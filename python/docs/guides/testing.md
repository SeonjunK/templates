# Testing Guide

## Running Tests

```bash
uv run pytest                                        # all tests
uv run pytest --tb=short -q                          # concise output
uv run pytest --cov=src --cov-report=term-missing    # with coverage
uv run pytest tests/unit/config/                     # specific directory
```

## File Structure

```
tests/
└── unit/
    └── <module>/
        └── test_<module>.py
```

Mirror the `src/` layout under `tests/unit/`. One test file per source module.

## Test Organisation

Group related tests into a class per logical unit. Name the class after the subject being tested.

```python
class TestConfig:          # tests for Settings loading
    def test_default(...)
    def test_yaml(...)

class TestGetSettings:     # tests for get_settings() caching behaviour
    def test_cached(...)
```

Each test method must have a single-sentence docstring describing the behaviour under test, not the implementation.

```python
def test_cached(self, ...) -> None:
    """get_settings() returns the same object on repeated calls."""
```

## Fixtures

Always use pytest built-in fixtures over manual setup:

| Fixture | Use case |
|:--------|:---------|
| `tmp_path: Path` | Temporary directory, unique per test |
| `monkeypatch: pytest.MonkeyPatch` | Environment variables, `sys.argv`, `os.getcwd()` |

## Configuration Isolation

Never rely on `config.yaml` in the project root during tests. Always create a fixture YAML in `tmp_path` and pass it via `_yaml_file`.

```python
def test_yaml(self, tmp_path: Path) -> None:
    yaml = tmp_path / "config.yaml"
    yaml.write_text("auth:\n  secret_key: test-secret\n")
    s = Settings(_yaml_file=yaml)
    assert s.auth.secret_key.get_secret_value() == "test-secret"
```

### Minimal required YAML

`AuthConfig.secret_key` has no default and must always be present:

```python
yaml.write_text("auth:\n  secret_key: secret\n")
```

### dotenv isolation

Pass `_env_file` alongside `_yaml_file` to load a fixture `.env` file:

```python
dotenv = tmp_path / ".env"
dotenv.write_text("APP__AUTH__SECRET_KEY=dotenv-secret\n")
s = Settings(_yaml_file=yaml, _env_file=dotenv)
```

## Environment Variables

Use `monkeypatch.setenv` — changes are automatically reverted after the test.

```python
def test_env(self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch) -> None:
    monkeypatch.setenv("APP__AUTH__SECRET_KEY", "env-secret")
    s = Settings(_yaml_file=yaml)
    assert s.auth.secret_key.get_secret_value() == "env-secret"
```

## CLI Arguments

Override `sys.argv` via `monkeypatch.setattr`. Include a dummy program name as the first element.

```python
monkeypatch.setattr(sys, "argv", ["prog", "--auth.secret_key=cli-secret"])
s = Settings(_yaml_file=yaml)
assert s.auth.secret_key.get_secret_value() == "cli-secret"
```

## SecretStr Assertions

Call `.get_secret_value()` to access the raw secret in assertions. Never compare `SecretStr` objects directly.

```python
# correct
assert s.auth.secret_key.get_secret_value() == "expected"

# wrong — always fails, SecretStr masks the value
assert s.auth.secret_key == "expected"
```

## get_settings() Cache

`get_settings()` uses `@lru_cache`. Always clear the cache before and after tests that call it, using a `try/finally` block to guarantee cleanup even on failure.

```python
def test_cached(self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch) -> None:
    get_settings.cache_clear()
    monkeypatch.chdir(tmp_path)
    (tmp_path / "config.yaml").write_text("auth:\n  secret_key: secret\n")
    try:
        assert get_settings() is get_settings()
    finally:
        get_settings.cache_clear()
```

Use `monkeypatch.chdir(tmp_path)` so `Settings()` resolves `config.yaml` from the temp directory.

## Expected Failures

Use `pytest.raises` as a context manager for validation errors:

```python
def test_missing_secret_key_raises(self, tmp_path: Path) -> None:
    """Raises ValidationError when secret_key is missing."""
    empty = tmp_path / "empty.yaml"
    empty.write_text("{}\n")
    with pytest.raises(ValidationError):
        Settings(_yaml_file=empty)
```
