FROM golang:1.16.5-alpine

ENV PATH="$PATH:$(go env GOPATH)/bin"
RUN apk update && apk add --no-cache bash \
  git \
  make \
  protoc~=3.13


RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest