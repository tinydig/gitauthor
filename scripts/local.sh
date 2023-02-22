#!/bin/sh

echo "Setup gitauthor cli locally"

go build -o ./bin/gitauthor ./cmd/main.go
cp ./bin/gitauthor /usr/local/bin
