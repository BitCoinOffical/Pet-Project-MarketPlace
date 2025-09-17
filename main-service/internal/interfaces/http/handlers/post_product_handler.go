package handlers

import (
	"encoding/json"
	"myapp/internal/interfaces/http/dto"
	"net/http"

	"github.com/go-playground/validator/v10"

	usecase "myapp/internal/usecase/products"
	"myapp/pkg/httphelper"
)

type PostProductHandler struct {
	usecase *usecase.UseCases
}

func NewPostProductHandler(usecase *usecase.UseCases) *PostProductHandler {
	return &PostProductHandler{usecase: usecase}
}

func (h *PostProductHandler) PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var validate = validator.New()
	var ProductCreateDTO dto.ProductCreateDTO

	if err := json.NewDecoder(r.Body).Decode(&ProductCreateDTO); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(ProductCreateDTO); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.usecase.Create.Execute(r.Context(), ProductCreateDTO); err != nil {
		http.Error(w, "Failed to create product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := httphelper.RespondJSON(w, http.StatusCreated, map[string]string{"message": "created"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
