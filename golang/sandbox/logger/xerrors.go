package logger

import (
	errs "errors"
	"log"
	"os"

	"github.com/pkg/errors"

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

func PkgErrorMain() {
	log.Fatalf("%+v\n", returnPkgError1())
}

func returnPkgError1() error {
	return returnPkgError()
}

func returnPkgError() error {
	return errors.WithStack(errs.New("error"))
}
