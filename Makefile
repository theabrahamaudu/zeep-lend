build:
	@go build -o bin/zeep-lend cmd/main.go

test:
	@go test -v ./bin/... ./cmd/... ./config/... ./db/... ./logger/... \
	./services/... ./types/... ./utils/...

run: build
	@./bin/zeep-lend