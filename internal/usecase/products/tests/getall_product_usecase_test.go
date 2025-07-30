package usecase_test

import (
	"context"
	"errors"
	"myapp/internal/interfaces/http/dto"
	usecase "myapp/internal/usecase/products"
	"testing"
)

func TestGetAllUseCase(t *testing.T) {
	tests := []struct {
		filter dto.ProductFilterDTO
		err    error
	}{
		{
			filter: dto.ProductFilterDTO{
				Category: nil,
				MinPrice: nil,
				MaxPrice: nil,
				Search:   nil,
				Page:     1,
			},
			err: nil,
		},
		{
			filter: dto.ProductFilterDTO{
				Category: StringPtr("electronics"),
				MinPrice: Float64Ptr(100.0),
				MaxPrice: Float64Ptr(500.0),
				Search:   StringPtr("smartphone"),
				Page:     2,
			},
			err: nil,
		},
		{
			filter: dto.ProductFilterDTO{
				Category: StringPtr(""),
				MinPrice: Float64Ptr(-10.0),
				MaxPrice: Float64Ptr(0),
				Search:   StringPtr(""),
				Page:     0,
			},
			err: nil,
		},
	}
	for _, i := range tests {
		mock := &Mockrepo{err: i.err}
		uc := usecase.NewGetAllUseCase(mock)
		var category, search string
		var minPrice, maxPrice float64

		if i.filter.Category != nil {
			category = *i.filter.Category
		}
		if i.filter.Search != nil {
			search = *i.filter.Search
		}
		if i.filter.MinPrice != nil {
			minPrice = *i.filter.MinPrice
		}
		if i.filter.MaxPrice != nil {
			maxPrice = *i.filter.MaxPrice
		}
		_, _, err := uc.GetAll(context.Background(), category, search, i.filter.Page, minPrice, maxPrice)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

func TestGetAllUseCase_error(t *testing.T) {
	tests := []struct {
		filter dto.ProductFilterDTO
		err    error
	}{
		{
			filter: dto.ProductFilterDTO{
				Category: nil,
				MinPrice: nil,
				MaxPrice: nil,
				Search:   nil,
				Page:     1,
			},
			err: errors.New("error"),
		},
		{
			filter: dto.ProductFilterDTO{
				Category: StringPtr("electronics"),
				MinPrice: Float64Ptr(100.0),
				MaxPrice: Float64Ptr(500.0),
				Search:   StringPtr("smartphone"),
				Page:     2,
			},
			err: errors.New("error"),
		},
		{
			filter: dto.ProductFilterDTO{
				Category: StringPtr(""),
				MinPrice: Float64Ptr(-10.0),
				MaxPrice: Float64Ptr(0),
				Search:   StringPtr(""),
				Page:     0,
			},
			err: errors.New("error"),
		},
	}
	for _, i := range tests {
		mock := &Mockrepo{err: i.err}
		uc := usecase.NewGetAllUseCase(mock)
		var category, search string
		var minPrice, maxPrice float64

		if i.filter.Category != nil {
			category = *i.filter.Category
		}
		if i.filter.Search != nil {
			search = *i.filter.Search
		}
		if i.filter.MinPrice != nil {
			minPrice = *i.filter.MinPrice
		}
		if i.filter.MaxPrice != nil {
			maxPrice = *i.filter.MaxPrice
		}
		_, _, err := uc.GetAll(context.Background(), category, search, i.filter.Page, minPrice, maxPrice)
		if err == nil {
			t.Errorf("unexpected error: %v", err)
		}
	}
}
