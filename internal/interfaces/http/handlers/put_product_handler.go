package handlers

import (
	"encoding/json"
	"myapp/internal/interfaces/http/dto"
	"net/http"
	"strconv"
	"strings"

	usecase "myapp/internal/usecase/products"
	"myapp/pkg/httphelper"

	"github.com/go-playground/validator/v10"
)

type PutProductHandler struct {
	usecase *usecase.UseCases
}

func NewPutByIdProductHandler(usecase *usecase.UseCases) *PutProductHandler {
	return &PutProductHandler{usecase: usecase}
}

func (h *PutProductHandler) PutByIdProductHandler(w http.ResponseWriter, r *http.Request) {
	var validate = validator.New()
	var ProductPutDTO dto.ProductPutDTO
	idstr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "id convert error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&ProductPutDTO); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(ProductPutDTO); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.usecase.PutByID.PutByID(r.Context(), id, ProductPutDTO); err != nil {
		http.Error(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := httphelper.RespondJSON(w, http.StatusOK, map[string]string{"message": "updated"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
