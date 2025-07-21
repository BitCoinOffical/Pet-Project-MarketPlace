package get_test

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"main.go/interfaces/http/dto"
	"main.go/internal/domain/products"
	get "main.go/usecase/products/method_get"
)

type GetByIdmockRepo struct {
	id        int
	product   products.Product
	returnErr error
}

func (m *GetByIdmockRepo) Delete(ctx context.Context, id int) error {
	return nil
}
func (m *GetByIdmockRepo) GetAll(ctx context.Context, filter *dto.ProductFilterDTO) ([]dto.ProductResponse, int, error) {
	return nil, 0, nil
}
func (m *GetByIdmockRepo) GetByIdRepo(ctx context.Context, id int) (*products.Product, error) {
	m.id = id
	return &m.product, m.returnErr
}

func (m *GetByIdmockRepo) Patch(ctx context.Context, id int, product *dto.ProductPatchDTO) error {
	return nil
}

func (m *GetByIdmockRepo) Put(ctx context.Context, id int, product *products.Product) error {
	return nil
}

func (m *GetByIdmockRepo) Create(ctx context.Context, p *products.Product) error {
	return nil
}

func TestGetByIdUseCase(t *testing.T) {
	tests := []struct {
		id  int
		err error
	}{
		{1, nil},
		{2, nil},
		{3, nil},
		{4, nil},
		{5, nil},
	}
	for _, i := range tests {
		t.Run("id_"+strconv.Itoa(i.id), func(t *testing.T) {
			mock := &GetByIdmockRepo{returnErr: i.err}
			uc := get.NewGetByIdProductUseCase(mock)
			_, err := uc.GetById(context.Background(), i.id)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if mock.id != i.id {
				t.Errorf("expected id %d, got %d", i.id, mock.id)
			}
		})
	}
}
func TestGetByIdUseCase_error(t *testing.T) {
	tests := []struct {
		id  int
		err error
	}{
		{1, errors.New("id not found")},
		{2, errors.New("id not found")},
		{3, errors.New("id not found")},
		{4, errors.New("id not found")},
		{5, errors.New("id not found")},
	}
	for _, i := range tests {
		t.Run("id_"+strconv.Itoa(i.id), func(t *testing.T) {
			mock := &GetByIdmockRepo{returnErr: i.err}
			uc := get.NewGetByIdProductUseCase(mock)
			_, err := uc.GetById(context.Background(), i.id)
			if err == nil {
				t.Error("expected error, got nil")
			}
		})
	}
}
