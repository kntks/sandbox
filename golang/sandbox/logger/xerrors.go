package logger

import (
	"log"
	"os"

	"golang.org/x/xerrors"
)

func XerrorsMain() {
	logger := log.New(os.Stdout, "[Error]", log.LstdFlags|log.Llongfile)

	if err := returnXError1(); err != nil {
		logger.Printf("%+v\n", err)
	}
}

func returnXError() error {
	return xerrors.New("this is error")
}

func returnXError1() error {
	err := returnXError()
	return xerrors.Errorf("this is error1: %w", err)
}
