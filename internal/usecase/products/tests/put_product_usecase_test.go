package usecase_test

import (
	"context"
	"errors"
	"myapp/internal/interfaces/http/dto"
	usecase "myapp/internal/usecase/products"
	"testing"
)

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
			mock := &Mockrepo{returnErr: i.err}
			uc := usecase.NewPutByIdProductUseCase(mock)
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
			mock := &Mockrepo{returnErr: i.err}
			uc := usecase.NewPutByIdProductUseCase(mock)
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
