package auth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateToken(userID uint64) (string, error) {
	perm := jwt.MapClaims{}

	perm["authorized"] = true
	perm["exp"] = time.Now().Add(time.Hour * 6).Unix()
	perm["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perm)

	return token.SignedString([]byte("minhatestesupersecreta"))
}

// GetUserID return the userId that is saved in token
func GetUserByID(c echo.Context) (uint64, error) {
	tokenString := getToken(c)

	token, err := jwt.Parse(tokenString, returnVerificationKey)

	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userID"]), 10, 64)

		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New("Invalid token")
}

func getToken(c echo.Context) string {
	token := c.Request().Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signature method! %s", token.Header["alg"])
	}

	return []byte("minhatestesupersecreta"), nil
}
