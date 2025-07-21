package get

import (
	"context"

	"main.go/internal/domain/products"
)

type GetByIdProductUseCase struct {
	Repo products.Repository
}

func NewGetByIdProductUseCase(repo products.Repository) *GetByIdProductUseCase {
	return &GetByIdProductUseCase{Repo: repo}
}
func (uc *GetByIdProductUseCase) GetById(ctx context.Context, id int) (*products.Product, error) {
	return uc.Repo.GetByIdRepo(ctx, id)
}
