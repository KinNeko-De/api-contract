#!/bin/bash
export PATH="$PATH:$(go env GOPATH)/bin"

outdated=( $(\ls -d */) )
for i in "${outdated[@]}"
do
    rm -r ${i}
    echo "Removed generated go code for $i"
done

cd ..

protobase=proto
protobasepackage=kinnekode
protopath=${protobase}/${protobasepackage}

cd ${protopath}
projects=( $(\ls -d */ ) )
cd -

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
	cd -
	echo "Generated mod files for: $i"
done

read -p "Press [Enter] to exit."
