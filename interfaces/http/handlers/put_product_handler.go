package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"main.go/interfaces/http/dto"
	"main.go/pkg/httphelper"
	usecase "main.go/usecase/products"
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
