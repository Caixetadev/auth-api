package security

import "golang.org/x/crypto/bcrypt"

// Hash receive a string and put a hash it
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(passwordHash, passowrd string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passowrd))
}
