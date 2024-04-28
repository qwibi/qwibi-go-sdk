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

release:
	@latest_tag=$$(git describe --tags --abbrev=0); \
	IFS='.' read -ra parts <<< "$$latest_tag"; \
	major=$${parts[0]}; \
	minor=$${parts[1]}; \
	patch=$${parts[2]}; \
	new_version=$$major.$$minor.$$((patch+1)); \
	today=$$(date +'%d/%m/%Y'); \
	git add -A; \
	git commit -m "Data: $$today"; \
	git tag -d latest || true; \
	git tag -a -f latest -m "Latest version"; \
	git tag $$new_version; \
	git push --tags --force; \
	git push --force


.DEFAULT_GOAL := get