PROJECTNAME=$(shell basename "$(PWD)")

GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
PKG = $(shell go list ./...)
PORT ?= 8080
ENTRYPOINT=./cmd/app
BINARY_NAME=$(PROJECTNAME)

dev: serve

## serve: Start in development mode. Auto-starts on code changes
serve:
	@echo "Building API"
	@fresh
