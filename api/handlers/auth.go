package handlers

import (
	"auth-api/api/db"
	"auth-api/api/models"
	"auth-api/api/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := user.Prepare("REGISTER"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	db, err := db.Connect()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	defer db.Close()

	repository := repository.NewRepositoryOfAuth(db)

	if err = repository.Create(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}
