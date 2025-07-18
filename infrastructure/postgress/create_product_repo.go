package postgress

import (
	"context"
	"database/sql"
	"errors"

	"main.go/internal/domain/products"
)

type CreateProductRepo struct {
	DB *sql.DB
}

func NewCreateProductRepo(db *sql.DB) *CreateProductRepo {
	return &CreateProductRepo{DB: db}
}

func (db *CreateProductRepo) Create(ctx context.Context, product *products.Product) error {
	query := `INSERT INTO products (name, category, price, in_stock) VALUES ($1, $2, $3, $4)`
	res, err := db.DB.ExecContext(ctx, query, &product.Name, &product.Category, &product.Price, &product.InStock)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no product found with given ID")
	}
	return nil
}
