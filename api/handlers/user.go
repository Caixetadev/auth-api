package handlers

import (
	"auth-api/api/auth"
	"auth-api/api/db"
	"auth-api/api/models"
	"auth-api/api/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUser returns current user
// @Summary Get current user
// @Description Get current user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.UserResponse
// @Failure 404
// @Router /me [get]
func GetUser(c echo.Context) error {
	userID, err := auth.GetUserByID(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	db, err := db.Connect()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	defer db.Close()

	repository := repository.NewRepositoryOfAuth(db)

	user, err := repository.GetUserByID(userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	userResponse := models.UserResponse{
		Email:    user.Email,
		LastName: user.LastName,
		Name:     user.Name,
	}

	return c.JSON(http.StatusOK, userResponse)
}
