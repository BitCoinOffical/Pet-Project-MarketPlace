package usecase_test

import (
	"context"
	"errors"
	"myapp/internal/interfaces/http/dto"
	usecase "myapp/internal/usecase/products"
	"testing"
)

func TestPostProductUseCase(t *testing.T) {

	tests := []struct {
		dto dto.ProductCreateDTO
		err error
	}{
		{dto.ProductCreateDTO{Name: "Молоко", Category: "Продукты", Price: 80, InStock: true}, nil},
		{dto.ProductCreateDTO{Name: "Сыр", Category: "Продукты", Price: 250, InStock: true}, nil},
		{dto.ProductCreateDTO{Name: "Шампунь", Category: "Гигиена", Price: 190.5, InStock: false}, nil},
		{dto.ProductCreateDTO{Name: "Ноутбук", Category: "Электроника", Price: 59999.99, InStock: true}, nil},
		{dto.ProductCreateDTO{Name: "Книга", Category: "Канцелярия", Price: 350, InStock: true}, nil},
	}
	for _, i := range tests {
		mock := &Mockrepo{returnErr: i.err}

		uc := usecase.NewPostProductUseCase(mock)

		err := uc.Execute(context.Background(), i.dto)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if mock.input.Name != i.dto.Name {
			t.Errorf("expected name %q, got %q", i.dto.Name, mock.input.Name)
		}

		if mock.input.Category != i.dto.Category {
			t.Errorf("expected category %q, got %q", i.dto.Category, mock.input.Category)
		}

		if mock.input.Price != i.dto.Price {
			t.Errorf("expected price %v, got %v", i.dto.Price, mock.input.Price)
		}

		if mock.input.InStock != i.dto.InStock {
			t.Errorf("expected InStock %v, got %v", i.dto.InStock, mock.input.InStock)
		}
	}

}

func TestPostProductUseCase_error(t *testing.T) {

	tests := []struct {
		dto dto.ProductCreateDTO
		err error
	}{
		{dto.ProductCreateDTO{Name: "Молоко", Category: "Продукты", Price: 80, InStock: true}, errors.New("insert error")},
		{dto.ProductCreateDTO{Name: "Сыр", Category: "Продукты", Price: 250, InStock: true}, errors.New("insert error")},
		{dto.ProductCreateDTO{Name: "Шампунь", Category: "Гигиена", Price: 190.5, InStock: false}, errors.New("insert error")},
		{dto.ProductCreateDTO{Name: "Ноутбук", Category: "Электроника", Price: 59999.99, InStock: true}, errors.New("insert error")},
		{dto.ProductCreateDTO{Name: "Книга", Category: "Канцелярия", Price: 350, InStock: true}, errors.New("insert error")},
	}
	for _, i := range tests {
		mock := &Mockrepo{returnErr: i.err}
		uc := usecase.NewPostProductUseCase(mock)

		err := uc.Execute(context.Background(), i.dto)
		if err == nil {
			t.Error("expected error, got nil")
		}
	}
}
