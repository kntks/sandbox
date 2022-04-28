//go:build mage

package main

import (
	"context"
	"log"

	"github.com/magefile/mage/sh"
)

func Target(ctx context.Context, name string) {
	log.Println("hello", name)
}

// Runs go mod download and then installs the binary.
func Build() error {
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "install", "./...")
}
