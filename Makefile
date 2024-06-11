build:
	@go build -o bin/zeep-lend cmd/main.go

test:
	@go test -v ./bin/... ./cmd/... ./config/... ./db/... ./logger/... \
	./services/... ./types/... ./utils/...

run: build
	@./bin/zeep-lend

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations \
	$(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down