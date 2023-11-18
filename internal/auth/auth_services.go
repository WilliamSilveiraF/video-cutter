package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}

func RegisterUser(email, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	newUser := User{
		Email:    email,
		Password: hashedPassword,
	}

	return InsertUser(newUser)
}