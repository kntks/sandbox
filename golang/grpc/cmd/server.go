package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "grpc/gen/go/proto/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedHogeServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) AtoB(ctx context.Context, in *pb.AA) (*pb.BB, error) {
	return &pb.BB{Name: "hello world"}, nil
}

const grpcPort = "7100"
const gwPort = "7090"

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	pb.RegisterHogeServer(s, &server{})
	log.Printf("Serving gRPC on 0.0.0.0:%s", grpcPort)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%s", grpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterHogeHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", gwPort),
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-Gateway on http://0.0.0.0:%s", gwPort)
	log.Fatalln(gwServer.ListenAndServe())
}
