#!/bin/bash
export GOPROXY=https://goproxy.io
export GO111MODULE=on

# please make sure running this script using ./script/setup.sh
# go run main.go
go build -o bin/fly main.go
