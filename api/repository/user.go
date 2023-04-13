package repository

import (
	"auth-api/api/models"
	"database/sql"
)

type User struct {
	db *sql.DB
}

func NewRepositoryOfAuth(db *sql.DB) *User {
	return &User{db}
}

func (u User) Create(user models.User) error {
	_, err := u.db.Exec("INSERT INTO users (name, lastname, email, password) VALUES ($1, $2, $3, $4)", user.Name, user.LastNmae, user.Email, user.Password)
	return err
}
