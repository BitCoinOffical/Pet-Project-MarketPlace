package usecase

import "main.go/internal/domain/products"

type UseCases struct {
	Create      *PostProductUseCase
	GetById     *GetByIdProductUseCase
	GetAll      *GetAllProductUseCase
	PutByID     *PutByIDProductUseCase
	DeletedByID *DeletedByIDProductUseCase
	PatchById   *PatchProductUseCase
}

func NewProductUseCase(Repository products.Repository) *UseCases {
	return &UseCases{
		Create:      NewPostProductUseCase(Repository),
		GetById:     NewGetByIdProductUseCase(Repository),
		GetAll:      NewGetAllUseCase(Repository),
		PutByID:     NewPutByIdProductUseCase(Repository),
		DeletedByID: NewDeletedByIDProductUseCase(Repository),
		PatchById:   NewPatchProductUseCase(Repository),
	}
}
