#!/bin/sh

echo "Building gitauthor cli"

go build -o ./bin/gitauthor ./cmd/main.go
