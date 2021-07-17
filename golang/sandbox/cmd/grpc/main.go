package main

import (
	"flag"
	"log"
	"sandbox/grpc/client"
	"sandbox/grpc/server"
)

func main() {
	clientFlag := flag.Bool("c", false, "execute code as client")
	serverFlag := flag.Bool("s", false, "execute code as server")

	flag.Parse()

	if *clientFlag && *serverFlag {
		log.Fatal("flag option either -c or -s")
	}
	if *clientFlag {
		//client.Run()
		client.Bidirectional()
	}
	if *serverFlag {
		server.Run()
	}
}
