package put

import (
	"context"

	"main.go/interfaces/http/dto"
	"main.go/internal/domain/products"
)

type PutByIDProductUseCase struct {
	Repo products.Repository
}

func NewPutByIdProductUseCase(repo products.Repository) *PutByIDProductUseCase {
	return &PutByIDProductUseCase{Repo: repo}
}

func (uc *PutByIDProductUseCase) PutByID(ctx context.Context, id int, dto dto.ProductPutDTO) error {
	product := &products.Product{
		Name:     dto.Name,
		Category: dto.Category,
		Price:    dto.Price,
		InStock:  dto.InStock,
	}
	return uc.Repo.Put(ctx, id, product)
}
