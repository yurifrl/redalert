#!/bin/bash

go get -u github.com/tools/godep
go get -u google.golang.org/grpc
godep save
