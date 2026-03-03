# Architecture

## Source Structure

```
src/
├── main.py          # CLI entry point (Typer)
├── log.py           # Observability setup (Logfire + OTLP)
└── config/
    ├── __init__.py  # Public API re-exports
    ├── base.py      # BaseConfig, ConfigModel (source priority definition)
    └── settings.py  # Domain Config classes, Settings, get_settings
```

## Documents

| Document | Description |
|:---------|:------------|
| [cli.md](cli.md) | CLI entry point — Typer setup, startup flow |
| [config.md](config.md) | Configuration system — class hierarchy, source priority, YAML injection, caching |
| [observability.md](observability.md) | Observability — Logfire setup, OTLP flow, log parameters |

---

## Build

Uses hatchling with `src/` mapped as the package root.

```toml
[tool.hatch.build.targets.wheel]
only-include = ["src"]

[tool.hatch.build.targets.wheel.sources]
"src" = ""
```

Installed package layout: `config/`, `main.py`, `log.py`
