package handlers

import (
	usecase "myapp/internal/usecase/products"
	"myapp/pkg/httphelper"
	"net/http"
	"strconv"
	"strings"
)

type DeletedByIDProductHandler struct {
	usecase *usecase.UseCases
}

func NewDeletedByIDProductHandler(usecase *usecase.UseCases) *DeletedByIDProductHandler {
	return &DeletedByIDProductHandler{usecase: usecase}
}

func (h *DeletedByIDProductHandler) DeletedByIDProductHandler(w http.ResponseWriter, r *http.Request) {
	idstr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.usecase.DeletedByID.DeletedByID(r.Context(), id); err != nil {
		http.Error(w, "Failed to deleted product: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := httphelper.RespondJSON(w, http.StatusNoContent, map[string]string{"message": "deleted"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
