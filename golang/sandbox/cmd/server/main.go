package main

import (
	_ "sandbox/server"
	"sandbox/server/simple"
)

func main() {
	//server.Start()
	simple.SimpleServer{}.Start()
}
