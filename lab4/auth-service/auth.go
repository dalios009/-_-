package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = "your-256-bit-secret" // This should be an environment variable in production

// Claims structure to store JWT claims
type Claims struct {
	UserID string `json:"sub"`
	jwt.RegisteredClaims
}

// Function to create a JWT token for a customer
func GenerateJWT(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// Function to validate JWT token
func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}
