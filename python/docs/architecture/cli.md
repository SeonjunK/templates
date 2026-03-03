# CLI Entry Point

## Overview

[Typer](https://typer.tiangolo.com/)-based CLI. Registered in `pyproject.toml` under `[project.scripts]` as `serve = "main:app"`.

## Startup Flow

```
uv run serve
    └── serve()          # Typer command
            ├── get_settings()
            ├── log.setup(settings.app, settings.otel)
            └── # Application initialization (TODO)
```
