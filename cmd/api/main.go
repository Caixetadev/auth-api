package main

import (
	"auth-api/api/db"
	"auth-api/api/handlers"
	"auth-api/api/models"
	"auth-api/api/routes"
	"auth-api/config"
	"database/sql"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "auth-api/docs"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// @title Auth API
// @version 1.0
// @description Authentication API Documentation.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()
	v := validator.New()

	config.InitEnvConfigs()

	e.Validator = &models.UserValidator{Validator: v}

	db := db.Connect()

	defer db.Close()

	Setup(db, e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":3333"))
}

func Setup(db *sql.DB, e *echo.Echo) {
	handlerAuth, handlerUser := InitHandlers(db)

	routes.InitAuthRoutes(e, handlerAuth)
	routes.InitUserRoutes(e, handlerUser)
}

func InitHandlers(db *sql.DB) (*handlers.AuthHanlder, *handlers.UserHanlder) {
	handlerAuth := handlers.NewAuthHandler(db)
	hanlderUser := handlers.NewUserHanlder(db)

	return handlerAuth, hanlderUser
}
