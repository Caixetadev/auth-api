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

// Register registers a new user
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User information"
// @Success 201 {object} map[string]string
// @Router /register [post]
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

	response := map[string]string{
		"message": "Usuário criado com sucesso",
	}

	return c.JSON(http.StatusCreated, response)
}

// Login logs in the user
// @Summary Log in the user
// @Description Log in the user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.LoginPayload true "User's credentials"
// @Success 200 {object} map[string]string
// @Failure 400
// @Router /login [post]
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
		"message": "Logado com sucesso",
		"token":   token,
	}

	return c.JSON(http.StatusOK, response)
}
