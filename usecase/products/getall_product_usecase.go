package usecase

import (
	"context"

	"main.go/interfaces/http/dto"
	"main.go/internal/domain/products"
)

const (
	Limit = 10
)

type GetAllProductUseCase struct {
	Repo products.Repository
}

func NewGetAllUseCase(repo products.Repository) *GetAllProductUseCase {
	return &GetAllProductUseCase{Repo: repo}
}

func (uc *GetAllProductUseCase) GetAllUseCase(ctx context.Context, category, search string, page int, min_price, max_price float64) ([]dto.ProductResponse, int, error) {
	filter := &dto.ProductFilterDTO{
		Category: &category,
		MinPrice: &min_price,
		MaxPrice: &max_price,
		Search:   &search,
		Page:     page,
		Limit:    Limit,
	}
	return uc.Repo.GetAll(ctx, filter)
}
