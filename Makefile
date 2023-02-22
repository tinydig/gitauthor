BINARY_NAME=gitauthor
.DEFAULT_GOAL := all

all: clean dep test build

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin cmd/main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux cmd/main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/${BINARY_NAME}-windows cmd/main.go

clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
	rm ./bin/${BINARY_NAME}-linux
	rm ./bin/${BINARY_NAME}-windows

test:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download
