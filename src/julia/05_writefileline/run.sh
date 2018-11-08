#!/bin/bash

set -x

go build main.go
./main
./main number.csv
./main number.csv out.txt
