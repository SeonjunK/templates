# Templates

Language-specific project templates with opinionated tooling and conventions.

## Templates

| Language | Tools |
|----------|-------|
| [Go](./go/) | golangci-lint, goimports, golines |
| [Python](./python/) | uv, ruff, mypy, pytest |

## Shared Components

| Component | Description |
|-----------|-------------|
| [go-hook](./go-hook/) | Language-agnostic Claude Code hook binary (Go) |

## Setup

```bash
# Build hook binary for all templates
make build-hook
```

## Usage

Copy the desired template directory as a starting point for a new project.

> [!NOTE]
> After copying, uncomment `.vscode/` in `.gitignore` to exclude IDE settings from version control.
> ```
> # .vscode/  →  .vscode/
> ```
