#!/usr/bin/env bash

PROJ=godis

mkdir -p output/bin
go build --race -v -o $PROJ
mv $PROJ output/bin