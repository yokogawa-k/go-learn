SHELL := /bin/bash

CURRENT_REVISION = $(shell git rev-parse --short HEAD)

SRC := $(shell find . -not -path "*/vendor/*" -name "*.go")
BIN := $(shell basename $(dir $(realpath $(firstword $(MAKEFILE_LIST)))))

all: fmt lint test

build: $(SRC)
	go get ./...
	go build -ldflags="-s -w -X main.COMMIT=$(CURRENT_REVISION)" -o $(BIN) .

fmt:
	gofmt -s -w $(SRC)

test: testdeps
	go test $(VERBOSE_FLAG) ./...

lint: testdeps
	golint -set_exit_status $(SRC)

testdeps:
	go get github.com/golang/lint/golint

.PHONY: all build fmt test lint testdeps
