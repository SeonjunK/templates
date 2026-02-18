# Python Project Template

A modern Python project template with uv, Ruff, pytest, and mypy.

## Setup

```bash
uv sync
```

## Commands

- Format: `uv run ruff format .`
- Lint: `uv run ruff check . --fix`
- Test: `uv run pytest`
- Type check: `uv run mypy src`
