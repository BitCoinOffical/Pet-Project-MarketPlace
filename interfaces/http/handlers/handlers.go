package handlers

import usecase "main.go/usecase/products"

type Handlers struct {
	Post        *PostProductHandler
	GetByID     *GetByIdProductHandler
	PutByID     *PutProductHandler
	DeletedByID *DeletedByIDProductHandler
	Patch       *PatchProductHandler
}

func NewHandler(usecase *usecase.UseCases) *Handlers {
	return &Handlers{
		Post:        NewPostProductHandler(usecase),
		GetByID:     NewGetByIdProductHandler(usecase),
		PutByID:     NewPutByIdProductHandler(usecase),
		DeletedByID: NewDeletedByIDProductHandler(usecase),
		Patch:       NewPatchProductHandler(usecase),
	}
}
