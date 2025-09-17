package search

import (
	"access_manager-service/internal/storage/dto"
	"context"
)

type UserSearcher interface {
	FindUserByEmail(ctx context.Context, email string) (dto.User, error)
	FindUserByUsername(ctx context.Context, username string) (string, string, error)
}
