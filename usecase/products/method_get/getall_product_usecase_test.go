package get_test

import (
	"context"
	"errors"
	"testing"

	"main.go/interfaces/http/dto"
	"main.go/internal/domain/products"
	get "main.go/usecase/products/method_get"
)

type mockRepo struct {
	filter  *dto.ProductFilterDTO
	filters []dto.ProductResponse
	num     int
	err     error
}

func (m *mockRepo) Delete(ctx context.Context, id int) error {
	return nil
}
func (m *mockRepo) GetAll(ctx context.Context, filter *dto.ProductFilterDTO) ([]dto.ProductResponse, int, error) {
	m.filter = filter
	return m.filters, m.num, m.err
}
func (m *mockRepo) GetByIdRepo(ctx context.Context, id int) (*products.Product, error) {
	return nil, nil
}

func (m *mockRepo) Patch(ctx context.Context, id int, product *dto.ProductPatchDTO) error {
	return nil
}

func (m *mockRepo) Put(ctx context.Context, id int, product *products.Product) error {
	return nil
}

func (m *mockRepo) Create(ctx context.Context, p *products.Product) error {
	return nil
}
func StringPtr(s string) *string {
	return &s
}

func Float64Ptr(f float64) *float64 {
	return &f
}
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
		mock := &mockRepo{err: i.err}
		uc := get.NewGetAllUseCase(mock)
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
		mock := &mockRepo{err: i.err}
		uc := get.NewGetAllUseCase(mock)
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
