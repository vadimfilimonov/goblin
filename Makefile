# make ./bin part of PATH for this Makefile
PWD   := $(PWD)
PATH  := $(PWD)/bin:$(PATH)
SHELL := env PATH=$(PATH) /bin/bash

build:
	go build -v ./...

tools:
	cd tools && go mod tidy && go generate -tags tools

install: tools
	go mod tidy

verify:
	go mod verify

lint:
	go vet ./...
	golangci-lint run ./...
	goconst ./...

format:
	gofmt -w .

test:
	go test ./...

test-coverage:
	go test ./... -cover

build-test-coverage-report:
	go test . -coverprofile=coverage.out
	go tool cover -html=coverage.out

docs:
	godoc .

.PHONY: build tools install verify lint format test test-coverage build-test-coverage-report docs
