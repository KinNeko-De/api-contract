#!/bin/bash

outdated=( $(ls -d */) )
for i in "${outdated[@]}"
do
    rm -r ${i}
    echo "Removed generated go code for $i"
done

cd ..

export PATH="$PATH:$(go env GOPATH)/bin"

protobase=proto
protopath=${protobase}/kinnekode
projects=( $(ls -d ${protopath})/*/ )
outputpath=golang

for i in "${projects[@]}"
do
    find ${i} -name *.proto \
    -exec tools/protoc/win64/bin/protoc.exe --proto_path=${protobase} --go_out=${outputpath} --go_opt=paths=source_relative --go-grpc_out=${outputpath} --go-grpc_opt=paths=source_relative {} +

    echo "Generated go code for: $i"
done

read -p "Press [Enter] to exit."