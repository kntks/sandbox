package main

import (
	"fmt"
	"sandbox/logger"
)

func main() {
	fmt.Println("zap production xerror")
	//logger.ZapProductionXError()

	fmt.Println("zap development xerror")
	//logger.ZapDevelopmentXError()

	fmt.Println("zap development error")
	//logger.ZapDevelopmentError()

	fmt.Println("zap xerrors")
	//logger.XerrorsMain()

	fmt.Println("pkg error")
	logger.PkgErrorMain()
}
