GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=jenkins-tools
VERSION := $(shell git describe --always --long --dirty)

.PHONY: list build build-linux clean

list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs
build:
	$(GOBUILD) -i -v -ldflags="-X 'jenkins-tools/cmd.Version=${VERSION}'" -o $(BINARY_NAME) 
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags="-X 'jenkins-tools/cmd.Version=${VERSION}' -s -w" -o $(BINARY_NAME).linux.amd64
build-darwin:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -ldflags="-X 'jenkins-tools/cmd.Version=${VERSION}' -s -w" -o $(BINARY_NAME).darwin.amd64
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME) $(BINARY_NAME).linux.amd64 $(BINARY_NAME).darwin.amd64

release: clean build-darwin build-linux

all: build-darwin build-linux
