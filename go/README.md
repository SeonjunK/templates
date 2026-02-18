# Go Project Template

A modern Go project template with golangci-lint, golines, and go test.

## Setup

```bash
go mod download
```

## Commands

- Format: `go fmt ./...`
- Format (long lines): `go tool golines -w .`
- Lint: `go tool golangci-lint run ./...`
- Test: `go test -race ./...`
- Build: `go build ./...`
