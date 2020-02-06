#!/usr/bin/env bash
export GOPATH=$(pwd)

rm -rf bin
GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o bin/nps

version='0'
small_version=$(($(cat version)+1)); printf ${small_version} > version

mkdir -p build/usr/bin
mkdir build/DEBIAN

cp bin/nps build/usr/bin

echo """package: nps
version: ${version}.${small_version}
architecture: amd64
maintainer: 17112yan@gmail.com
description: Network process

""" > build/DEBIAN/control

if [ ! -d 'dist' ]; then mkdir dist; fi
sudo dpkg -b build dist/nps_${version}.${small_version}_linux_amd64.deb

rm -rf build
