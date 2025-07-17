package dto

import _ "github.com/go-playground/validator/v10"

type ProductPatchDTO struct {
	Name     *string  `json:"name" validate:"omitempty,min=2,max=255"`
	Category *string  `json:"category" validate:"omitempty,min=2,max=32"`
	Price    *float64 `json:"price" validate:"omitempty,gte=1,lte=10000000000"`
	InStock  *bool    `json:"in_stock" validate:"omitempty"`
}
