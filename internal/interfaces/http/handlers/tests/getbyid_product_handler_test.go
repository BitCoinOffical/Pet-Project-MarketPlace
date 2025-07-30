package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"myapp/internal/domain/products"
	"myapp/internal/interfaces/http/handlers"
	usecase "myapp/internal/usecase/products"
)

func TestGetById(t *testing.T) {
	mock := &mockRepo{
		pr: &products.Product{
			ID:       1,
			Name:     "iPhone 15",
			Category: "Smartphone",
			Price:    1099.99,
			InStock:  true,
		},
		err: nil,
	}

	usecases := &usecase.UseCases{GetById: usecase.NewGetByIdProductUseCase(mock)}
	handler := handlers.NewGetByIdProductHandler(usecases)

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	w := httptest.NewRecorder()

	handler.GetByIdProductHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}

	expected := `{"ID":1,"Name":"iPhone 15","Category":"Smartphone","Price":1099.99,"InStock":true,"CreatedAt":"0001-01-01T00:00:00Z"}`
	clean := func(s string) string {
		return strings.TrimSpace(strings.ReplaceAll(s, " ", ""))
	}

	if clean(string(body)) != clean(expected) {
		t.Errorf("unexpected response body:\nwant: %s\ngot:  %s", expected, string(body))
	}
}
