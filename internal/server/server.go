package server

import (
	"database/sql"
	"log"
	"net/http"

	"main.go/infrastructure/postgress"
	"main.go/interfaces/http/handlers"
	usecase "main.go/usecase/products"
)

func Run() error {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=5286 dbname=productsdb sslmode=disable")
	if err != nil {
		return err
	}
	if err := runMigrations(db); err != nil {
		return err
	}
	ProductRepo := postgress.NewProductRepo(db)
	productUsecase := usecase.NewProductUseCase(ProductRepo)
	handler := handlers.NewHandler(productUsecase)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /products", handler.Post.PostProductHandler)
	mux.HandleFunc("GET /products/", handler.GetByID.GetByIdProductHandler)
	mux.HandleFunc("PUT /products/", handler.PutByID.PutByIdProductHandler)
	mux.HandleFunc("DELETE /products/", handler.DeletedByID.DeletedByIDProductHandler)
	mux.HandleFunc("PATCH /products/", handler.Patch.PatchProductHandler)
	log.Println("Server started at :8080")
	return http.ListenAndServe(":8080", mux)
}

func runMigrations(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		category VARCHAR(100),
		price NUMERIC,
		in_stock BOOLEAN,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}
