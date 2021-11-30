package repository

import (
	"github.com/Alexander272/games-library/internal/user"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Session user.ISesRepo
	User    user.IUserRepo
}

func NewRepo(db *mongo.Database, redis *redis.Client) *Repo {
	return &Repo{
		Session: user.NewSessionRepo(redis),
		User:    user.NewUserRepo(db, usersCollection),
	}
}
