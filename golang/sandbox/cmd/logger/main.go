package main

import (
	"sandbox/logger"
)

func main() {
	logger.ZapProduction()

	logger.ZapDevelopment()

	logger.XerrorsMain()
}
