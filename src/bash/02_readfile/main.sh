#!/bin/bash

set -eu

i=0
while read -r line
do
	line=$(echo "$line" | sed -r -e 's@^\s+@@g' -e 's@\s+$@@g')
	if [ -z "$line" ]; then
		continue
	fi
	i=$((i+1))
	printf "%03d:%s\n" "$i" "$line"
done < "$1"

