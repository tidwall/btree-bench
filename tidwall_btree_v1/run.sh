#!/bin/bash

set -e
cd $(dirname "${BASH_SOURCE[0]}")

echo
printf "\e[0;32mtidwall/btree v1\e[0m (Go, btree.Set) (previous version) \n"
go run main.go
