MIGRATE_BIN := $(shell which migrate)

build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom

migration:
ifdef MIGRATE_BIN
	@$(MIGRATE_BIN) create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))
else
	$(error "migrate command not found. Please install 'migrate' binary.")
endif

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down