package auth

import (
	"github.com/aenmurtic/be-hijooin-admin/internal/user"
	"github.com/golang-jwt/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer " // Standard prefix
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}

		tokenString := authHeader[len(BearerSchema):] // Strip prefix

		token, err := ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["id"].(float64))

		var existingUser user.User
		if result := db.First(&existingUser, userID); result.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}

		c.Set("userID", userID) // Set in Gin context for access in route handlers
		c.Next()
	}
}
