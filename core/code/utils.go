package code

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	expiration = time.Duration(
		int64(jwtConfig.Int("expirationTimeSeconds"))) * time.Second
)

// Claims stores claims of verification token
type Claims struct {
	UserID                     string
	CreationUTC, ExpirationUTC int64
	jwt.StandardClaims
}

// claims creates claims for verification token.
func claims(userID string) Claims {
	return Claims{
		UserID:        userID,
		CreationUTC:   time.Now().Unix(),
		ExpirationUTC: time.Now().Add(expiration).Unix(),
	}
}

// jwtKeyFunc signs a token.
func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return secret, nil
}
