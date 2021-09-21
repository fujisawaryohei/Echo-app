package utils

import "golang.org/x/crypto/bcrypt"

func Hashed(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword)
}
