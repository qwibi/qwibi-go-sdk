GOPATH := $(HOME)/go
SRC_PATH := $(GOPATH)/src/qwibi.com

.PHONY: build get

build:
	go build -v ./cmd/...

get:
	go get -v ./cmd/...

test:
	go test ./...

push:
	go mod tidy && git add -A && git commit -m "WIP: Test" | true && git push

.DEFAULT_GOAL := get