package get

import (
	"context"

	"main.go/interfaces/http/dto"
	"main.go/internal/domain/products"
)

const (
	Limit = 10
	Page  = 1
)

type GetAllProductUseCase struct {
	Repo products.Repository
}

func NewGetAllUseCase(repo products.Repository) *GetAllProductUseCase {
	return &GetAllProductUseCase{Repo: repo}
}

func (uc *GetAllProductUseCase) GetAll(ctx context.Context, category, search string, page int, min_price, max_price float64) ([]dto.ProductResponse, int, error) {
	if page == 0 {
		page = Page
	}

	var categoryPtr *string
	if category != "" {
		categoryPtr = &category
	}

	var searchPtr *string
	if search != "" {
		searchPtr = &search
	}

	var minPricePtr *float64
	if min_price != 0 {
		minPricePtr = &min_price
	}

	var maxPricePtr *float64
	if max_price != 0 {
		maxPricePtr = &max_price
	}

	filter := &dto.ProductFilterDTO{
		Category: categoryPtr,
		Search:   searchPtr,
		MinPrice: minPricePtr,
		MaxPrice: maxPricePtr,
		Page:     page,
		Limit:    Limit,
	}
	return uc.Repo.GetAll(ctx, filter)
}
