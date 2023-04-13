package models

import "github.com/go-playground/validator"

type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name" validate:"required"`
	LastNmae string `json:"lastname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserValidator struct {
	Validator *validator.Validate
}

func (u *UserValidator) Validate(i interface{}) error {
	return u.Validator.Struct(i)
}
