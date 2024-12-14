package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password, salt string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword, password, salt string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt))
	if err != nil {
		return err
	}

	return nil
}
