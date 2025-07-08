all: compile
	
compile: *.go config.yaml
	@echo "–> compiling"
	go build

run: compile
	@echo "–> running the app"
	go run main.go

# Datbase 
MIGRATE_TOOL := /home/dci-student/go/bin/migrate
DB_URL := "postgresql://postgres:LearnGowithRelations@localhost:5432/goplay"
MIGRATE_DIR := migrations

MIGRATE := ${MIGRATE_TOOL} -path ${MIGRATE_DIR} -database ${DB_URL}

migrate-up:
	@${MIGRATE} up

migrate-down:
	@${MIGRATE} down

clean:
	@echo "–> cleaning"
	@rm -f bin

reload: clean run