FROM golang:1.16.5-stretch

RUN apt update && apt install unzip

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

# https://grpc.io/docs/protoc-installation/#install-pre-compiled-binaries-any-os
ENV PATH=$PATH:/root/.local/bin
ENV PROTO_VARSION=3.17.3
ENV PROTOC_ZIP=protoc-${PROTO_VARSION}-linux-x86_64.zip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTO_VARSION}/$PROTOC_ZIP 
RUN unzip protoc-${PROTO_VARSION}-linux-x86_64.zip  -d /root/.local

WORKDIR /go/src