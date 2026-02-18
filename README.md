# Templates

Language-specific project templates with opinionated tooling and conventions.

## Templates

| Language | Tools |
|----------|-------|
| [Go](./go/) | golangci-lint, goimports, golines |
| [Python](./python/) | uv, ruff, mypy, pytest |
| [Rust](./rust/) | cargo, clippy, rustfmt |

## Usage

Copy the desired template directory as a starting point for a new project.

> [!NOTE]
> After copying, uncomment `.vscode/` in `.gitignore` to exclude IDE settings from version control.
> ```
> # .vscode/  →  .vscode/
> ```
