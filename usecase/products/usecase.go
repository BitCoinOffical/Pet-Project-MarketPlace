package usecase

import (
	"main.go/internal/domain/products"
	delete "main.go/usecase/products/method_delete"
	get "main.go/usecase/products/method_get"
	patch "main.go/usecase/products/method_patch"
	post "main.go/usecase/products/method_post"
	put "main.go/usecase/products/method_put"
)

type UseCases struct {
	Create      *post.PostProductUseCase
	GetById     *get.GetByIdProductUseCase
	GetAll      *get.GetAllProductUseCase
	PutByID     *put.PutByIDProductUseCase
	DeletedByID *delete.DeletedByIDProductUseCase
	PatchById   *patch.PatchProductUseCase
}

func NewProductUseCase(Repository products.Repository) *UseCases {
	return &UseCases{
		Create:      post.NewPostProductUseCase(Repository),
		GetById:     get.NewGetByIdProductUseCase(Repository),
		GetAll:      get.NewGetAllUseCase(Repository),
		PutByID:     put.NewPutByIdProductUseCase(Repository),
		DeletedByID: delete.NewDeletedByIDProductUseCase(Repository),
		PatchById:   patch.NewPatchProductUseCase(Repository),
	}
}
