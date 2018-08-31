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

db_add_auth:
	@echo "Adding users to Dev DB"
	@sleep 3
	@docker exec -it mongo mongo < build/mongo.js