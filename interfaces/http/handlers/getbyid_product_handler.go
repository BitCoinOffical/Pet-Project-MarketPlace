package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"main.go/pkg/httphelper"
	usecase "main.go/usecase/products"
)

type GetByIdProductHandler struct {
	usecase *usecase.UseCases
}

func NewGetByIdProductHandler(usecase *usecase.UseCases) *GetByIdProductHandler {
	return &GetByIdProductHandler{usecase: usecase}
}

func (h *GetByIdProductHandler) GetByIdProductHandler(w http.ResponseWriter, r *http.Request) {
	idstr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.usecase.GetById.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to get product: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := httphelper.RespondJSON(w, http.StatusOK, res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
