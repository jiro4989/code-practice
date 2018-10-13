#!/bin/bash

# $# は引数の数
if [ $# -lt 2 ]; then
	echo need 2 arguments. 1>&2
	exit 1
fi

echo "args[0]:$0"
echo "args[1]:$1"
echo "args[2]:$2"
sum=$(($1+$2))
echo "sum:$sum"
