SHELL := /bin/bash

all: test

test-integration-docker:
	go test -tags=integration -docker

test-integration-not-docker:
	go test -tags=integration

test-normal:
	go test

test: test-integration-docker test-integration-not-docker test-normal
