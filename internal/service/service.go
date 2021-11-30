package service

import (
	"time"

	"github.com/Alexander272/games-library/internal/repository"
	"github.com/Alexander272/games-library/internal/user"
	"github.com/Alexander272/games-library/pkg/auth"
	"github.com/Alexander272/games-library/pkg/hasher"
	"github.com/Alexander272/games-library/pkg/storage"
)

type Services struct {
	Auth user.IAuthService
	User user.IUserService
}

type Deps struct {
	Repos           *repository.Repo
	StorageProvider storage.Provider
	Hasher          hasher.IPasswordHasher
	TokenManager    auth.ITokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	Domain          string
}

func NewServices(deps Deps) *Services {
	return &Services{
		Auth: user.NewAuthService(
			deps.Repos.User,
			deps.Repos.Session,
			deps.TokenManager,
			deps.Hasher,
			deps.AccessTokenTTL,
			deps.RefreshTokenTTL,
			deps.Domain,
		),
		User: user.NewUserService(deps.Repos.User, deps.Hasher),
	}
}
