package postgress

import (
	"context"
	"database/sql"

	"main.go/interfaces/http/dto"
	"main.go/internal/domain/products"
)

type ProductRepo struct {
	CreateProductDB  *CreateProductRepo
	DeletedByIdDB    *DeletedByIdProductRepo
	GetByIdProductDB *GetByIdRepo
	PatchByIdDB      *PatchByIdProductRepo
	PutByIdDB        *PutByIdProductRepo
	GetAllDB         *GetAll
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{
		CreateProductDB:  NewCreateProductRepo(db),
		DeletedByIdDB:    NewDeletedByIDProductRepo(db),
		GetByIdProductDB: NewGetByIdRepo(db),
		PatchByIdDB:      NewPatchByIdProductRepo(db),
		PutByIdDB:        NewPutByIdProductRepo(db),
		GetAllDB:         NewGetAll(db),
	}
}

func (p *ProductRepo) Create(ctx context.Context, product *products.Product) error {
	return p.CreateProductDB.Create(ctx, product)
}

func (p *ProductRepo) GetByIdRepo(ctx context.Context, id int) (*products.Product, error) {
	return p.GetByIdProductDB.GetByIdRepo(ctx, id)
}

func (p *ProductRepo) Delete(ctx context.Context, id int) error {
	return p.DeletedByIdDB.Delete(ctx, id)
}

func (p *ProductRepo) Patch(ctx context.Context, id int, product *dto.ProductPatchDTO) error {
	return p.PatchByIdDB.Patch(ctx, id, product)
}

func (p *ProductRepo) Put(ctx context.Context, id int, product *products.Product) error {
	return p.PutByIdDB.Put(ctx, id, product)

}

func (p *ProductRepo) GetAll(ctx context.Context, filter *dto.ProductFilterDTO) ([]dto.ProductResponse, int, error) {
	return p.GetAllDB.GetAll(ctx, filter)
}
