#!/bin/bash

#  docker build --tag protobuf `pwd`/protobuf/
docker run --rm -v `pwd`/protobuf/proto:/go/src -it protobuf /bin/sh
 