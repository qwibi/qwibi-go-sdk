GOPATH := $(HOME)/go
SRC_PATH := $(GOPATH)/src/qwibi.com

.PHONY: build get

build:
	go build -v ./...

get:
	go get -v ./...

test:
	go test ./...

push:
	go mod tidy && git add -A && git commit -m "WIP: Update" | true && git push

.DEFAULT_GOAL := get