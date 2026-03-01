# Go Project Template

A modern Go project template with golangci-lint, golines, and go test.

## Structure

```
go/
├── cmd/           # Executable binaries (main package)
├── internal/      # Private packages
├── pkg/           # Public packages
├── docs/
│   ├── actions/       # Action logs
│   ├── adr/           # Architecture Decision Records
│   ├── architecture/  # Architecture documentation
│   ├── guides/        # Development guides
│   └── poc/           # Proof of Concept documents
├── go.mod         # Module definition
├── go.sum         # Dependency checksums
└── .golangci.yml  # golangci-lint configuration
```

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

## Documentation

- [Architecture](./docs/architecture/README.md) - Project structure and design
- [Guides](./docs/guides/README.md) - Development guidelines and conventions
- [Actions](./docs/actions/README.md) - Action logs and records
- [ADR](./docs/adr/README.md) - Architecture Decision Records
- [PoC](./docs/poc/README.md) - Proof of Concept documents

## Development

See [CLAUDE.md](./CLAUDE.md) for coding conventions and guidelines.
