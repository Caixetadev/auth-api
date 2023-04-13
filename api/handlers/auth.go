package handlers

import (
	"auth-api/api/auth"
	"auth-api/api/db"
	"auth-api/api/models"
	"auth-api/api/repository"
	"auth-api/api/security"
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

func Login(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := db.Connect()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	defer db.Close()

	repository := repository.NewRepositoryOfAuth(db)

	userSavedInBank, err := repository.GetUserByEmail(user.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = security.VerifyPassword(userSavedInBank.Password, user.Password); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	token, err := auth.CreateToken(userSavedInBank.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]string{
		"token": token,
	}

	return c.JSON(http.StatusOK, response)
}
