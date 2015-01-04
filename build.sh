#!/bin/sh

rm  ./stage/container-counter.*.64
GOOS=linux go build -o ./stage/container-counter.linux.64
