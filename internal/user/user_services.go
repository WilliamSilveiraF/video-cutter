package user

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}

func LoginUser(email, password string) (bool, error) {
	user, err := RetrieveUser(email)
	if err != nil {
		return false, err
	}

	return CheckPasswordHash(password, user.Password), nil
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