FILES = cmd
BINARY = main

VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`

build:
	go build ./cmd

run:
	go run ./cmd

build-static:
	CGO_ENABLED=0 GOOS=linux go build -ldflags='-s -w -extldflags "-static"' -o bin/$(BINARY) $(FILES)

test:
	go test ./...

deps:
	go mod download
