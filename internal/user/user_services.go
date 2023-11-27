package user

import (
	"time"
	"os"
	"errors"
    "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	
)

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

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

func RegisterUser(email, password string) (int, error) {
    hashedPassword, err := HashPassword(password)
    if err != nil {
        return 0, err
    }

    newUser := User{
        Email:    email,
        Password: hashedPassword,
    }

    userID, err := InsertUser(newUser)
    if err != nil {
        return 0, err
    }

    return userID, nil
}

func UpdateUserPassword(email, oldPassword, newPassword string) error {
	user, err := RetrieveUser(email)
	if err != nil {
		return err
	}

	if !CheckPasswordHash(oldPassword, user.Password) {
		return errors.New("old password does not match")
	}

	hashedNewPassword, err := HashPassword(newPassword)
	if err != nil {
		return err
	}

	return UpdatePassword(email, hashedNewPassword)
}