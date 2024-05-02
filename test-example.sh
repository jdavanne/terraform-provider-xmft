#!/usr/bin/env bash

set -euo pipefail

for folder in ./examples/resources/* ; do
    (rm -rf ./tmp
    mkdir ./tmp
    cd ./tmp
    echo "$folder"
    cp "../$folder/"* .
    cp ../examples/provider/* .
    terraform init
    terraform plan )
done
