package code

import "github.com/dgrijalva/jwt-go"

// Claims is used to store the claims
type Claims struct {
	UserID                     string
	CreationUTC, ExpirationUTC int64
	jwt.StandardClaims
}
