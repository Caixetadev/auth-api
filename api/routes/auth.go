package routes

import (
	"auth-api/api/handlers"

	"github.com/labstack/echo/v4"
)

func InitAuthRoutes(e *echo.Echo) {
	e.POST("/register", handlers.Register)
}
