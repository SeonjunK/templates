from __future__ import annotations

from functools import lru_cache
from typing import Literal

from pydantic import Field, SecretStr

from .base import BaseConfig, ConfigModel


class AppConfig(ConfigModel):
    name: str = "my-app"
    env: str = "development"
    debug: bool = False
    log_level: Literal["DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL"] = "INFO"


class ServerConfig(ConfigModel):
    host: str = "127.0.0.1"
    port: int = 8000
    workers: int = 1
    reload: bool = False


class DatabaseConfig(ConfigModel):
    url: SecretStr = SecretStr("sqlite:///./dev.db")
    pool_size: int = 5
    max_overflow: int = 10
    echo: bool = False


class AuthConfig(ConfigModel):
    secret_key: SecretStr
    algorithm: str = "HS256"
    expire_minutes: int = 60


class OtelConfig(ConfigModel):
    endpoint: str = ""
    headers: dict[str, str] = Field(default_factory=dict)
    secret_headers: dict[str, SecretStr] = Field(default_factory=dict)
    send_to_logfire: bool = False

    @property
    def all_headers(self) -> dict[str, str]:
        """Return merged headers for OTLP exporter use.

        secret_headers values take precedence over headers on key conflict.
        """
        return {
            **self.headers,
            **{k: v.get_secret_value() for k, v in self.secret_headers.items()},
        }


class Settings(BaseConfig):
    app: AppConfig = Field(default_factory=AppConfig)
    server: ServerConfig = Field(default_factory=ServerConfig)
    database: DatabaseConfig = Field(default_factory=DatabaseConfig)
    auth: AuthConfig
    otel: OtelConfig = Field(default_factory=OtelConfig)


@lru_cache(maxsize=1)
def get_settings() -> Settings:
    return Settings()
