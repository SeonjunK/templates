# Python Project Template

A modern Python project template with uv, Ruff, pytest, and mypy.

## Structure

```
python/
├── src/
│   ├── log.py             # Logfire setup
│   ├── main.py            # CLI entry point (Typer)
│   └── config/            # Configuration module
│       ├── __init__.py
│       ├── base.py        # BaseConfig, ConfigModel
│       └── settings.py    # Settings, get_settings
├── tests/
│   └── unit/
│       └── config/
│           └── test_config.py
├── docs/
│   ├── architecture/      # Architecture documentation
│   ├── guides/            # Development guides
│   ├── adr/               # Architecture Decision Records
│   ├── actions/           # Action logs
│   └── poc/               # PoC documents and datasets
├── config.yaml            # Default configuration file
├── .env.example           # Environment variable examples
└── pyproject.toml
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
