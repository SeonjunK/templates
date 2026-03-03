from __future__ import annotations

from contextvars import ContextVar
from pathlib import Path
from typing import Any

from pydantic import BaseModel
from pydantic.config import ConfigDict
from pydantic_settings import (
    BaseSettings,
    CliSettingsSource,
    PydanticBaseSettingsSource,
    SettingsConfigDict,
    YamlConfigSettingsSource,
)

_yaml_file_var: ContextVar[Path | None] = ContextVar("yaml_file", default=None)


class ConfigModel(BaseModel):
    """Base class for nested configuration sections."""

    model_config = ConfigDict(extra="ignore")


class BaseConfig(BaseSettings):
    """Base settings class with ContextVar-based YAML path injection.

    Source priority (highest to lowest):
        init kwargs > CLI args > environment variables > .env file > YAML file

    Usage::

        settings = Settings()                              # uses config.yaml
        settings = Settings(_yaml_file=Path("test.yaml"))  # override for testing
    """

    model_config = SettingsConfigDict(
        env_prefix="APP__",
        env_nested_delimiter="__",
        extra="ignore",
    )

    def __init__(self, _yaml_file: Path | None = None, **data: Any) -> None:
        effective_yaml = (
            _yaml_file if _yaml_file is not None else Path.cwd() / "config.yaml"
        )
        token = _yaml_file_var.set(effective_yaml)
        try:
            super().__init__(**data)
        finally:
            _yaml_file_var.reset(token)

    @classmethod
    def settings_customise_sources(
        cls,
        settings_cls: type[BaseSettings],
        init_settings: PydanticBaseSettingsSource,
        env_settings: PydanticBaseSettingsSource,
        dotenv_settings: PydanticBaseSettingsSource,
        file_secret_settings: PydanticBaseSettingsSource,
    ) -> tuple[PydanticBaseSettingsSource, ...]:
        return (
            init_settings,
            CliSettingsSource(
                settings_cls, cli_ignore_unknown_args=True, cli_parse_args=True
            ),
            env_settings,
            dotenv_settings,
            *(
                (YamlConfigSettingsSource(settings_cls, yaml_file=f),)
                if (f := _yaml_file_var.get()) is not None and f.exists()
                else ()
            ),
        )
