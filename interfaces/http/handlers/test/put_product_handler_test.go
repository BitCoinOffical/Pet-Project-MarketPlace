package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"main.go/interfaces/http/handlers"
	usecase "main.go/usecase/products"
	put "main.go/usecase/products/method_put"
)

func TestPutProduct(t *testing.T) {
	mock := &mockRepo{err: nil}

	usecases := usecase.UseCases{PutByID: put.NewPutByIdProductUseCase(mock)}
	handler := handlers.NewPutByIdProductHandler(&usecases)

	body := `{"name":"iPhone 15","category":"Smartphone","price":1099.99,"in_stock":true}`
	req := httptest.NewRequest(http.MethodPut, "/products/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.PutByIdProductHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", resp.StatusCode)
	}

	if mock.num != 1 {
		t.Errorf("expected ID 1, got %d", mock.num)
	}

	if mock.pr == nil || mock.pr.Name != "iPhone 15" {
		t.Errorf("expected name 'iPhone 15', got %+v", mock.pr)
	}
}
