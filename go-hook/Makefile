# Go Hook CLI Makefile.

.PHONY: fmt lint test build tidy

fmt:
	go fmt ./...
	go tool golines -w .

lint:
	go tool golangci-lint run ./...

test:
	go test -race ./...

build:
	go build -o bin/hook ./cmd/hook
	@echo "Built bin/hook"

tidy:
	go mod tidy
