package main

import (
	"testing"

	"go.uber.org/goleak"
)

func TestTicker(t *testing.T) {
	defer goleak.VerifyNone(t)
	ticker()
}
