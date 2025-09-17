package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func NewClient(ctx context.Context, cfg Config) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return db, nil
}
