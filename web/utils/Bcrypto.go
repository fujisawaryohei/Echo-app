package utils

import "golang.org/x/crypto/bcrypt"

func Hashed(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword)
}

func Compare(hashedPassword string, inputPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	if err != nil {
		return err
	}
	return nil
}
