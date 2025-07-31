package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"myapp/infrastructure/cache"
	"myapp/internal/interfaces/http/dto"
	"myapp/internal/interfaces/http/handlers"
	usecase "myapp/internal/usecase/products"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func TestGetAll(t *testing.T) {
	s := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{
		Addr:     s.Addr(),
		Password: "",
		DB:       0,
	})
	mock := &mockRepo{
		filters: []dto.ProductResponse{
			{
				ID:       1,
				Name:     "iPhone 15",
				Category: "Smartphone",
				Price:    1099.99,
				InStock:  true,
			},
		},
		err: nil,
	}

	usecase := usecase.NewProductUseCase(mock)
	cache := cache.NewGetAllCash(usecase, rdb)
	handler := handlers.NewGetAllHandler(cache)

	req := httptest.NewRequest(http.MethodGet, "/products/", nil)
	w := httptest.NewRecorder()

	handler.GetAllHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}
	expected := `{"items":[{"id":1,"name":"iPhone 15","category":"Smartphone","price":1099.99,"in_stock":true}],"total_count":1}`
	clean := func(s string) string {
		return strings.TrimSpace(strings.ReplaceAll(s, " ", ""))
	}

	if clean(string(body)) != clean(expected) {
		t.Errorf("unexpected response body:\nwant: %s\ngot:  %s", expected, string(body))
	}
}
