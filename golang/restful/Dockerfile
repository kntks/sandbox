FROM golang:alpine

# go testする時に gcc, musl-devが必要
# https://github.com/docker-library/golang/issues/86#issuecomment-191333335

RUN apk add --no-cache git gcc make musl-dev
RUN go get -u \
  github.com/cespare/reflex \
  github.com/ramya-rao-a/go-outline \
  github.com/rogpeppe/godef \
  golang.org/x/tools/gopls