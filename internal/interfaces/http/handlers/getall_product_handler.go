package handlers

import (
	"myapp/internal/interfaces/http/dto"
	"net/http"
	"strconv"

	usecase "myapp/internal/usecase/products"
	"myapp/pkg/httphelper"
)

type GetAllHandler struct {
	usecase *usecase.UseCases
}

func NewGetAllHandler(usecase *usecase.UseCases) *GetAllHandler {
	return &GetAllHandler{usecase: usecase}
}

func (uc *GetAllHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	var min_price, max_price float64
	var page = 1

	category := r.URL.Query().Get("category")
	search := r.URL.Query().Get("search")

	min_priceSTR := r.URL.Query().Get("min_price")
	if min_priceSTR != "" {
		var err error
		min_price, err = strconv.ParseFloat(min_priceSTR, 64)
		if err != nil {
			http.Error(w, "Invalid min price: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	max_priceSTR := r.URL.Query().Get("max_price")
	if max_priceSTR != "" {
		var err error
		max_price, err = strconv.ParseFloat(max_priceSTR, 64)
		if err != nil {
			http.Error(w, "Invalid max price: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	pagestr := r.URL.Query().Get("page")
	if pagestr != "" {
		var err error
		page, err = strconv.Atoi(pagestr)
		if err != nil {
			http.Error(w, "Invalid page: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	products, total_count, err := uc.usecase.GetAll.GetAll(r.Context(), category, search, page, min_price, max_price)
	if err != nil {
		http.Error(w, "Failed get products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res := &dto.ProductList{
		Items:      products,
		TotalCount: total_count,
	}

	if err := httphelper.RespondJSON(w, http.StatusOK, res); err != nil {
		http.Error(w, "Failed get json witch products: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
