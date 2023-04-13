package main

import (
	"auth-api/api/models"
	"auth-api/api/routes"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	v := validator.New()

	e.Validator = &models.UserValidator{Validator: v}

	routes.InitAuthRoutes(e)
	routes.InitUserRoutes(e)

	e.Logger.Fatal(e.Start(":3333"))
}
