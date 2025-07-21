package post_test

import (
	"context"
	"errors"
	"testing"

	"main.go/interfaces/http/dto"
	"main.go/internal/domain/products"
	post "main.go/usecase/products/method_post"
)

type mockRepo struct {
	input     *products.Product
	returnErr error
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
	return nil
}

func (m *mockRepo) Create(ctx context.Context, p *products.Product) error {
	m.input = p
	return m.returnErr
}

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
		mock := &mockRepo{returnErr: i.err}

		uc := post.NewPostProductUseCase(mock)

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
		mock := &mockRepo{returnErr: i.err}
		uc := post.NewPostProductUseCase(mock)

		err := uc.Execute(context.Background(), i.dto)
		if err == nil {
			t.Error("expected error, got nil")
		}
	}
}
