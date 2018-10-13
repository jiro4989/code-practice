#!/bin/bash

set -eu

now=$(date +'%Y%m%d_%H%M%S')
echo "$1" > "$now.txt"

