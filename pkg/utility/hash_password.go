package utility

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	var generatedPassword string
	if password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		generatedPassword = string(hashedPassword)
	}

	return generatedPassword
}

func ComparePassword(hashedPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}
