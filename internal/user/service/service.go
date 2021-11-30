package service

import (
	"context"
	"net/http"

	"github.com/Alexander272/games-library/internal/user/models"
)

const CookieName = "session"

type IUser interface {
	Create(ctx context.Context, dto models.CreateUserDTO) (string, error)
	GetAll(ctx context.Context) ([]models.User, error)
	GetById(ctx context.Context, userId string) (models.User, error)
	Update(ctx context.Context, dto models.UpdateUserDTO) error
	Delete(ctx context.Context, userId string) error
}

type IAuth interface {
	SignIn(ctx context.Context, dto models.SignInUserDTO, ua, ip string) (models.Token, http.Cookie, error)
	SignOut(ctx context.Context, token string) (http.Cookie, error)
	Refresh(ctx context.Context, refToken, ua, ip string) (models.Token, http.Cookie, error)
}
