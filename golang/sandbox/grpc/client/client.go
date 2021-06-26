package client

import (
	"context"
	"log"
	"sandbox/protobuf"

	"google.golang.org/grpc"
)

func Run() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := protobuf.NewHogeClient(conn)

	response, err := c.AtoB(context.Background(), &protobuf.AA{Name: "BB"})
	if err != nil {
		log.Fatalf("Error when calling AtoB: %s", err)
	}
	log.Printf("Response from server: %+v", response)

}
