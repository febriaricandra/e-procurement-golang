package middleware

import (
	"errors"
	"fmt"
	"my-procurement-system/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// internal/middleware/auth.go
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		token, err := utils.ParseAndValidateToken(tokenString)
		if err != nil {
			handleTokenError(c, err)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// Validasi tipe token
		if claims["type"] != "access" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token type"})
			return
		}

		c.Set("userID", uint(claims["user_id"].(float64)))
		c.Next()
	}
}

func extractToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header required")
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return "", fmt.Errorf("invalid token format")
	}

	return splitToken[1], nil
}

func handleTokenError(c *gin.Context, err error) {
	if errors.Is(err, jwt.ErrTokenMalformed) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "malformed token"})
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired or not active"})
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
	}
}
