package main

import (
	"auth-api/api/models"
	"auth-api/api/routes"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "auth-api/docs"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()
	v := validator.New()

	e.Validator = &models.UserValidator{Validator: v}

	routes.InitAuthRoutes(e)
	routes.InitUserRoutes(e)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":3333"))
}
