#!/bin/bash
export PATH="$PATH:$(go env GOPATH)/bin"

protobase=proto
protobasepackage=kinnekode
protopath=${protobase}/${protobasepackage}

cd ${protobasepackage}
outdated=( $(\ls -d */) )
for o in "${outdated[@]}"
do
    find ${o} -name *.pb.go \
    -exec rm {} +
    echo "Removed generated go code for $o"
    rm ${o}/go.mod
    rm ${o}/go.sum
    echo "Removed generated go mod files for $o"
done
cd - > /dev/null

cd ..

cd ${protopath}
projects=( $(\ls -d */ ) )
cd - > /dev/null

outputpath=golang

for i in "${projects[@]}"
do
    find ${protopath}/${i} -name *.proto \
    -exec tools/protoc/win64/bin/protoc.exe --proto_path=${protobase} --go_out=${outputpath} --go_opt=paths=source_relative --go-grpc_out=${outputpath} --go-grpc_opt=paths=source_relative {} +
    echo "Generated go code for: $i"
	
	cd ${outputpath}/${protobasepackage}/${i}
	modulename=$(echo github.com/kinneko-de/api-contract/golang/${protobasepackage}/${i} | sed 's/.$//')
	go mod init ${modulename}
	go mod tidy
	cd - > /dev/null
	echo "Generated mod files for: $i"
done

read -p "Press [Enter] to exit."
