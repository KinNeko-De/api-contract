#!/bin/bash

protobase=proto
protobasepackage=kinnekode
protogoogle=proto/google
protopath=${protobase}/${protobasepackage}

# cd ${protobasepackage}
# outdated=( $(\ls -d */) )
# for o in "${outdated[@]}"
# do
#     find ${o} -name *.pb.go \
#     -exec rm {} +
#     echo "Removed generated js code for $o"
# done
# cd - > /dev/null

cd ..

cd ${protopath}
projects=( $(\ls -d */ ) )
cd - > /dev/null

outputpath=js

for i in "${projects[@]}"
do
    find ${protopath}/${i} -name *.proto \
    -exec tools/protoc/win64/bin/protoc.exe --proto_path=${protobase} --js_out=import_style=commonjs:${outputpath} --grpc-web_out=import_style=typescript,mode=grpcwebtext:${outputpath} {} +
    echo "Generated js code for: $i"
done

read -p "Press [Enter] to exit."
