#!/bin/bash

set -x

clojure main.clj
clojure main.clj number.csv
