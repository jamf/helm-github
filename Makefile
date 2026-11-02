SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: test build

.PHONY: build
build:
	go build -o bin/helmgithub

.PHONY: test
test:
	# main's init() requires a token; the unit tests never call the API.
	GITHUB_TOKEN=unused go test ./... -v
