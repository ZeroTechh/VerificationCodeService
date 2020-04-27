package code

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func expiredToken() (string, error) {
	c := claims("test")
	c.ExpirationUTC = time.Now().Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return t.SignedString(secret)
}

func TestCode(t *testing.T) {
	assert := assert.New(t)

	// Testing creation of a token
	token, err := Create("Test")
	assert.NoError(err)
	assert.NotZero(token)

	// Testing Validation of a token
	valid, id := Validate(token)
	assert.True(valid)
	assert.Equal("Test", id)

	// Testing if validate can detect invalid token
	invalid := "eyJhbGciOiJQUzM4NCIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.MqF1AKsJkijKnfqEI3VA1OnzAL2S4eIpAuievMgD3tEFyFMU67gCbg-fxsc5dLrxNwdZEXs9h0kkicJZ70mp6p5vdv-j2ycDKBWg05Un4OhEl7lYcdIsCsB8QUPmstF-lQWnNqnq3wra1GynJrOXDL27qIaJnnQKlXuayFntBF0j-82jpuVdMaSXvk3OGaOM-7rCRsBcSPmocaAO-uWJEGPw_OWVaC5RRdWDroPi4YL4lTkDEC-KEvVkqCnFm_40C-T_siXquh5FVbpJjb3W2_YvcqfDRj44TsRrpVhk6ohsHMNeUad_cxnFnpolIKnaXq_COv35e9EgeQIPAbgIeg"
	valid, id = Validate(invalid)
	assert.False(valid)
	assert.Zero(id)

	// Testing if validate can detect expired token
	token, err = expiredToken()
	assert.NoError(err)
	valid, id = Validate(token)
	assert.False(valid)
	assert.Zero(id)
}
