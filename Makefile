include .env
export

dev: 
	@make db_build
	@make serve
	@make load_fixtures

## serve: Start in development mode. Auto-starts on code changes
serve:
	@echo "Building API"
	@docker-compose build --no-cache state-api
	@docker-compose up -d state-api
	@docker exec -it state-api dep ensure --vendor-only
	@docker exec -it state-api go install -v

db_build:
	@echo "Setting up Dev DB"
	@docker-compose up --build -d mongo

load_fixtures:
	@docker exec -it state-api ./fixtures --host ${MONGO_HOST} --db ${MONGO_DB} --user ${MONGO_USER} --pass ${MONGO_PW}

test:
	@docker exec -it state-api go test ./...