package code

import (
	"time"

	"github.com/ZeroTechh/hades"
	"github.com/dgrijalva/jwt-go"
)

var (
	config     = hades.GetConfig("main.yaml", []string{"config", "../config", "../../config"})
	jwtConfig  = config.Map("JWT")
	secret     = []byte(jwtConfig.Str("secret"))
	expiration = time.Duration(int64(jwtConfig.Int("expirationTimeSeconds"))) * time.Second
)

// Code is used to create and validate verification code
type Code struct{}

// Create is used to create a jwt for verification token
func (code Code) Create(userID string) string {
	claims := createClaims(userID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(secret)
	return tokenString
}

// Validate is used to validate a jwt
func (code Code) Validate(tokenString string) (bool, string) {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, jwtKeyFunc)
	if err != nil {
		return false, ""
	}

	expirationTime := time.Unix(claims.ExpirationUTC, 0)
	if isExpired(expirationTime) {
		return false, ""
	}

	return token.Valid, claims.UserID
}
