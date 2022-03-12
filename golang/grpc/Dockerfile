FROM golang:1.18rc1-bullseye

WORKDIR /go/src

RUN apt update && \
    apt install -y protobuf-compiler=3.12.4-1

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
