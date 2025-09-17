package db

import (
	"access_manager-service/config"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func UsersDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPass,
		cfg.PostgresDB,
		cfg.PostgresSslmode))
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(context.Background()); err != nil {
		db.Close()
		return nil, err
	}
	query := `CREATE TABLE IF NOT EXISTS users(
		id BIGSERIAL PRIMARY KEY,
		name TEXT NOT NULL CHECK (char_length(name) BETWEEN 2 AND 16),
		username TEXT NOT NULL UNIQUE CHECK (char_length(username) BETWEEN 2 AND 16),
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL
	);`
	_, err = db.ExecContext(context.Background(), query)
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
