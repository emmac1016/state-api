include .env
export

dev: 
	@make db_build
	@make load_fixtures
	@make serve

## serve: Start in development mode. Auto-starts on code changes
serve:
	@echo "Building API"
	@docker-compose up --build -d state-api

db_build:
	@echo "Setting up Dev DB"
	@docker-compose up --build -d mongo

load_fixtures:
	./fixtures --host ${MONGO_HOST} --db ${MONGO_DB} --user ${MONGO_USER} --pass ${MONGO_PW}

serve_and_test:
	@make dev
	@make test

test:
	@docker exec -it state-api go test ./...