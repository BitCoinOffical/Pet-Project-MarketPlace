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

type PatchProductHandler struct {
	usecase *usecase.UseCases
}

func NewPatchProductHandler(usecase *usecase.UseCases) *PatchProductHandler {
	return &PatchProductHandler{usecase: usecase}
}

func (uc *PatchProductHandler) PatchProductHandler(w http.ResponseWriter, r *http.Request) {
	var ProductPatchDTO dto.ProductPatchDTO
	var validate = validator.New()
	idstr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&ProductPatchDTO); err != nil {
		http.Error(w, "Failed decode json: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := validate.Struct(ProductPatchDTO); err != nil {
		http.Error(w, "validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := uc.usecase.PatchById.PatchProductUseCase(r.Context(), id, &ProductPatchDTO); err != nil {
		http.Error(w, "Patch failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := httphelper.RespondJSON(w, http.StatusOK, map[string]string{"message": "patched"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
