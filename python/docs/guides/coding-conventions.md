# Coding Conventions

## Formatting

- **Formatter**: Ruff (`uv run ruff format .`)
- **Line length**: 88 characters
- **Quote style**: double quotes
- **Indent style**: spaces

## Linting

- **Linter**: Ruff (`uv run ruff check . --fix`)
- **Enabled rule sets**: `E`, `F`, `I`, `N`, `W`, `UP`, `B`, `S`
- **Test file exceptions**: `S101` (assert), `S104` (bind-all-interfaces) are allowed in `tests/**`

## Type Checking

- **Tool**: mypy in strict mode (`uv run mypy src`)
- All public functions and methods must have type annotations
- Use `from __future__ import annotations` for forward references

## Security Fields

### Single Sensitive Fields

Secret values (passwords, tokens, signing keys, connection URLs) **must** use `SecretStr` from pydantic.

```python
from pydantic import SecretStr

class AuthConfig(ConfigModel):
    secret_key: SecretStr  # never str
```

Affected fields in the current codebase:
- `DatabaseConfig.url` — database connection string
- `AuthConfig.secret_key` — JWT signing key

### Sensitive Dict Fields

When a config section contains a dict where some values may be sensitive (e.g., auth headers, signed params), split into two fields and expose a merge method:

- `<name>: dict[str, str]` — non-sensitive entries
- `secret_<name>: dict[str, SecretStr]` — sensitive entries
- `@property all_<name>(self) -> dict[str, str]` — merged result for actual use; `secret_<name>` takes precedence on key conflict

```python
class OtelConfig(ConfigModel):
    headers: dict[str, str] = Field(default_factory=dict)
    secret_headers: dict[str, SecretStr] = Field(default_factory=dict)

    @property
    def all_headers(self) -> dict[str, str]:
        return {
            **self.headers,
            **{k: v.get_secret_value() for k, v in self.secret_headers.items()},
        }
```

Example: `OtelConfig` uses `headers` + `secret_headers` → `all_headers` property for the OTLP exporter.

## Configuration

- Never hard-code secret values in source files or `config.yaml`
- Use environment variables for all secrets: `APP__<SECTION>__<FIELD>=...`
- See `.env.example` for all available environment variables

## Testing

- Test files: `tests/**` directory, named `test_*.py`
- Test functions: named `test_*`
- Use `Settings(_yaml_file=Path("fixture.yaml"))` for config isolation in tests
- Call `get_settings.cache_clear()` when testing with `get_settings()`

## Python Version

- Minimum supported: Python 3.10
- Use `match` statements, `X | Y` union types, and other 3.10+ features freely
