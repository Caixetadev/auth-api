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
	_, err := u.db.Exec("INSERT INTO users (name, lastname, email, password) VALUES ($1, $2, $3, $4)", user.Name, user.LastName, user.Email, user.Password)
	return err
}

func (u User) GetUserByEmail(email string) (models.User, error) {
	line, err := u.db.Query("SELECT id, password FROM users WHERE email = $1", email)

	if err != nil {
		return models.User{}, err
	}

	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
