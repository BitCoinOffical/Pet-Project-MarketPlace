package postgress

import (
	"context"
	"database/sql"

	"main.go/internal/domain/products"
)

type GetByIdRepo struct {
	DB *sql.DB
}

func NewGetByIdRepo(db *sql.DB) *GetByIdRepo {
	return &GetByIdRepo{DB: db}
}

func (db *GetByIdRepo) GetByIdRepo(ctx context.Context, id int) (*products.Product, error) {
	query := `SELECT * FROM products WHERE id = $1`
	row := db.DB.QueryRowContext(ctx, query, id)
	var product products.Product
	err := row.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.InStock, &product.CreatedAt)
	return &product, err
}
