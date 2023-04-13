package models

import (
	"auth-api/api/security"

	"github.com/go-playground/validator"
)

type User struct {
	ID       uint64 `json:"id,omitempty"`
	Name     string `json:"name,omitempty" validate:"required"`
	LastName string `json:"lastname,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
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
