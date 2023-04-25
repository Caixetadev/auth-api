package handlers

import (
	"auth-api/api/auth"
	"auth-api/api/models"
	"auth-api/api/repository"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHanlder struct {
	db *sql.DB
}

func NewUserHanlder(db *sql.DB) *UserHanlder {
	return &UserHanlder{db: db}
}

// GetUser returns current user
// @Summary Get current user
// @Description Get current user
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.UserResponse
// @Failure 404
// @Router /me [get]
func (h *UserHanlder) GetUser(c echo.Context) error {
	userID, err := auth.GetUserByID(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	repository := repository.NewRepositoryOfAuth(h.db)

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
