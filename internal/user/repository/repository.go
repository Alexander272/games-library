package repository

import (
	"context"

	"github.com/Alexander272/games-library/internal/user/models"
)

type IUser interface {
	Create(ctx context.Context, user models.User) (id string, err error)
	GetAll(ctx context.Context) (users []models.User, err error)
	GetById(ctx context.Context, userId string) (user models.User, err error)
	GetByEmail(ctx context.Context, email string) (user models.User, err error)
	Update(ctx context.Context, user models.User) error
	Remove(ctx context.Context, userId string) error
}

type ISession interface{}
