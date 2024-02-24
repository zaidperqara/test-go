package common

import (
	"time"

	"github.com/aenmurtic/be-hijooin-admin/internal/user"
	"github.com/golang-jwt/jwt"
	"os"
)

type UserID uint

type Claims struct {
	UserID UserID `json:"id"`
	jwt.StandardClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(user *user.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := Claims{
		UserID: UserID(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
