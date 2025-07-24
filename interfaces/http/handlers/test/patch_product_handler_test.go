package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"main.go/interfaces/http/handlers"
	usecase "main.go/usecase/products"
	patch "main.go/usecase/products/method_patch"
)

func TestPatch(t *testing.T) {
	mock := &mockRepo{
		err: nil,
	}

	usecases := usecase.UseCases{PatchById: patch.NewPatchProductUseCase(mock)}
	handler := handlers.NewPatchProductHandler(&usecases)

	body := `{"name":"iPhone 15","category":"Smartphone","price":1099.99}`

	req := httptest.NewRequest(http.MethodPatch, "/products/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.PatchProductHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	if mock.num != 1 {
		t.Errorf("expected ID 1, got %d", mock.num)
	}

	if mock.ProductPatchDTO == nil || mock.ProductPatchDTO.Name == nil || *mock.ProductPatchDTO.Name != "iPhone 15" {
		t.Errorf("expected name 'iPhone 15', got %+v", mock.ProductPatchDTO)
	}

}
