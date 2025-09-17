package handlers_test

import (
	"context"
	"myapp/internal/interfaces/http/dto"

	"myapp/internal/domain/products"
)

type mockRepo struct {
	filter          *dto.ProductFilterDTO
	filters         []dto.ProductResponse
	pr              *products.Product
	ProductPatchDTO *dto.ProductPatchDTO
	num             int
	err             error
	CalledWithID    int
	returnErr       error
}

func (m *mockRepo) Delete(ctx context.Context, id int) error {
	m.CalledWithID = id
	return m.returnErr
}
func (m *mockRepo) GetAll(ctx context.Context, filter *dto.ProductFilterDTO) ([]dto.ProductResponse, int, error) {
	m.filter = filter
	return m.filters, len(m.filters), m.err
}
func (m *mockRepo) GetByIdRepo(ctx context.Context, id int) (*products.Product, error) {
	m.num = id
	return m.pr, m.returnErr
}

func (m *mockRepo) Patch(ctx context.Context, id int, product *dto.ProductPatchDTO) error {
	m.num = id
	m.ProductPatchDTO = product
	return m.returnErr
}

func (m *mockRepo) Put(ctx context.Context, id int, product *products.Product) error {
	m.num = id
	m.pr = product
	return m.returnErr
}

func (m *mockRepo) Create(ctx context.Context, p *products.Product) error {
	m.pr = p
	return m.returnErr
}
func StringPtr(s string) *string {
	return &s
}

func Float64Ptr(f float64) *float64 {
	return &f
}
