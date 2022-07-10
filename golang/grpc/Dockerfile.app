FROM golang:1.18.3-alpine3.16
WORKDIR /go/src
# RUN apk add git
RUN go install github.com/cosmtrek/air@latest
