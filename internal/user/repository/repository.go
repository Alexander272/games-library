package repository

import (
	"context"

	"github.com/Alexander272/games-library/internal/user/models"
	"github.com/Alexander272/games-library/internal/user/repository/redis"
)

type IUser interface {
	Create(ctx context.Context, user models.User) (string, error)
	GetAll(ctx context.Context) ([]models.User, error)
	GetById(ctx context.Context, userId string) (models.User, error)
	GetByEmail(ctx context.Context, email string) (models.User, error)
	Update(ctx context.Context, user models.User) error
	Remove(ctx context.Context, userId string) error
}

type ISession interface {
	Create(ctx context.Context, token string, data redis.SessionData) error
	GetDel(ctx context.Context, key string) (data redis.SessionData, err error)
	Delete(ctx context.Context, key string) error
}
