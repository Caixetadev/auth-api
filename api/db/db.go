package db

import (
	"auth-api/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connect opens the database connection and returns it
func Connect() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@db/%s?sslmode=disable", config.EnvConfigs.PostgresUser, config.EnvConfigs.PostgresPassword, config.EnvConfigs.PostgresDB))

	if err != nil {
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		db.Close()
		panic(err.Error())
	}

	return db
}
