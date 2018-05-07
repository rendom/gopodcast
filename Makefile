exec = .
VERSION = `git describe --tags`
BUILD_DATE=`date +%FT%T%z`
BUILD_NAME="gopodcast"
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.BuildDate=${BUILD_DATE}"

.PHONY: all build build-release test deps build-protos

all: test clean build

deps:
	dep ensure

clean:
	rm -rf dist/*

build:
	go build -v -o dist/$(BUILD_NAME) $(exe)

build-release:
	@echo "not implemented"

test:
	go test ./...

build-protos:
	./generateProtos.sh
