package database

import (
	"context"
	"database/sql"
	"errors"

	"myapp/internal/domain/products"
)

type PutByIdProductRepo struct {
	DB *sql.DB
}

func NewPutByIdProductRepo(db *sql.DB) *PutByIdProductRepo {
	return &PutByIdProductRepo{DB: db}
}
func (db *PutByIdProductRepo) Put(ctx context.Context, id int, product *products.Product) error {
	query := `UPDATE products SET name = $1, category = $2, price = $3, in_stock = $4 WHERE id = $5`
	res, err := db.DB.ExecContext(ctx, query, &product.Name, &product.Category, &product.Price, &product.InStock, id)
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
