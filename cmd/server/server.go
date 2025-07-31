package server

import (
	"database/sql"
	"log"
	"myapp/infrastructure/cache"
	"myapp/infrastructure/database"
	"myapp/internal/interfaces/http/handlers"
	usecase "myapp/internal/usecase/products"
	"net/http"

	_ "github.com/lib/pq"
)

func Run() error {
	db, err := sql.Open("postgres", "host=db port=5432 user=postgres password=5286 dbname=productsdb sslmode=disable")
	if err != nil {
		return err
	}
	if err := runMigrations(db); err != nil {
		return err
	}
	rdb, err := cache.Redisinit()
	if err != nil {
		return err
	}
	ProductRepo := database.NewProductRepo(db)
	productUsecase := usecase.NewProductUseCase(ProductRepo)
	productCache := cache.NewGetAllCash(productUsecase, rdb)
	handler := handlers.NewHandler(productUsecase, productCache)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /products", handler.Post.PostProductHandler)
	mux.HandleFunc("GET /products", handler.GetAll.GetAllHandler)
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
