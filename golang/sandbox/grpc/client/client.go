package client

import (
	"context"
	"io"
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

func Bidirectional() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := protobuf.NewHogeClient(conn)
	done := make(chan struct{})
	stream, err := c.AtoBstream(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s ", in.Name)
		}
	}()

	if err := stream.Send(&protobuf.AA{Name: "hogehoge"}); err != nil {
		log.Fatal(err)
	}
	stream.CloseSend()
	<-done
}
