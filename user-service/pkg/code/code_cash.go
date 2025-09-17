package code

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const (
	VALUE    = 72
	LIFETIME = 5
)

type CodeInRedis struct {
	rdb *redis.Client
}

func NewSaveCodeInRedis(rdb *redis.Client) *CodeInRedis {
	return &CodeInRedis{rdb: rdb}
}

func (rdb *CodeInRedis) SaveCodeRedis(ctx context.Context, email, code string) {
	key := fmt.Sprintf("LoginCode:%s", email)
	rdb.rdb.Set(ctx, key, code, time.Minute*LIFETIME)
}

func (rdb *CodeInRedis) CheckCodeRedis(ctx context.Context, key, code string) error {
	val, err := rdb.rdb.Get(ctx, key).Result()
	if err == redis.Nil || val != code {
		return errors.New("invalid code")
	}
	if err != nil {
		return err
	}
	err = rdb.rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
func (rdb *CodeInRedis) CreateSessionToken(w http.ResponseWriter, ctx context.Context, UserID int) error {
	sessionID := uuid.NewString()
	err := rdb.rdb.Set(ctx, sessionID, UserID, time.Hour*VALUE).Err()
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session_ID",
		Value:    sessionID,
		Expires:  time.Now().Add(time.Hour * VALUE),
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}
