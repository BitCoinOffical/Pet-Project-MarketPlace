package database

import (
	"context"
	"database/sql"
	"fmt"
	"myapp/internal/interfaces/http/dto"
	"myapp/pkg/sqlbuilder"
)

type GetAll struct {
	DB *sql.DB
}

func NewGetAll(db *sql.DB) *GetAll {
	return &GetAll{DB: db}
}

func (db *GetAll) GetAll(ctx context.Context, filter *dto.ProductFilterDTO) ([]dto.ProductResponse, int, error) {
	res, limitOffset, args, err := sqlbuilder.BuildGetAllQuery(filter)
	if err != nil {
		return nil, 0, err
	}
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM products %s", res)
	var totalCount int
	err = db.DB.QueryRowContext(ctx, countQuery, args[:len(args)-2]...).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}
	query := fmt.Sprintf(`SELECT id, name, category, price, in_stock FROM products %s ORDER BY price %s`, res, limitOffset)
	rows, err := db.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	var allproduct []dto.ProductResponse
	defer rows.Close()
	for rows.Next() {
		var product dto.ProductResponse
		if err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.InStock); err != nil {
			return nil, 0, err
		}
		allproduct = append(allproduct, product)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	return allproduct, totalCount, nil
}
