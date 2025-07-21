package delete_test

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"main.go/interfaces/http/dto"
	"main.go/internal/domain/products"
	delete "main.go/usecase/products/method_delete"
)

type mockRepo struct {
	id        int
	returnErr error
}

func (m *mockRepo) Delete(ctx context.Context, id int) error {
	m.id = id
	return m.returnErr
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
	return nil
}

func TestDeleteByIdUseCase(t *testing.T) {
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
			mock := &mockRepo{returnErr: i.err}
			uc := delete.NewDeletedByIDProductUseCase(mock)

			err := uc.DeletedByID(context.Background(), i.id)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if mock.id != i.id {
				t.Errorf("expected id %d, got %d", i.id, mock.id)
			}
		})
	}
}

func TestDeleteByIdUseCase_erorr(t *testing.T) {
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
			mock := &mockRepo{returnErr: i.err}
			uc := delete.NewDeletedByIDProductUseCase(mock)

			err := uc.DeletedByID(context.Background(), i.id)
			if err == nil {
				t.Error("expected error, got nil")
			}
		})
	}
}
