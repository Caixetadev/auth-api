package routes

import (
	"auth-api/api/handlers"

	"github.com/labstack/echo/v4"
)

func InitAuthRoutes(e *echo.Echo, h *handlers.AuthHanlder) {
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
}
