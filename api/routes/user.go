package routes

import (
	"auth-api/api/handlers"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo) {
	e.GET("/me", handlers.GetUser)
}
