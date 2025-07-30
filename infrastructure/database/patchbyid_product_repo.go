package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"myapp/internal/interfaces/http/dto"
	"myapp/pkg/sqlbuilder"
)

type PatchByIdProductRepo struct {
	DB *sql.DB
}

func NewPatchByIdProductRepo(db *sql.DB) *PatchByIdProductRepo {
	return &PatchByIdProductRepo{DB: db}
}
func (db *PatchByIdProductRepo) Patch(ctx context.Context, id int, product *dto.ProductPatchDTO) error {
	set, args, num, err := sqlbuilder.BuildPatchQuery(id, product)
	if err != nil {
		return err
	}
	query := fmt.Sprintf(`UPDATE products SET %s WHERE id = $%d`, strings.Join(set, ", "), num)
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
