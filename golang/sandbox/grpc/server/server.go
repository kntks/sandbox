package server

import (
	"context"
	"log"
	"net"
	"sandbox/protobuf"

	"google.golang.org/grpc"
)

type Hoge struct {
	protobuf.UnimplementedHogeServer
}

func (h Hoge) AtoB(ctx context.Context, aa *protobuf.AA) (*protobuf.BB, error) {
	switch aa.Name {
	case "AA":
		return &protobuf.BB{
			Name: "this is BB",
		}, nil
	case "BB":
		return &protobuf.BB{
			Name: "this is AA",
		}, nil
	default:
		return &protobuf.BB{
			Name: "this is default",
		}, nil
	}
}

func Run() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	hogeSrv := Hoge{}
	grpcServer := grpc.NewServer()

	protobuf.RegisterHogeServer(grpcServer, hogeSrv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
