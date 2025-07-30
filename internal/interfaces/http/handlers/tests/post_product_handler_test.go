package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"myapp/internal/interfaces/http/handlers"
	usecase "myapp/internal/usecase/products"
)

func TestPostProduct(t *testing.T) {
	mock := &mockRepo{err: nil}

	usecases := usecase.UseCases{Create: usecase.NewPostProductUseCase(mock)}
	handler := handlers.NewPostProductHandler(&usecases)

	body := `{"name":"iPhone 15","category":"Smartphone","price":1099.99,"in_stock":true}`
	req := httptest.NewRequest(http.MethodPost, "/products/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.PostProductHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	if mock.pr == nil || mock.pr.Name != "iPhone 15" {
		t.Errorf("expected name 'iPhone 15', got %+v", mock.pr)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected 201 Created, got %d", resp.StatusCode)
	}
}
