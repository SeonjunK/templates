# Configuration System

## Class Hierarchy

```
ConfigModel (BaseModel)      ← Nested sections (extra="ignore")
    └── AppConfig
    └── ServerConfig
    └── DatabaseConfig
    └── AuthConfig
    └── OtelConfig

BaseConfig (BaseSettings)    ← Source priority definition
    └── Settings             ← Actual usage class
```

---

## Source Priority (highest to lowest)

| Priority | Source | Example |
|:--------:|:-------|:--------|
| 1 | `init kwargs` | `Settings(auth=AuthConfig(...))` |
| 2 | CLI arguments | `--auth.secret_key=...` |
| 3 | Environment variables | `APP__AUTH__SECRET_KEY=...` |
| 4 | `.env` file | `APP__AUTH__SECRET_KEY=...` |
| 5 | `config.yaml` | `auth: secret_key: ...` |

- `env_prefix = "APP__"`, `env_nested_delimiter = "__"`
- `CliSettingsSource(cli_ignore_unknown_args=True)` — prevents argv conflicts when running pytest

---

## YAML Path Injection

`BaseConfig.__init__` passes the YAML path via `ContextVar` to the point where `settings_customise_sources` is called. This is thread-safe and used for test isolation.

```python
Settings(_yaml_file=Path("test.yaml"))  # Override in tests
Settings()                              # Production: uses cwd/config.yaml
```

---

## Security Fields

Secret values must use `SecretStr`.

- `DatabaseConfig.url` — database connection string
- `AuthConfig.secret_key` — JWT signing key (required, no default)

---

## Caching Strategy

`get_settings()` is decorated with `@lru_cache(maxsize=1)` and returns a single instance for the lifetime of the process. In tests, use `Settings(_yaml_file=...)` directly and call `get_settings.cache_clear()` to isolate between test cases.
