package handler

import (
	"context"

	"github.com/ZeroTechh/VelocityCore/logger"
	proto "github.com/ZeroTechh/VelocityCore/proto/VerificationCodeService"
	"github.com/ZeroTechh/blaze"
	"github.com/ZeroTechh/hades"
	"go.uber.org/zap"

	"github.com/ZeroTechh/VerificationCodeService/core/code"
)

var (
	config = hades.GetConfig("main.yaml", []string{"config", "../config"})
	log    = logger.GetLogger(
		config.Map("service").Str("logFile"),
		config.Map("service").Bool("debug"),
	)
)

// Handler is used to handle all verification code service functions
type Handler struct {
	code code.Code
}

// Create is used to create a verification code jwt
func (handler Handler) Create(
	ctx context.Context,
	request *proto.UserData) (*proto.TokenData, error) {
	funcLog := blaze.NewFuncLog(
		"VerificationCodeService.Handler.Create",
		log,
		zap.String("UserID", request.UserID),
	)
	funcLog.Started()

	token := handler.code.Create(request.UserID)

	funcLog.Completed(zap.String("Token", token))

	return &proto.TokenData{Token: token}, nil
}

// Validate is used to validate a verification code jwt
func (handler Handler) Validate(
	ctx context.Context,
	request *proto.TokenData) (*proto.Valid, error) {
	funcLog := blaze.NewFuncLog(
		"VerificationCodeService.Handler.Validate",
		log,
		zap.String("Token", request.Token),
	)
	funcLog.Started()

	valid, userID := handler.code.Validate(request.Token)

	funcLog.Completed(zap.Bool("Valid", valid))
	return &proto.Valid{Valid: valid, UserID: userID}, nil
}
