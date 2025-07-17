package usecase

import "main.go/internal/domain/products"

type UseCases struct {
	Create      *PostProductUseCase
	GetById     *GetByIdProductUseCase
	PutByID     *PutByIDProductUseCase
	DeletedByID *DeletedByIDProductUseCase
	Patch       *PatchProductUseCase
}

func NewProductUseCase(Repository products.Repository) *UseCases {
	return &UseCases{
		Create:      NewPostProductUseCase(Repository),
		GetById:     NewGetByIdProductUseCase(Repository),
		PutByID:     NewPutByIdProductUseCase(Repository),
		DeletedByID: NewDeletedByIDProductUseCase(Repository),
		Patch:       NewPatchProductUseCase(Repository),
	}
}
