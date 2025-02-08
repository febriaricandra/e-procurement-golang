package utils

import (
	"fmt"
	"my-procurement-system/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(getJWTSecret())

func getJWTSecret() string {
	return config.JWTSecret()
}

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
		"iss":     "e-procurement",
		"aud":     "user",
		"type":    "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iss":     "your-app-name",
		"aud":     "user",
		"type":    "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseAndValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validasi signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Validasi claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Validasi issuer
			if claims["iss"] != "e-procurement" {
				return nil, fmt.Errorf("invalid issuer")
			}

			// Validasi audience
			if claims["aud"] != "user" {
				return nil, fmt.Errorf("invalid audience")
			}
		}

		return jwtSecret, nil
	})
}
