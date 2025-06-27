.PHONY: build run infra generate db_ping migrate_up

default: build

infra:
	docker compose up -d database

build:
	go build -o bin/main ./cmd/main

run: infra build
	./bin/main

generate:
	sqlc generate -f ./db/sqlc.yaml

db_ping:
	@source .env && \
	PGPASSWORD=$${DB_PASSWORD} psql -h localhost -U "$${DB_USER}" -d "$${DB_NAME}" -p "$${DB_PORT}" -c "\q" && \
	echo "Pong"

db_connect:
	@source .env && \
	PGPASSWORD=$${DB_PASSWORD} psql -h localhost -U "$${DB_USER}" -d "$${DB_NAME}" -p "$${DB_PORT}"

migrate_up:
	@source .env && \
	goose postgres "host=$${DB_HOST} port=$${DB_PORT} user=$${DB_USER} password=$${DB_PASSWORD} dbname=$${DB_NAME} sslmode=disable" up
