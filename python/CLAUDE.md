# Python Project Template

## Project Structure
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

## Commands (uv)
- Sync: `uv sync`
- Run: `uv run serve`
- Format: `uv run ruff format .`
- Lint: `uv run ruff check . --fix`
- Test: `uv run pytest`
- Type check: `uv run mypy src`
- Coverage: `uv run pytest --cov=src --cov-report=term-missing`

## Code Style
- Use Ruff for formatting and linting
- Line length: 88 characters
- Quote style: double quotes
- Secret fields (passwords, tokens, keys, connection URLs) must use `SecretStr`

## Required Tools

VS Code extensions:
- `ms-python.python`
- `ms-python.vscode-pylance`
- `charliermarsh.ruff`
- `ms-python.mypy-type-checker`
- `ryanluker.vscode-coverage-gutters`
- `emeraldwalk.runonsave`

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
