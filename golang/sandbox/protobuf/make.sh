#!/bin/bash

#  docker build --tag protobuf `pwd`/protobuf/
docker run --rm -v `pwd`/protobuf:/go/src -it protobuf \
  protoc -I=proto --go_out=./ --go_opt=paths=source_relative  \
   --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/sample.proto
 