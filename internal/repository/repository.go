package repository

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct{}

func NewRepo(db *mongo.Database, redis redis.Cmdable) *Repo {
	return &Repo{}
}
