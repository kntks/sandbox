package logger

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

func ZapProductionXError() {
	config := zap.NewProductionConfig()
	config.DisableStacktrace = true
	logger, _ := config.Build()
	defer logger.Sync()

	if err := returnxError1(); err != nil {
		logger.Sugar().Errorf("%+v\n", err)
	}
}

func ZapDevelopmentXError() {
	config := zap.NewDevelopmentConfig()
	config.DisableStacktrace = true
	logger, _ := config.Build()
	defer logger.Sync()

	if err := returnxError1(); err != nil {
		logger.Sugar().Errorf("%+v\n", err)
	}
}

func returnxError() error {
	return xerrors.New("error")
}

func returnxError1() error {
	err := returnxError()
	return xerrors.Errorf("error1: %w", err)
}

func ZapDevelopmentError() {
	config := zap.NewDevelopmentConfig()
	config.DisableStacktrace = true
	logger, _ := config.Build()
	defer logger.Sync()

	if err := returnError1(); err != nil {
		logger.Sugar().Errorf("%+v\n", err)
	}
}

func returnError() error {
	return errors.New("error")
}

func returnError1() error {
	err := returnError()
	return fmt.Errorf("error1: %w", err)
}
