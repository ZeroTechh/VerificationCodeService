package code

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func createExpiredToken() string {
	claims := createClaims("test")
	claims.ExpirationUTC = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(secret)
	return tokenString
}

func TestCode(t *testing.T) {
	assert := assert.New(t)
	code := Code{}

	// Testing creation of a token
	token := code.Create("Test")
	assert.NotZero(token)

	// Testing Validation of a token
	valid, id := code.Validate(token)
	assert.True(valid)
	assert.Equal("Test", id)

	// Testing if validate can detect invalid token
	invalidAccessToken := "eyJhbGciOiJQUzM4NCIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.MqF1AKsJkijKnfqEI3VA1OnzAL2S4eIpAuievMgD3tEFyFMU67gCbg-fxsc5dLrxNwdZEXs9h0kkicJZ70mp6p5vdv-j2ycDKBWg05Un4OhEl7lYcdIsCsB8QUPmstF-lQWnNqnq3wra1GynJrOXDL27qIaJnnQKlXuayFntBF0j-82jpuVdMaSXvk3OGaOM-7rCRsBcSPmocaAO-uWJEGPw_OWVaC5RRdWDroPi4YL4lTkDEC-KEvVkqCnFm_40C-T_siXquh5FVbpJjb3W2_YvcqfDRj44TsRrpVhk6ohsHMNeUad_cxnFnpolIKnaXq_COv35e9EgeQIPAbgIeg"
	valid, id = code.Validate(invalidAccessToken)
	assert.False(valid)
	assert.Zero(id)

	// Testing if validate can detect expired token
	expiredToken := createExpiredToken()
	valid, id = code.Validate(expiredToken)
	assert.False(valid)
	assert.Zero(id)
}
