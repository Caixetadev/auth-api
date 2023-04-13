package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userID uint64) (string, error) {
	perm := jwt.MapClaims{}

	perm["authorized"] = true
	perm["exp"] = time.Now().Add(time.Hour * 6).Unix()
	perm["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perm)

	return token.SignedString([]byte("minhatestesupersecreta"))
}
