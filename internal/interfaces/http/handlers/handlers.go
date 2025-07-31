package handlers

import (
	cache "myapp/infrastructure/cache"
	usecase "myapp/internal/usecase/products"
)

type Handlers struct {
	Post        *PostProductHandler
	GetByID     *GetByIdProductHandler
	GetAll      *GetAllHandler
	PutByID     *PutProductHandler
	DeletedByID *DeletedByIDProductHandler
	Patch       *PatchProductHandler
}

func NewHandler(usecase *usecase.UseCases, cacheusecase *cache.GetAllCashe) *Handlers {
	return &Handlers{
		Post:        NewPostProductHandler(usecase),
		GetByID:     NewGetByIdProductHandler(usecase),
		GetAll:      NewGetAllHandler(cacheusecase),
		PutByID:     NewPutByIdProductHandler(usecase),
		DeletedByID: NewDeletedByIDProductHandler(usecase),
		Patch:       NewPatchProductHandler(usecase),
	}
}
