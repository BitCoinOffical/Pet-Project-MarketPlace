package auth

import (
	"access_manager-service/internal/storage/dto"
	"context"
)

type Registrator interface {
	UserRegister(ctx context.Context, regDTO *dto.Register) error
}
