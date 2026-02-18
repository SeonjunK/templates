# Go Project Template

## Project Structure
```
go/
├── cmd/           # Executable binaries (main package)
├── internal/      # Private packages
├── pkg/           # Public packages
├── go.mod         # Module definition
├── go.sum         # Dependency checksums
└── .golangci.yml  # golangci-lint configuration
```

## Commands

- Format: `go fmt ./...`
- Format (long lines): `go tool golines -w .`
- Lint: `go tool golangci-lint run ./...`
- Test: `go test -race ./...`
- Test (verbose): `go test -v -race ./...`
- Build: `go build ./...`
- Add dependency: `go get <package>`
- Tidy: `go mod tidy`

## Code Style

- `gofmt` formatting required (applied automatically on save)
- `goimports` for import organization (applied automatically on save in VS Code)
- Errors must be handled or explicitly ignored (`_ =`)
- Public functions and types require godoc comments
- Test files use `_test.go` suffix

## Required Tools

`golangci-lint` and `golines` are managed via `go.mod` tool directive (no separate install needed):

```bash
# Run tools directly via go tool (uses versions pinned in go.mod)
go tool golangci-lint run ./...
go tool golines -w .
```

```bash
# goimports (for VS Code)
go install golang.org/x/tools/cmd/goimports@latest
```

VS Code extensions:
- `golang.go`