#!/usr/bin/env bash

mkdir -p output/server output/client

echo "build godis server"
cd server || exit
./build.sh
cp -r output ../output/server
cd ..

echo "build godis client"
cd client || exit
./build.sh
cp -r output ../output/client
cd ..