#!/bin/bash

export PATH="$PATH:$(go env GOPATH)/bin"

protobase=proto
protobasepackage=kinnekode
protopath=${protobase}/${protobasepackage}

rm -r ${protobasepackage}

cd ..

cd ${protopath}
projects=( $(command \ls -d */) )

cd - > /dev/null

outputpath=doc/${protobasepackage}
for i in "${projects[@]}"
do
    projectoutputpath=${outputpath}/${i}
    if [[ "$i" != "protobuf/" ]]
    then
    protobuf=${protopath}/protobuf
    else
    protobuf=""
    fi
    additionalprotos="$protobuf $protobase/google"
    
    mkdir -p ${projectoutputpath}

    find ${protopath}/${i} ${additionalprotos} -name *.proto \
    -exec tools/protoc/win64/bin/protoc.exe --proto_path=${protobase} --doc_out=${projectoutputpath} --doc_opt=markdown,index.md {} +

    echo "Generated doc for: $i"
done

read -p "Press [Enter] to exit."