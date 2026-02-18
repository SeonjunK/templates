# Python Project Template

## Project Structure
```
python/
├── src/           # Source code
├── tests/         # Test files
└── pyproject.toml # Project configuration
```

## Commands (uv)
- Sync: `uv sync`
- Format: `uv run ruff format .`
- Lint: `uv run ruff check . --fix`
- Test: `uv run pytest`
- Type check: `uv run mypy src`

## Code Style
- Use Ruff for formatting and linting
- Line length: 88 characters
- Quote style: double quotes

## Required Tools

VS Code extensions:
- `ms-python.python`
- `ms-python.vscode-pylance`
- `charliermarsh.ruff`
- `ms-python.mypy-type-checker`
- `ryanluker.vscode-coverage-gutters`
- `emeraldwalk.runonsave`
