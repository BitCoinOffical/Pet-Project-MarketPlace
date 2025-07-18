package products

import (
	"context"

	"main.go/interfaces/http/dto"
)

type Repository interface {
	Create(ctx context.Context, product *Product) error
	GetByIdRepo(ctx context.Context, id int) (*Product, error)
	GetAll(ctx context.Context, filter *dto.ProductFilterDTO) ([]dto.ProductResponse, int, error)
	Put(ctx context.Context, id int, product *Product) error
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, id int, product *dto.ProductPatchDTO) error
}
