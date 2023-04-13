package handlers

import (
	"auth-api/api/auth"
	"auth-api/api/db"
	"auth-api/api/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

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

	return c.JSON(http.StatusOK, user)
}
