package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Alexander272/games-library/pkg/logger"
	"github.com/go-redis/redis/v8"
)

type SessionRepo struct {
	db *redis.Client
}

func NewSessionRepo(db *redis.Client) *SessionRepo {
	return &SessionRepo{db: db}
}

type SessionData struct {
	UserId string
	Email  string
	Role   string
	Ua     string
	Ip     string
	Exp    time.Duration
}

func (d SessionData) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}

func (r *SessionRepo) Create(ctx context.Context, token string, data SessionData) error {
	res := r.db.Set(ctx, token, data, data.Exp)
	if res.Err() != nil {
		return fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	logger.Debug(res.Result())

	return nil
}

func (r *SessionRepo) GetDel(ctx context.Context, key string) (data SessionData, err error) {
	cmd := r.db.GetDel(ctx, key)
	if cmd.Err() != nil {
		return data, fmt.Errorf("failed to execute query. error: %w", cmd.Err())
	}

	if err := cmd.Scan(&data); err != nil {
		// todo дописать ошибку
		return data, fmt.Errorf("failed to ... . error: %w", err)
	}
	logger.Info(data)
	return data, nil
}

func (r *SessionRepo) Delete(ctx context.Context, key string) error {
	res := r.db.Del(ctx, key)
	if res.Err() != nil {
		return fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	logger.Debug(res)
	return nil
}
