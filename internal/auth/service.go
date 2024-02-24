package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
}
