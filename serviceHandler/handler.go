package serviceHandler

import (
	"context"

	proto "github.com/ZeroTechh/VelocityCore/proto/VerificationCodeService"

	"github.com/ZeroTechh/VerificationCodeService/core/code"
)

// Handler is used to handle all verification code service functions
type Handler struct {
	code code.Code
}

// Create is used to create a verification code jwt
func (handler Handler) Create(
	ctx context.Context,
	request *proto.UserData) (*proto.TokenData, error) {
	token := handler.code.Create(request.UserID)
	return &proto.TokenData{Token: token}, nil
}

// Validate is used to validate a verification code jwt
func (handler Handler) Validate(
	ctx context.Context,
	request *proto.TokenData) (*proto.Valid, error) {
	valid, userID := handler.code.Validate(request.Token)
	return &proto.Valid{Valid: valid, UserID: userID}, nil
}
