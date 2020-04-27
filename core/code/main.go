package code

import (
	"time"

	"github.com/ZeroTechh/hades"
	"github.com/dgrijalva/jwt-go"
)

var (
	config    = hades.GetConfig("main.yaml", []string{"config", "../config", "../../config"})
	jwtConfig = config.Map("JWT")
	secret    = []byte(jwtConfig.Str("secret"))
)

// Create creates jwt as a verification code.
func Create(userID string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims(userID))
	return t.SignedString(secret)
}

// Validate validates a verification code.
func Validate(token string) (bool, string) {
	var claims Claims

	t, err := jwt.ParseWithClaims(token, &claims, jwtKeyFunc)
	if err != nil {
		return false, ""
	}

	if time.Now().After(time.Unix(claims.ExpirationUTC, 0)) {
		return false, ""
	}

	return t.Valid, claims.UserID
}
