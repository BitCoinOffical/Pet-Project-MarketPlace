package dto

import _ "github.com/go-playground/validator/v10"

type Register struct {
	Name     string `json:"name" validate:"required,min=2,max=16"`
	UserName string `json:"username" validate:"required,min=2,max=16"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
