package usecase

import (
	"context"

	"main.go/internal/domain/products"
)

type DeletedByIDProductUseCase struct {
	Repo products.Repository
}

func NewDeletedByIDProductUseCase(repo products.Repository) *DeletedByIDProductUseCase {
	return &DeletedByIDProductUseCase{Repo: repo}
}

func (uc *DeletedByIDProductUseCase) DeletedByID(ctx context.Context, id int) error {
	return uc.Repo.Delete(ctx, id)
}
