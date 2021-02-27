#!/bin/sh

echo "Building the system「Host」 ..."

cd Host/proxy

docker-compose build

docker-compose up -d

cd ../cpu

go build

./cpu


