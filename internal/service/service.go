package service

import (
	"time"

	"github.com/Alexander272/games-library/internal/repository"
	"github.com/Alexander272/games-library/pkg/auth"
	"github.com/Alexander272/games-library/pkg/hasher"
	"github.com/Alexander272/games-library/pkg/storage"
)

type Services struct{}

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
	return &Services{}
}
