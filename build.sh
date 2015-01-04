#!/bin/sh

rm  ./build/container-counter.*.64
GOOS=linux GOARCH=amd64 go build -o ./build/container-counter.linux.64
