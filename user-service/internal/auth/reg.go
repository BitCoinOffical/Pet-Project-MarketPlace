package auth

import (
	"access_manager-service/internal/storage/dto"
	passwordhash "access_manager-service/pkg/hash"
	"context"
	"database/sql"
	"errors"
)

type Registration struct {
	db *sql.DB
}

func NewRegistation(db *sql.DB) *Registration {
	return &Registration{db: db}
}

func (r *Registration) UserRegister(ctx context.Context, regDTO *dto.Register) error {
	passhash, err := passwordhash.HashPassword(regDTO.Password)
	if err != nil {
		return err
	}
	reg := `INSERT INTO users (name, username, email, password_hash) VALUES ($1, $2, $3, $4)`
	res, err := r.db.ExecContext(ctx, reg, regDTO.Name, regDTO.UserName, regDTO.Email, passhash)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("failed registration user")
	}
	return nil
}
