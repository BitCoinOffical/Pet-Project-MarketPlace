package usecase_test

import (
	"context"
	"myapp/internal/domain/products"
	"myapp/internal/interfaces/http/dto"
)

type Mockrepo struct {
	id           int
	returnErr    error
	filter       *dto.ProductFilterDTO
	filters      []dto.ProductResponse
	patchproduct *dto.ProductPatchDTO
	num          int
	err          error
	product      products.Product
	input        *products.Product
}

func (m *Mockrepo) Delete(ctx context.Context, id int) error {
	m.id = id
	return m.returnErr
}
func (m *Mockrepo) GetAll(ctx context.Context, filter *dto.ProductFilterDTO) ([]dto.ProductResponse, int, error) {
	m.filter = filter
	return m.filters, m.num, m.err
}
func (m *Mockrepo) GetByIdRepo(ctx context.Context, id int) (*products.Product, error) {
	m.id = id
	return &m.product, m.returnErr
}

func (m *Mockrepo) Patch(ctx context.Context, id int, product *dto.ProductPatchDTO) error {
	m.id = id
	m.patchproduct = product
	return m.returnErr
}

func (m *Mockrepo) Put(ctx context.Context, id int, product *products.Product) error {
	m.id = id
	m.product = *product
	return m.returnErr
}

func (m *Mockrepo) Create(ctx context.Context, p *products.Product) error {
	m.input = p
	return m.returnErr
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
