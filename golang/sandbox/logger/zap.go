package logger

import (
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

func ZapProduction() {
	config := zap.NewProductionConfig()
	config.DisableStacktrace = true
	logger, _ := config.Build()
	defer logger.Sync()

	if err := returnError1(); err != nil {
		logger.Sugar().Errorf("%+v\n", err)
	}
}

func ZapDevelopment() {
	config := zap.NewDevelopmentConfig()
	config.DisableStacktrace = true
	logger, _ := config.Build()
	defer logger.Sync()

	if err := returnError1(); err != nil {
		logger.Sugar().Errorf("%+v\n", err)
	}
}

func returnError() error {
	return xerrors.New("error")
}

func returnError1() error {
	err := returnError()
	return xerrors.Errorf("error1: %w", err)
}
