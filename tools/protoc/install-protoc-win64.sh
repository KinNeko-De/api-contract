#! /bin/sh
ossystem=win64
currentversion="22.0"


rm -rf ${ossystem}
mkdir ${ossystem}
downloaduri=https://github.com/protocolbuffers/protobuf/releases/download/v${currentversion}/protoc-${currentversion}-${ossystem}.zip
echo "downloading from: $downloaduri"
curl --location ${downloaduri} --remote-name
unzip "protoc-$currentversion-${ossystem}.zip" -d ${ossystem}
rm "protoc-$currentversion-${ossystem}.zip"

read -p "Press [Enter] to exit."