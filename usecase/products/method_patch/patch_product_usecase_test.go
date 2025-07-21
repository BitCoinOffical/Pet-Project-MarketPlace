package patch_test

import (
	"context"
	"errors"
	"testing"

	"main.go/interfaces/http/dto"
	"main.go/internal/domain/products"
	patch "main.go/usecase/products/method_patch"
)

type mockRepo struct {
	id        int
	returnErr error
	product   dto.ProductPatchDTO
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
	m.id = id
	m.product = *product
	return m.returnErr
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

func BoolPtr(b bool) *bool {
	return &b
}
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
		mock := &mockRepo{returnErr: i.err}
		uc := patch.NewPatchProductUseCase(mock)
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

		mock := &mockRepo{returnErr: i.err}
		uc := patch.NewPatchProductUseCase(mock)
		err := uc.PatchProductUseCase(context.Background(), i.id, i.dto)
		if err == nil {
			t.Error("expected error, got nil")
		}
		if mock.id != i.id {
			t.Errorf("expected id %d, got %d", i.id, mock.id)
		}

	}
}
