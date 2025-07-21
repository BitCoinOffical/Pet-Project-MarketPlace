package post

import (
	"context"

	"main.go/interfaces/http/dto"
	"main.go/internal/domain/products"
)

type PostProductUseCase struct {
	Repo products.Repository
}

func NewPostProductUseCase(repo products.Repository) *PostProductUseCase {
	return &PostProductUseCase{Repo: repo}
}

func (uc *PostProductUseCase) Execute(ctx context.Context, dto dto.ProductCreateDTO) error {
	product := &products.Product{
		Name:     dto.Name,
		Category: dto.Category,
		Price:    dto.Price,
		InStock:  dto.InStock,
	}
	return uc.Repo.Create(ctx, product)
}
