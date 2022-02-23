FROM golang:1.18beta1-bullseye

# go testする時に gcc, musl-devが必要
# https://github.com/docker-library/golang/issues/86#issuecomment-191333335

# RUN apt install musl-dev
RUN go get -v\
  github.com/cespare/reflex@latest \
  github.com/ramya-rao-a/go-outline@latest \
  github.com/rogpeppe/godef@latest \
  golang.org/x/tools/gopls@latest