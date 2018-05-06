#!/bin/sh
cd protos
for protoFile in *.proto ; do \
    protoc --go_out=plugins=grpc:../gproto $protoFile; \
done
