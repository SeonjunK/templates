# Python Project Template

A modern Python project template with uv, Ruff, pytest, and mypy.

## Structure

```
python/
├── src/           # Source code
├── tests/         # Test files
├── docs/
│   ├── actions/       # Action logs
│   ├── adr/           # Architecture Decision Records
│   ├── architecture/  # Architecture documentation
│   ├── guides/        # Development guides
│   └── poc/           # Proof of Concept documents
└── pyproject.toml # Project configuration
```

## Setup

```bash
uv sync
```

## Commands

- Format: `uv run ruff format .`
- Lint: `uv run ruff check . --fix`
- Test: `uv run pytest`
- Type check: `uv run mypy src`

## Documentation

- [Architecture](./docs/architecture/README.md) - Project structure and design
- [Guides](./docs/guides/README.md) - Development guidelines and conventions
- [Actions](./docs/actions/README.md) - Action logs and records
- [ADR](./docs/adr/README.md) - Architecture Decision Records
- [PoC](./docs/poc/README.md) - Proof of Concept documents

## Development

See [CLAUDE.md](./CLAUDE.md) for coding conventions and guidelines.
