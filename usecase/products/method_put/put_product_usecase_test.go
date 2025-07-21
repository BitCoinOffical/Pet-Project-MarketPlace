package put_test

import (
	"context"
	"errors"
	"testing"

	"main.go/interfaces/http/dto"
	"main.go/internal/domain/products"
	put "main.go/usecase/products/method_put"
)

type mockRepo struct {
	id        int
	returnErr error
	product   products.Product
}

func (m *mockRepo) Delete(ctx context.Context, id int) error {
	return nil
}
func (m *mockRepo) GetAll(ctx context.Context, filter *dto.ProductFilterDTO) ([]dto.ProductResponse, int, error) {
	return nil, 0, nil
}
func (m *mockRepo) GetByIdRepo(ctx context.Context, id int) (*products.Product, error) {
	return nil, nil
}

func (m *mockRepo) Patch(ctx context.Context, id int, product *dto.ProductPatchDTO) error {
	return nil
}

func (m *mockRepo) Put(ctx context.Context, id int, product *products.Product) error {
	m.id = id
	m.product = *product
	return m.returnErr
}

func (m *mockRepo) Create(ctx context.Context, p *products.Product) error {
	return nil
}

func TestPutUseCase(t *testing.T) {
	tests := []struct {
		id  int
		dto dto.ProductPutDTO
		err error
	}{
		{1, dto.ProductPutDTO{Name: "Молоко", Category: "Продукты", Price: 80, InStock: true}, nil},
		{2, dto.ProductPutDTO{Name: "Сыр", Category: "Продукты", Price: 250, InStock: true}, nil},
		{3, dto.ProductPutDTO{Name: "Шампунь", Category: "Гигиена", Price: 190.5, InStock: false}, nil},
		{4, dto.ProductPutDTO{Name: "Ноутбук", Category: "Электроника", Price: 59999.99, InStock: true}, nil},
		{5, dto.ProductPutDTO{Name: "Книга", Category: "Канцелярия", Price: 350, InStock: true}, nil},
	}
	for _, i := range tests {
		t.Run("testing_"+i.dto.Name, func(t *testing.T) {
			mock := &mockRepo{returnErr: i.err}
			uc := put.NewPutByIdProductUseCase(mock)
			err := uc.PutByID(context.Background(), i.id, i.dto)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if mock.id != i.id {
				t.Errorf("expected id %d, got %d", i.id, mock.id)
			}
		})
	}
}
func TestPutUseCase_error(t *testing.T) {
	tests := []struct {
		id  int
		dto dto.ProductPutDTO
		err error
	}{
		{1, dto.ProductPutDTO{Name: "Молоко", Category: "Продукты", Price: 80, InStock: true}, errors.New("not found")},
		{2, dto.ProductPutDTO{Name: "Сыр", Category: "Продукты", Price: 250, InStock: true}, errors.New("not found")},
		{3, dto.ProductPutDTO{Name: "Шампунь", Category: "Гигиена", Price: 190.5, InStock: false}, errors.New("not found")},
		{4, dto.ProductPutDTO{Name: "Ноутбук", Category: "Электроника", Price: 59999.99, InStock: true}, errors.New("not found")},
		{5, dto.ProductPutDTO{Name: "Книга", Category: "Канцелярия", Price: 350, InStock: true}, errors.New("not found")},
	}
	for _, i := range tests {
		t.Run("testing_"+i.dto.Name, func(t *testing.T) {
			mock := &mockRepo{returnErr: i.err}
			uc := put.NewPutByIdProductUseCase(mock)
			err := uc.PutByID(context.Background(), i.id, i.dto)
			if err == nil {
				t.Error("expected error, got nil")
			}
			if mock.id != i.id {
				t.Errorf("expected id %d, got %d", i.id, mock.id)
			}
		})
	}
}
