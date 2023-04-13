package models

import (
	"auth-api/api/security"

	"github.com/go-playground/validator"
)

type User struct {
	ID       uint64 `json:"id,omitempty"`
	Name     string `json:"name" validate:"required"`
	LastName string `json:"lastname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserValidator struct {
	Validator *validator.Validate
}

func (u *UserValidator) Validate(i interface{}) error {
	return u.Validator.Struct(i)
}

func (u *User) Prepare(step string) error {
	if step == "REGISTER" {
		passwordHash, err := security.Hash(u.Password)

		if err != nil {
			return err
		}

		u.Password = string(passwordHash)
	}

	return nil
}
