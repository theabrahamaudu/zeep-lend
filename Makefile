build:
	@go build -o bin/zeep-lend cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/zeep-lend