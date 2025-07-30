package usecase_test

import (
	"context"
	"errors"
	usecase "myapp/internal/usecase/products"
	"strconv"
	"testing"
)

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
			mock := &Mockrepo{returnErr: i.err}
			uc := usecase.NewGetByIdProductUseCase(mock)
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
			mock := &Mockrepo{returnErr: i.err}
			uc := usecase.NewGetByIdProductUseCase(mock)
			_, err := uc.GetById(context.Background(), i.id)
			if err == nil {
				t.Error("expected error, got nil")
			}
		})
	}
}
