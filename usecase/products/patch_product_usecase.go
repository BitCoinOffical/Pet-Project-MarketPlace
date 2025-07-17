package usecase

import (
	"context"

	"main.go/interfaces/http/dto"
	"main.go/internal/domain/products"
)

type PatchProductUseCase struct {
	Repo products.Repository
}

func NewPatchProductUseCase(repo products.Repository) *PatchProductUseCase {
	return &PatchProductUseCase{Repo: repo}
}

func (uc *PatchProductUseCase) PatchProductUseCase(ctx context.Context, id int, dto *dto.ProductPatchDTO) error {
	return uc.Repo.Patch(ctx, id, dto)
}
