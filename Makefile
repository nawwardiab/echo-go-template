include .env

all: compile
	
compile: *.go config.yaml
	@echo "–> compiling"
	go build

run: compile
	@echo "–> running the app"
	go run main.go

# Datbase 
DB_URL := postgres://${DB_USER}:${DB_PWD}@${DB_HOST}:${DB_PORT}/${DB_NAME}
MIGRATE_TOOL := /home/dci-student/go/bin/migrate
MIGRATE_DIR := migrations

MIGRATE := ${MIGRATE_TOOL} -path ${MIGRATE_DIR} -database ${DB_URL}

migrate-up:
	@${MIGRATE} up

migrate-down:
	@${MIGRATE} down

clean:
	@echo "–> cleaning"
	@rm -f echo-server

reload: clean run