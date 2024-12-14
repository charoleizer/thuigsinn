package pkg

import (
	"errors"
	"regexp"
)

// IsValidEmail checks if the email is valid using a regular expression
func IsValidEmail(email string) error {
	var re = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("email must contain only lowercase letters, numbers, and special characters ._%+-, and have a valid domain")
	}

	return nil
}

// IsValidPassword checks if the password is valid (minimum 8 characters, at least one letter and one number)
func IsValidPassword(password string) error {
	// Checks if it is at least 8 characters long
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	// Checks if it contains at least one letter and one number
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)
	if !hasUppercase || !hasLowercase || !hasDigit || !hasSpecial {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}

	return nil
}

// IsValidUsername checks if the username is valid (only letters, numbers, underscores, between 3 and 16 characters)
func IsValidUsername(username string) error {
	var re = regexp.MustCompile(`^[a-zA-Z0-9_]{3,16}$`)
	if !re.MatchString(username) {
		return errors.New("username must be between 3 and 16 characters long and contain only letters, numbers, and underscores")
	}

	return nil
}
