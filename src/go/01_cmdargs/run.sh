#!/bin/bash

set -x

go build main.go
./main
./main 1
./main 2 3
