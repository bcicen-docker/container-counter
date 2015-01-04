#!/bin/sh

rm  container-counter.*.64
GOOS=linux go build -o container-counter.linux.64
