package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Connect opens the database connection and returns it
func Connect() (*sql.DB, error) {
	// config, err := config.LoadConfig(".")

	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }

	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@db/auth-api?sslmode=disable")

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
