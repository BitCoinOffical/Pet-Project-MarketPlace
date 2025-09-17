package rediscash

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
	cashdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := cashdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return cashdb, nil
}

func RedisCashInit() (*redis.Client, error) {
	ctx := context.Background()
	cfg := Config{
		Addr:     "redis:6380",
		Password: "",
		DB:       0,
	}
	rdb, err := NewClient(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
