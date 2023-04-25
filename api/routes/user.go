package routes

import (
	"auth-api/api/handlers"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo, h *handlers.UserHanlder) {
	e.GET("/me", h.GetUser)
}
