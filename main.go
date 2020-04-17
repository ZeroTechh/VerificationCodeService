package main

import (
	"go.uber.org/zap"

	"github.com/ZeroTechh/VelocityCore/logger"
	proto "github.com/ZeroTechh/VelocityCore/proto/VerificationCodeService"
	"github.com/ZeroTechh/VelocityCore/services"
	"github.com/ZeroTechh/VelocityCore/utils"
	"github.com/ZeroTechh/hades"

	"github.com/ZeroTechh/VerificationCodeService/serviceHandler"
)

func main() {
	// Loading the logger
	config := hades.GetConfig("main.yaml", []string{"config"})
	log := logger.GetLogger(
		config.Map("service").Str("logFile"),
		config.Map("service").Bool("debug"),
	)

	defer utils.HandlePanic(log)

	grpcServer, listner := utils.CreateGRPCServer(
		services.VerificationCodeService,
		log,
	)

	handler := serviceHandler.Handler{}

	proto.RegisterVerificationCodeServer(grpcServer, handler)

	log.Info("Service Started")
	if err := grpcServer.Serve(*listner); err != nil {
		log.Fatal("Service Failed With Error", zap.Error(err))
	}
}
