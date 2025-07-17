package dto

import _ "github.com/go-playground/validator/v10"

type ProductCreateDTO struct {
	Name     string  `json:"name" validate:"required,min=2,max=255,alphanum"`
	Category string  `json:"category" validate:"required,min=2,max=32"`
	Price    float64 `json:"price" validate:"required,gte=1,lte=10000000000"`
	InStock  bool    `json:"in_stock"`
}
