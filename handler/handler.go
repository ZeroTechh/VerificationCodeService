package handler

import (
	"context"

	proto "github.com/ZeroTechh/VelocityCore/proto/VerificationCodeService"

	"github.com/ZeroTechh/VerificationCodeService/core/code"
)

// Handler handles all verification code service functions.
type Handler struct{}

// Create creates a verification token.
func (Handler) Create(ctx context.Context, request *proto.UserData) (*proto.TokenData, error) {
	t, err := code.Create(request.UserID)
	return &proto.TokenData{Token: t}, err
}

// Validate validates verification token.
func (Handler) Validate(ctx context.Context, request *proto.TokenData) (*proto.Valid, error) {
	valid, userID := code.Validate(request.Token)
	return &proto.Valid{Valid: valid, UserID: userID}, nil
}
