# Go Hook CLI

Language-agnostic Claude Code hook binary written in Go.

## Project Structure
```
go-hook/
├── cmd/hook/       # CLI entry point (thin router)
├── internal/
│   ├── hook/       # Core types, input parsing, config loading
│   ├── response/   # Hook response types and builders
│   ├── guard/      # PreToolUse guard handlers (read, write, bash)
│   ├── format/     # PostToolUse format handler
│   └── verify/     # Stop verify handler
├── docs/
│   ├── architecture/  # Architecture documentation
│   └── guides/        # Development guides
├── go.mod          # Module definition
├── go.sum          # Dependency checksums
└── .golangci.yml   # golangci-lint configuration
```

## Commands

- Format: `go fmt ./...`
- Format (long lines): `go tool golines -w .`
- Lint: `go tool golangci-lint run ./...`
- Test: `go test -race ./...`
- Test (verbose): `go test -v -race ./...`
- Build: `go build -o bin/hook ./cmd/hook`
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
go tool golangci-lint run ./...
go tool golines -w .
```

```bash
# goimports (for VS Code)
go install golang.org/x/tools/cmd/goimports@latest
```

VS Code extensions:
- `golang.go`

## Hook Architecture

This binary serves as the shared hook implementation for all project templates.
Each template includes shell wrappers in `.claude/hooks/` that delegate to `bin/hook`.

### Commands
- `hook guard-read` — PreToolUse: blocks reading sensitive files
- `hook guard-write` — PreToolUse: blocks writing sensitive files
- `hook guard-bash` — PreToolUse: blocks dangerous bash commands
- `hook format` — PostToolUse: runs configured formatters on changed files
- `hook verify` — Stop: runs sequential verification steps

### Configuration
- `.claude/guard.json` — blocked file patterns and bash commands
- `.claude/hooks.json` — format rules (by extension) and verify steps

## Documentation

### docs/architecture/
Maintains documents that reflect the structure and design of source code and project files.

- Any change to code or files must be immediately reflected in the corresponding architecture document.
- Documents always represent the current state of the actual source code and files.
- Planned but not yet implemented content is marked with the `PLANNED` tag.

### docs/guides/
Maintains documents containing rules and guidelines that must be followed during development.

- Includes coding conventions, collaboration rules, and workflow processes.
- Always read the relevant guides before starting work and follow them.
