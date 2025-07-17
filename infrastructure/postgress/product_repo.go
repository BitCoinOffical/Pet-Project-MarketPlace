package postgress

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"main.go/interfaces/http/dto"
	"main.go/internal/domain/products"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (db *ProductRepo) Create(ctx context.Context, product *products.Product) error {
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

func (db *ProductRepo) Get(ctx context.Context, id int) (*products.Product, error) {
	query := `SELECT * FROM products WHERE id = $1`
	row := db.DB.QueryRowContext(ctx, query, id)
	var product products.Product
	err := row.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.InStock, &product.CreatedAt)
	return &product, err
}

func (db *ProductRepo) Put(ctx context.Context, id int, product *products.Product) error {
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

func (db *ProductRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE id = $1`
	res, err := db.DB.ExecContext(ctx, query, id)
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

func (db *ProductRepo) Patch(ctx context.Context, id int, product *dto.ProductPatchDTO) error {
	var (
		set  = []string{}
		args = []interface{}{}
		num  = 1
	)
	if product.Name != nil {
		set = append(set, fmt.Sprintf("name = $%d", num))
		args = append(args, product.Name)
		num++
	}
	if product.Category != nil {
		set = append(set, fmt.Sprintf("category = $%d", num))
		args = append(args, product.Category)
		num++
	}
	if product.Price != nil {
		set = append(set, fmt.Sprintf("price = $%d", num))
		args = append(args, product.Price)
		num++
	}
	if product.InStock != nil {
		set = append(set, fmt.Sprintf("in_stock = $%d", num))
		args = append(args, product.InStock)
		num++
	}
	if len(set) == 0 {
		return errors.New("empty set")
	}
	query := fmt.Sprintf(`UPDATE products SET %s WHERE id = $%d"`, strings.Join(set, ", "), num)
	res, err := db.DB.ExecContext(ctx, query, args...)
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
