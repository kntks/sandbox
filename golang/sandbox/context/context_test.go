package context

import (
	"os"
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
	code := m.Run()
	os.Exit(code)
}

func TestLeakParent(t *testing.T) {
	leakParent()
}

func TestParent1(t *testing.T) {
	parent1()
}
func TestParent2(t *testing.T) {
	parent2()
}
