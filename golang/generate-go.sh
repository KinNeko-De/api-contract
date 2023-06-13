#!/bin/bash

cd ..

export PATH="$PATH:$(go env GOPATH)/bin"

protobase=proto
protopath=${protobase}/kinnekode
projects=( $(ls ${protopath}) )
outputpath=golang

for i in "${projects[@]}"
do
    find ${protobase}/${i} -name *.pb.go \
    -exec rm {} +
    echo "Removed generated go code"

    find ${protopath}/${i} -name *.proto \
    -exec tools/protoc/win64/bin/protoc.exe --proto_path=${protobase} --go_out=${outputpath} --go_opt=paths=source_relative --go-grpc_out=${outputpath} --go-grpc_opt=paths=source_relative {} +

    echo "Generated go code for: $i"
done

read -p "Press [Enter] to exit."