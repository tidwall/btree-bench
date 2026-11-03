#!/bin/bash

set -e
cd $(dirname "${BASH_SOURCE[0]}")

echo
printf "\e[0;32mrichardartoul/bbtree\e[0m (Go)\n"
go run main.go
