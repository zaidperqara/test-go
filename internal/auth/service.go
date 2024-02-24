package auth

import (
	"errors"
	"os"
	"time"

	"github.com/aenmurtic/be-hijooin-admin/internal/user"
	"github.com/golang-jwt/jwt"
)

// Replace with a stronger secret
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID uint `json:"id"`
	// ... Add other relevant claims
	jwt.StandardClaims
}

func GenerateToken(user *user.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24) // Example expiration

	claims := Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	}), nil
}
