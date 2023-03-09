#!/bin/bash
CURRENT_DIR=$(pwd)
for x in $(find ${CURRENT_DIR}/proto/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/proto -I /usr/local/include --go_out=${CURRENT_DIR} \
   --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
done; 