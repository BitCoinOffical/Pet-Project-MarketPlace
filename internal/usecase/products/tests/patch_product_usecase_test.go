package usecase_test

import (
	"context"
	"errors"
	"myapp/internal/interfaces/http/dto"
	usecase "myapp/internal/usecase/products"
	"testing"
)

func TestPatchUseCase(t *testing.T) {
	tests := []struct {
		id  int
		dto *dto.ProductPatchDTO
		err error
	}{
		{1, &dto.ProductPatchDTO{Name: StringPtr("Молоко"), Category: StringPtr("Продукты"), Price: Float64Ptr(80), InStock: BoolPtr(true)}, nil},
		{2, &dto.ProductPatchDTO{Name: StringPtr("Сыр"), Category: StringPtr("Продукты"), Price: Float64Ptr(250), InStock: BoolPtr(true)}, nil},
		{3, &dto.ProductPatchDTO{Name: StringPtr("Шампунь"), Category: StringPtr("Гигиена"), Price: Float64Ptr(190.5), InStock: BoolPtr(false)}, nil},
		{4, &dto.ProductPatchDTO{Name: StringPtr("Ноутбук"), Category: StringPtr("Электроника"), Price: Float64Ptr(59999.99), InStock: BoolPtr(true)}, nil},
		{5, &dto.ProductPatchDTO{Name: StringPtr("Книга"), Category: StringPtr("Канцелярия"), Price: Float64Ptr(350), InStock: BoolPtr(true)}, nil},
	}
	for _, i := range tests {
		mock := &Mockrepo{returnErr: i.err}
		uc := usecase.NewPatchProductUseCase(mock)
		err := uc.PatchProductUseCase(context.Background(), i.id, i.dto)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if mock.id != i.id {
			t.Errorf("expected id %d, got %d", i.id, mock.id)
		}
	}
}

func TestPatchUseCase_error(t *testing.T) {
	tests := []struct {
		id  int
		dto *dto.ProductPatchDTO
		err error
	}{
		{1, &dto.ProductPatchDTO{Name: StringPtr("Молоко"), Category: StringPtr("Продукты"), Price: Float64Ptr(80), InStock: BoolPtr(true)}, errors.New("not found")},
		{2, &dto.ProductPatchDTO{Name: StringPtr("Сыр"), Category: StringPtr("Продукты"), Price: Float64Ptr(250), InStock: BoolPtr(true)}, errors.New("not found")},
		{3, &dto.ProductPatchDTO{Name: StringPtr("Шампунь"), Category: StringPtr("Гигиена"), Price: Float64Ptr(190.5), InStock: BoolPtr(false)}, errors.New("not found")},
		{4, &dto.ProductPatchDTO{Name: StringPtr("Ноутбук"), Category: StringPtr("Электроника"), Price: Float64Ptr(59999.99), InStock: BoolPtr(true)}, errors.New("not found")},
		{5, &dto.ProductPatchDTO{Name: StringPtr("Книга"), Category: StringPtr("Канцелярия"), Price: Float64Ptr(350), InStock: BoolPtr(true)}, errors.New("not found")},
	}
	for _, i := range tests {

		mock := &Mockrepo{returnErr: i.err}
		uc := usecase.NewPatchProductUseCase(mock)
		err := uc.PatchProductUseCase(context.Background(), i.id, i.dto)
		if err == nil {
			t.Error("expected error, got nil")
		}
		if mock.id != i.id {
			t.Errorf("expected id %d, got %d", i.id, mock.id)
		}

	}
}
