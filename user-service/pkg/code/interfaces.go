package code

import (
	"context"
	"net/http"
)

type CodeSaver interface {
	SaveCodeRedis(ctx context.Context, email, code string)
	CheckCodeRedis(ctx context.Context, key, code string) error
	CreateSessionToken(w http.ResponseWriter, ctx context.Context, UserID int) error
}
