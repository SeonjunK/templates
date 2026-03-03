from __future__ import annotations

import sys
from pathlib import Path

import pytest
from pydantic import SecretStr, ValidationError

from config import (
    AppConfig,
    AuthConfig,
    DatabaseConfig,
    ServerConfig,
    Settings,
    get_settings,
)


class TestConfig:
    def test_missing_raises_error(self, tmp_path: Path) -> None:
        """Raises ValidationError when secret_key is missing."""
        empty = tmp_path / "empty.yaml"
        empty.write_text("""\
{}
""")
        with pytest.raises(ValidationError):
            Settings(_yaml_file=empty)

    def test_default(self, tmp_path: Path) -> None:
        """Fields not specified in YAML are filled with model defaults."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("""\
auth:
  secret_key: yaml-secret
""")
        s = Settings(_yaml_file=yaml)

        assert s.auth.secret_key.get_secret_value() == "yaml-secret"
        assert s.auth.algorithm == "HS256"
        assert s.auth.expire_minutes == 60

        assert s.app.name == "my-app"
        assert s.app.env == "development"
        assert s.app.debug is False
        assert s.app.log_level == "INFO"

        assert s.server.host == "127.0.0.1"
        assert s.server.port == 8000
        assert s.server.workers == 1
        assert s.server.reload is False

        assert s.database.url.get_secret_value() == "sqlite:///./dev.db"
        assert s.database.pool_size == 5
        assert s.database.max_overflow == 10
        assert s.database.echo is False

    def test_yaml(self, tmp_path: Path) -> None:
        """All fields in YAML are loaded correctly."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("""\
auth:
  secret_key: yaml-secret
  algorithm: RS256
  expire_minutes: 30

app:
  name: yaml-app
  env: production
  debug: true
  log_level: DEBUG

server:
  host: 0.0.0.0
  port: 9000
  workers: 4
  reload: true

database:
  url: postgresql://localhost/db
  pool_size: 10
  max_overflow: 20
  echo: true
""")
        s = Settings(_yaml_file=yaml)

        assert s.auth.secret_key.get_secret_value() == "yaml-secret"
        assert s.auth.algorithm == "RS256"
        assert s.auth.expire_minutes == 30

        assert s.app.name == "yaml-app"
        assert s.app.env == "production"
        assert s.app.debug is True
        assert s.app.log_level == "DEBUG"

        assert s.server.host == "0.0.0.0"
        assert s.server.port == 9000
        assert s.server.workers == 4
        assert s.server.reload is True

        assert s.database.url.get_secret_value() == "postgresql://localhost/db"
        assert s.database.pool_size == 10
        assert s.database.max_overflow == 20
        assert s.database.echo is True

    def test_dotenv(self, tmp_path: Path) -> None:
        """All fields in a dotenv file are loaded correctly."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("{}\n")
        dotenv = tmp_path / ".env"
        dotenv.write_text("""\
APP__AUTH__SECRET_KEY=dotenv-secret
APP__AUTH__ALGORITHM=RS256
APP__AUTH__EXPIRE_MINUTES=30
APP__APP__NAME=dotenv-app
APP__APP__ENV=production
APP__APP__DEBUG=true
APP__APP__LOG_LEVEL=DEBUG
APP__SERVER__HOST=0.0.0.0
APP__SERVER__PORT=9000
APP__SERVER__WORKERS=4
APP__SERVER__RELOAD=true
APP__DATABASE__URL=postgresql://localhost/db
APP__DATABASE__POOL_SIZE=10
APP__DATABASE__MAX_OVERFLOW=20
APP__DATABASE__ECHO=true
""")
        s = Settings(_yaml_file=yaml, _env_file=dotenv)

        assert s.auth.secret_key.get_secret_value() == "dotenv-secret"
        assert s.auth.algorithm == "RS256"
        assert s.auth.expire_minutes == 30

        assert s.app.name == "dotenv-app"
        assert s.app.env == "production"
        assert s.app.debug is True
        assert s.app.log_level == "DEBUG"

        assert s.server.host == "0.0.0.0"
        assert s.server.port == 9000
        assert s.server.workers == 4
        assert s.server.reload is True

        assert s.database.url.get_secret_value() == "postgresql://localhost/db"
        assert s.database.pool_size == 10
        assert s.database.max_overflow == 20
        assert s.database.echo is True

    def test_env(self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch) -> None:
        """All fields are loaded correctly via environment variables."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("{}\n")
        monkeypatch.setenv("APP__AUTH__SECRET_KEY", "env-secret")
        monkeypatch.setenv("APP__AUTH__ALGORITHM", "RS256")
        monkeypatch.setenv("APP__AUTH__EXPIRE_MINUTES", "30")
        monkeypatch.setenv("APP__APP__NAME", "env-app")
        monkeypatch.setenv("APP__APP__ENV", "production")
        monkeypatch.setenv("APP__APP__DEBUG", "true")
        monkeypatch.setenv("APP__APP__LOG_LEVEL", "DEBUG")
        monkeypatch.setenv("APP__SERVER__HOST", "0.0.0.0")
        monkeypatch.setenv("APP__SERVER__PORT", "9000")
        monkeypatch.setenv("APP__SERVER__WORKERS", "4")
        monkeypatch.setenv("APP__SERVER__RELOAD", "true")
        monkeypatch.setenv("APP__DATABASE__URL", "postgresql://localhost/db")
        monkeypatch.setenv("APP__DATABASE__POOL_SIZE", "10")
        monkeypatch.setenv("APP__DATABASE__MAX_OVERFLOW", "20")
        monkeypatch.setenv("APP__DATABASE__ECHO", "true")
        s = Settings(_yaml_file=yaml)

        assert s.auth.secret_key.get_secret_value() == "env-secret"
        assert s.auth.algorithm == "RS256"
        assert s.auth.expire_minutes == 30

        assert s.app.name == "env-app"
        assert s.app.env == "production"
        assert s.app.debug is True
        assert s.app.log_level == "DEBUG"

        assert s.server.host == "0.0.0.0"
        assert s.server.port == 9000
        assert s.server.workers == 4
        assert s.server.reload is True

        assert s.database.url.get_secret_value() == "postgresql://localhost/db"
        assert s.database.pool_size == 10
        assert s.database.max_overflow == 20
        assert s.database.echo is True

    def test_cli(self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch) -> None:
        """All fields are loaded correctly via CLI arguments."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("{}\n")
        monkeypatch.setattr(
            sys,
            "argv",
            [
                "prog",
                "--auth.secret_key=cli-secret",
                "--auth.algorithm=RS256",
                "--auth.expire_minutes=30",
                "--app.name=cli-app",
                "--app.env=production",
                "--app.debug=true",
                "--app.log_level=DEBUG",
                "--server.host=0.0.0.0",
                "--server.port=9000",
                "--server.workers=4",
                "--server.reload=true",
                "--database.url=postgresql://localhost/db",
                "--database.pool_size=10",
                "--database.max_overflow=20",
                "--database.echo=true",
            ],
        )
        s = Settings(_yaml_file=yaml)

        assert s.auth.secret_key.get_secret_value() == "cli-secret"
        assert s.auth.algorithm == "RS256"
        assert s.auth.expire_minutes == 30

        assert s.app.name == "cli-app"
        assert s.app.env == "production"
        assert s.app.debug is True
        assert s.app.log_level == "DEBUG"

        assert s.server.host == "0.0.0.0"
        assert s.server.port == 9000
        assert s.server.workers == 4
        assert s.server.reload is True

        assert s.database.url.get_secret_value() == "postgresql://localhost/db"
        assert s.database.pool_size == 10
        assert s.database.max_overflow == 20
        assert s.database.echo is True

    def test_init(self, tmp_path: Path) -> None:
        """All fields are loaded correctly when Config objects are passed as kwargs."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("{}\n")
        s = Settings(
            _yaml_file=yaml,
            auth=AuthConfig(
                secret_key=SecretStr("init-secret"),
                algorithm="RS256",
                expire_minutes=30,
            ),
            app=AppConfig(
                name="init-app", env="production", debug=True, log_level="DEBUG"
            ),
            server=ServerConfig(host="0.0.0.0", port=9000, workers=4, reload=True),
            database=DatabaseConfig(
                url="postgresql://localhost/db",
                pool_size=10,
                max_overflow=20,
                echo=True,
            ),
        )

        assert s.auth.secret_key.get_secret_value() == "init-secret"
        assert s.auth.algorithm == "RS256"
        assert s.auth.expire_minutes == 30

        assert s.app.name == "init-app"
        assert s.app.env == "production"
        assert s.app.debug is True
        assert s.app.log_level == "DEBUG"

        assert s.server.host == "0.0.0.0"
        assert s.server.port == 9000
        assert s.server.workers == 4
        assert s.server.reload is True

        assert s.database.url.get_secret_value() == "postgresql://localhost/db"
        assert s.database.pool_size == 10
        assert s.database.max_overflow == 20
        assert s.database.echo is True

    def test_priority_dotenv_over_yaml(self, tmp_path: Path) -> None:
        """dotenv > yaml: values in dotenv take precedence over YAML."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("""\
auth:
  secret_key: yaml-secret
""")
        dotenv = tmp_path / ".env"
        dotenv.write_text("""\
APP__AUTH__SECRET_KEY=dotenv-secret
""")
        s = Settings(_yaml_file=yaml, _env_file=dotenv)
        assert s.auth.secret_key.get_secret_value() == "dotenv-secret"

    def test_priority_env_over_dotenv(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch
    ) -> None:
        """env > dotenv: environment variables take precedence over dotenv file."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("""\
auth:
  secret_key: yaml-secret
""")
        dotenv = tmp_path / ".env"
        dotenv.write_text("""\
APP__AUTH__SECRET_KEY=dotenv-secret
""")
        monkeypatch.setenv("APP__AUTH__SECRET_KEY", "env-secret")
        s = Settings(_yaml_file=yaml, _env_file=dotenv)
        assert s.auth.secret_key.get_secret_value() == "env-secret"

    def test_priority_cli_over_env(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch
    ) -> None:
        """CLI > env: CLI arguments take precedence over environment variables."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("""\
auth:
  secret_key: yaml-secret
""")
        monkeypatch.setenv("APP__AUTH__SECRET_KEY", "env-secret")
        monkeypatch.setattr(sys, "argv", ["prog", "--auth.secret_key=cli-secret"])
        s = Settings(_yaml_file=yaml)
        assert s.auth.secret_key.get_secret_value() == "cli-secret"

    def test_priority_init_over_cli(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch
    ) -> None:
        """init > CLI: init kwargs take precedence over CLI arguments."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("""\
auth:
  secret_key: yaml-secret
""")
        monkeypatch.setattr(sys, "argv", ["prog", "--auth.secret_key=cli-secret"])
        s = Settings(
            _yaml_file=yaml,
            auth=AuthConfig(secret_key=SecretStr("init-secret")),
        )
        assert s.auth.secret_key.get_secret_value() == "init-secret"


class TestOtelConfig:
    def test_send_to_logfire_default(self, tmp_path: Path) -> None:
        """send_to_logfire defaults to False."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("auth:\n  secret_key: secret\n")
        s = Settings(_yaml_file=yaml)
        assert s.otel.send_to_logfire is False

    def test_send_to_logfire_yaml(self, tmp_path: Path) -> None:
        """send_to_logfire can be enabled via YAML."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("""\
auth:
  secret_key: secret
otel:
  send_to_logfire: true
""")
        s = Settings(_yaml_file=yaml)
        assert s.otel.send_to_logfire is True

    def test_send_to_logfire_env(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch
    ) -> None:
        """send_to_logfire can be enabled via environment variable."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("auth:\n  secret_key: secret\n")
        monkeypatch.setenv("APP__OTEL__SEND_TO_LOGFIRE", "true")
        s = Settings(_yaml_file=yaml)
        assert s.otel.send_to_logfire is True

    def test_otel_defaults(self, tmp_path: Path) -> None:
        """OtelConfig fields have correct defaults."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("auth:\n  secret_key: secret\n")
        s = Settings(_yaml_file=yaml)
        assert s.otel.endpoint == ""
        assert s.otel.headers == {}
        assert s.otel.secret_headers == {}
        assert s.otel.send_to_logfire is False

    def test_otel_secret_headers_yaml(self, tmp_path: Path) -> None:
        """secret_headers loaded from YAML are stored as SecretStr."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("""\
auth:
  secret_key: secret
otel:
  secret_headers:
    Authorization: Bearer token-from-yaml
""")
        s = Settings(_yaml_file=yaml)
        assert isinstance(s.otel.secret_headers["Authorization"], SecretStr)
        assert (
            s.otel.secret_headers["Authorization"].get_secret_value()
            == "Bearer token-from-yaml"
        )

    def test_otel_secret_headers_env(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch
    ) -> None:
        """secret_headers loaded from env var are stored as SecretStr."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("auth:\n  secret_key: secret\n")
        monkeypatch.setenv(
            "APP__OTEL__SECRET_HEADERS", '{"Authorization": "Bearer env-token"}'
        )
        s = Settings(_yaml_file=yaml)
        assert isinstance(s.otel.secret_headers["Authorization"], SecretStr)
        assert (
            s.otel.secret_headers["Authorization"].get_secret_value()
            == "Bearer env-token"
        )

    def test_otel_all_headers_merges(self, tmp_path: Path) -> None:
        """all_headers property returns merged headers and secret_headers."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("""\
auth:
  secret_key: secret
otel:
  headers:
    X-Tenant-ID: tenant-a
  secret_headers:
    Authorization: Bearer token
""")
        s = Settings(_yaml_file=yaml)
        merged = s.otel.all_headers
        assert merged == {"X-Tenant-ID": "tenant-a", "Authorization": "Bearer token"}

    def test_otel_all_headers_secret_precedence(self, tmp_path: Path) -> None:
        """secret_headers takes precedence over headers on key conflict."""
        yaml = tmp_path / "config.yaml"
        yaml.write_text("""\
auth:
  secret_key: secret
otel:
  headers:
    Authorization: plain-value
  secret_headers:
    Authorization: secret-value
""")
        s = Settings(_yaml_file=yaml)
        assert s.otel.all_headers["Authorization"] == "secret-value"


class TestGetSettings:
    def test_returns_settings_instance(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch
    ) -> None:
        """get_settings() returns a valid Settings instance."""
        get_settings.cache_clear()
        monkeypatch.chdir(tmp_path)
        (tmp_path / "config.yaml").write_text("auth:\n  secret_key: cached-secret\n")
        try:
            s = get_settings()
            assert s.auth.secret_key.get_secret_value() == "cached-secret"
        finally:
            get_settings.cache_clear()

    def test_cached(self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch) -> None:
        """get_settings() returns the same object on repeated calls."""
        get_settings.cache_clear()
        monkeypatch.chdir(tmp_path)
        (tmp_path / "config.yaml").write_text("auth:\n  secret_key: cached-secret\n")
        try:
            assert get_settings() is get_settings()
        finally:
            get_settings.cache_clear()
