package postgress

import (
	"context"
	"database/sql"
	"errors"
)

type DeletedByIdProductRepo struct {
	DB *sql.DB
}

func NewDeletedByIDProductRepo(db *sql.DB) *DeletedByIdProductRepo {
	return &DeletedByIdProductRepo{DB: db}
}

func (db *DeletedByIdProductRepo) Delete(ctx context.Context, id int) error {
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
