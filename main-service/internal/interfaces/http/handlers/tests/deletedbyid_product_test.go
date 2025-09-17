package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"myapp/internal/interfaces/http/handlers"
	usecase "myapp/internal/usecase/products"
)

func TestDeleteHandlers(t *testing.T) {
	repo := &mockRepo{}
	usecases := &usecase.UseCases{
		DeletedByID: usecase.NewDeletedByIDProductUseCase(repo),
	}

	handler := handlers.NewDeletedByIDProductHandler(usecases)

	req := httptest.NewRequest(http.MethodDelete, "/products/7", nil)
	w := httptest.NewRecorder()

	handler.DeletedByIDProductHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if repo.CalledWithID != 7 {
		t.Errorf("expected ID 7, got %d", repo.CalledWithID)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected status 204 No Content, got %d", resp.StatusCode)
	}
}
