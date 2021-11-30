package user

import (
	"time"

	"github.com/Alexander272/games-library/internal/user/repository"
	"github.com/Alexander272/games-library/internal/user/service"
	"github.com/Alexander272/games-library/pkg/auth"
	"github.com/Alexander272/games-library/pkg/hasher"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepo interface {
	repository.IUser
}
type ISesRepo interface {
	repository.ISession
}

type IUserService interface {
	service.IUser
}
type IAuthService interface {
	service.IAuth
}

func NewUserRepo(db *mongo.Database, collection string) IUserRepo {
	return repository.NewUserRepo(db, collection)
}
func NewSessionRepo(redis *redis.Client) ISesRepo {
	return repository.NewSessionRepo(redis)
}

func NewUserService(repo repository.IUser, hasher hasher.IPasswordHasher) IUserService {
	return service.NewUserService(repo, hasher)
}
func NewAuthService(repo repository.IUser, ses repository.ISession, tokenManager auth.ITokenManager, hasher hasher.IPasswordHasher,
	accessTTL time.Duration, refreshTTL time.Duration, domain string) IAuthService {
	return service.NewAuthService(
		repo, ses, tokenManager, hasher, accessTTL, refreshTTL, domain,
	)
}
