# Go project Makefile.

.PHONY: fmt lint test build tidy

fmt:
	go fmt ./...
	go tool golines -w .

lint:
	go tool golangci-lint run ./...

test:
	go test -race ./...

build:
	go build ./...

tidy:
	go mod tidy
