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

tag:
	git fetch --tags
	latest_tag=$$(git describe --tags $$(git rev-list --tags --max-count=1))
	major=$$(echo $$latest_tag | cut -d. -f1)
	minor=$$(echo $$latest_tag | cut -d. -f2)
	new_minor=$$(($$minor + 1))
	new_version=$$major.$$new_minor.0
	git tag $$new_version
	#git tag -f latest
	git add -A
	git commit -m "$(DATE)"
	git push | true
	git push --tags -f

release:
	@latest_tag=$$(git describe --tags --abbrev=0); \
	IFS='.' read -ra parts <<< "$$latest_tag"; \
	major=$${parts[0]}; \
	minor=$${parts[1]}; \
	patch=$${parts[2]}; \
	new_version=$$major.$$minor.$$((patch+1)); \
	today=$$(date +'%d/%m/%Y'); \
	git add -A; \
	git commit -m "$$today"; \
	git tag $$new_version; \
	git push --tags

.DEFAULT_GOAL := get