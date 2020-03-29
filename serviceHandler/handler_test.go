package serviceHandler

import (
	"context"
	"testing"

	proto "github.com/ZeroTechh/VelocityCore/proto/VerificationCodeService"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	assert := assert.New(t)

	handler := Handler{}

	token, err := handler.Create(context.TODO(), &proto.UserData{
		UserID: "Test",
	})
	assert.NoError(err)
	assert.NotZero(token.Token)

	valid, err := handler.Validate(context.TODO(), &proto.TokenData{
		Token: token.Token,
	})
	assert.NoError(err)
	assert.True(valid.Valid)
	assert.Equal("Test", valid.UserID)
}
