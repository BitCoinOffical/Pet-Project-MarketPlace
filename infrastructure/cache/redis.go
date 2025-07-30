package cashe

import (
	"context"
	"time"

	"main.go/infrastructure/redis/storage"
)

func redisinit() error {
	cfg := storage.Config{
		Addr:        "localhost:6379",
		Password:    "",
		User:        "default",
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 10 * time.Second,
		Timeout:     5 * time.Second,
	}

	db, err :- storage.NewClient(context.Background(), cfg)
	if err !=  nil{
		return err
	}
	
}
