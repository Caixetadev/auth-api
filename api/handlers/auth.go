package handlers

import (
	"auth-api/api/auth"
	"auth-api/api/models"
	"auth-api/api/repository"
	"auth-api/api/security"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHanlder struct {
	db *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHanlder {
	return &AuthHanlder{db: db}
}

// Register registers a new user
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User information"
// @Success 201 {object} map[string]string
// @Router /register [post]
func (h *AuthHanlder) Register(c echo.Context) error {
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

	repository := repository.NewRepositoryOfAuth(h.db)

	if err := repository.Create(user); err != nil {
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
func (h *AuthHanlder) Login(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	repository := repository.NewRepositoryOfAuth(h.db)

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
