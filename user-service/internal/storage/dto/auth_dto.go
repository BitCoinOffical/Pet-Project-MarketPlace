package dto

import _ "github.com/go-playground/validator/v10"

type LoginWithPassword struct {
	UserName string `json:"username" validate:"required,min=2,max=16"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginWithEmailCode struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required,len=6,numeric"`
}
