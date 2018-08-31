include .env
export
# PROJECTNAME=$(shell basename "$(PWD)")

# GOCMD = go
# GOBUILD = $(GOCMD) build
# GOTEST = $(GOCMD) test
# PKG = $(shell go list ./...)
# PORT ?= 8080
# ENTRYPOINT=./
# BINARY_NAME=$(PROJECTNAME)
DB_CONTAINER=mongo

dev: 
	@make db_build
	@make db_add_auth
	#@make serve

## serve: Start in development mode. Auto-starts on code changes
serve:
	@echo "Building API"
	@docker-compose up --build state-api

db_build:
	@echo "Setting up Dev DB"
	@docker-compose up --build -d mongo

load_fixtures:
	./fixtures --host ${MONGO_HOST} --db ${MONGO_DB} --user ${MONGO_USER} --pass ${MONGO_PW}