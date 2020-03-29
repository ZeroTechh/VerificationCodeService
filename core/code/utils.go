package code

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// isExpired is used to check if expirationTime has already occured
func isExpired(expirationTime time.Time) bool {
	return time.Now().After(expirationTime)
}

// createClaims is used to create claims
func createClaims(userID string) Claims {
	creationTime := time.Now()
	expirationTime := creationTime.Add(expiration)

	return Claims{
		UserID:        userID,
		CreationUTC:   creationTime.Unix(),
		ExpirationUTC: expirationTime.Unix(),
	}
}

// jwtKeyFunc is used to sign a token
func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return secret, nil
}
