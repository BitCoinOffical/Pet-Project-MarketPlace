package search

import (
	"access_manager-service/internal/storage/dto"
	"context"
	"database/sql"
)

type Searcher struct {
	db *sql.DB
}

func NewSearcher(db *sql.DB) *Searcher {
	return &Searcher{db: db}
}

func (db *Searcher) FindUserByEmail(ctx context.Context, email string) (dto.User, error) {
	query := `SELECT id, name, username, email FROM users WHERE email ILIKE $1`
	row := db.db.QueryRowContext(ctx, query, email)
	var user dto.User
	err := row.Scan(&user.ID, &user.Name, &user.UserName, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}
	return user, nil
}

func (db *Searcher) FindUserByUsername(ctx context.Context, username string) (string, string, error) {
	query := `SELECT username, password FROM users WHERE username = $1`
	row := db.db.QueryRowContext(ctx, query, username)
	var dbUsername, dbPassword string
	err := row.Scan(&dbUsername, &dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", err
		}
		return "", "", err
	}
	return dbUsername, dbPassword, err
}
