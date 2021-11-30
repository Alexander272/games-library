package repository

import (
	"context"

	"github.com/Alexander272/games-library/internal/user/models"
	user "github.com/Alexander272/games-library/internal/user/repository/mongo"
	ses "github.com/Alexander272/games-library/internal/user/repository/redis"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
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
	Create(ctx context.Context, token string, data ses.SessionData) error
	GetDel(ctx context.Context, key string) (data ses.SessionData, err error)
	Delete(ctx context.Context, key string) error
}

func NewUserRepo(db *mongo.Database, collection string) IUser {
	return user.NewUserRepo(db, collection)
}

func NewSessionRepo(redis *redis.Client) ISession {
	return ses.NewSessionRepo(redis)
}
