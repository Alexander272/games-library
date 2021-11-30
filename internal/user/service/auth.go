package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Alexander272/games-library/internal/user/models"
	"github.com/Alexander272/games-library/internal/user/repository"
	"github.com/Alexander272/games-library/internal/user/repository/redis"
	"github.com/Alexander272/games-library/pkg/auth"
	"github.com/Alexander272/games-library/pkg/hasher"
	"github.com/Alexander272/games-library/pkg/logger"
)

type AuthService struct {
	repo            repository.IUser
	session         repository.ISession
	tokenManager    auth.ITokenManager
	hasher          hasher.IPasswordHasher
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	domain          string
}

func NewAuthService(repo repository.IUser, ses repository.ISession, tokenManager auth.ITokenManager, hasher hasher.IPasswordHasher,
	accessTTL, refreshTTL time.Duration, domain string) *AuthService {
	return &AuthService{
		repo:            repo,
		session:         ses,
		tokenManager:    tokenManager,
		hasher:          hasher,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
		domain:          domain,
	}
}

func (s *AuthService) SignIn(ctx context.Context, dto models.SignInUserDTO, ua, ip string) (token models.Token, cookie http.Cookie, err error) {
	user, err := s.repo.GetByEmail(ctx, dto.Email)
	if err != nil {
		logger.Errorf("failed to find user by email. errors: %s", err.Error())
		if errors.Is(err, models.ErrUserNotFound) {
			return token, cookie, err
		}
		return token, cookie, fmt.Errorf("invalid credentials")
	}

	if ok := s.hasher.CheckPasswordHash(dto.Password, user.Password); !ok {
		return token, cookie, fmt.Errorf("invalid credentials")
	}

	accessToken, err := s.tokenManager.NewJWT(user.Id, user.Email, user.Role, s.accessTokenTTL)
	if err != nil {
		logger.Errorf("failed to generate jwt token. error: %s", err.Error())
		return token, cookie, fmt.Errorf("failed to sign in")
	}
	refreshToken, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		logger.Errorf("failed to generate refresh token. error: %s", err.Error())
		return token, cookie, fmt.Errorf("failed to sign in")
	}

	data := redis.SessionData{
		UserId: user.Id,
		Email:  user.Email,
		Role:   user.Role,
		Ua:     ua,
		Ip:     ip,
		Exp:    s.refreshTokenTTL,
	}

	if err = s.session.Create(ctx, refreshToken, data); err != nil {
		logger.Errorf("failed to create session. error: %s", err.Error())
		return token, cookie, fmt.Errorf("failed to create session")
	}

	cookie = http.Cookie{
		Name:     CookieName,
		Value:    refreshToken,
		MaxAge:   int(s.refreshTokenTTL.Seconds()),
		Path:     "/",
		Domain:   s.domain,
		Secure:   false,
		HttpOnly: true,
	}

	return models.Token{AccessToken: accessToken}, cookie, nil
}

func (s *AuthService) SignOut(ctx context.Context, token string) (cookie http.Cookie, err error) {
	cookie = http.Cookie{
		Name:     CookieName,
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		Domain:   s.domain,
		Secure:   false,
		HttpOnly: true,
	}

	if err = s.session.Delete(ctx, token); err != nil {
		logger.Errorf("failed to remove session. error: %s", err.Error())
		return http.Cookie{}, fmt.Errorf("failed to remove session")
	}
	return cookie, nil
}

func (s *AuthService) Refresh(ctx context.Context, refToken, ua, ip string) (token models.Token, cookie http.Cookie, err error) {
	data, err := s.session.GetDel(ctx, refToken)
	if err != nil {
		logger.Errorf("failed to getdel session. error: %s", err.Error())
		return token, cookie, fmt.Errorf("failed to get session")
	}

	if ua != data.Ua || ip != data.Ip {
		if ua != data.Ua {
			logger.Errorf("ua do not match. ua from redis %s, ua from request %s", data.Ua, ua)
		}
		if ip != data.Ip {
			logger.Errorf("ip do not match. ip from redis %s, ip from request %s", data.Ip, ip)
		}
		return token, cookie, fmt.Errorf("invalid credentials")
	}

	accessToken, err := s.tokenManager.NewJWT(data.UserId, data.Email, data.Role, s.accessTokenTTL)
	if err != nil {
		logger.Errorf("failed to generate jwt token. error: %s", err.Error())
		return token, cookie, fmt.Errorf("failed to sign in")
	}
	refreshToken, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		logger.Errorf("failed to generate refresh token. error: %s", err.Error())
		return token, cookie, fmt.Errorf("failed to sign in")
	}

	newData := redis.SessionData{
		UserId: data.UserId,
		Email:  data.Email,
		Role:   data.Role,
		Ua:     ua,
		Ip:     ip,
		Exp:    s.refreshTokenTTL,
	}

	if err = s.session.Create(ctx, refreshToken, newData); err != nil {
		logger.Errorf("failed to create session. error: %s", err.Error())
		return token, cookie, fmt.Errorf("failed to create session")
	}

	cookie = http.Cookie{
		Name:     CookieName,
		Value:    refreshToken,
		MaxAge:   int(s.refreshTokenTTL.Seconds()),
		Path:     "/",
		Domain:   s.domain,
		Secure:   false,
		HttpOnly: true,
	}
	token = models.Token{AccessToken: accessToken}

	return token, cookie, nil
}

func (s *AuthService) TokenParse(token string) (userId string, role string, err error) {
	claims, err := s.tokenManager.Parse(token)
	if err != nil {
		return userId, role, fmt.Errorf("failed to parse token. error: %s", err.Error())
	}
	return claims["userId"].(string), claims["role"].(string), err
}
