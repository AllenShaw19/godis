#!/usr/bin/env bash

mkdir -p output/server output/client

echo "build godis server"
cd server
./build.sh
cp -r output output/server

echo "build godis client"
cd client
./build.sh
cp -r output output/client
