# Config Package

## Overview

`internal/config` manages application configuration using [Viper](https://github.com/spf13/viper).

## Directory Structure

```
internal/config/
├── config.go       # Config struct definitions and Load function
└── config_test.go  # Unit tests

configs/
└── config.yaml     # Example config file with default values
```

## Configuration Sections

| Section    | Type             | Description                                      |
|:-----------|:-----------------|:-------------------------------------------------|
| `server`   | `ServerConfig`   | HTTP server address and timeout settings         |
| `database` | `DatabaseConfig` | DB connection DSN and connection pool settings   |
| `logging`  | `LoggingConfig`  | Log level and output format settings             |

## Priority Order

Settings are merged in the following order (higher takes precedence):

1. Environment variables (`SERVER_ADDR`, `DATABASE_DSN`, ...)
2. Config file (`./configs/config.yaml` or `./config.yaml`)
3. Code defaults

## Environment Variable Mapping

Viper automatically maps environment variables by replacing `.` with `_`.

| Config Key                    | Environment Variable             |
|:------------------------------|:---------------------------------|
| `server.addr`                 | `SERVER_ADDR`                    |
| `server.read_timeout`         | `SERVER_READ_TIMEOUT`            |
| `database.dsn`                | `DATABASE_DSN`                   |
| `database.max_open_conns`     | `DATABASE_MAX_OPEN_CONNS`        |
| `logging.level`               | `LOGGING_LEVEL`                  |
| `logging.format`              | `LOGGING_FORMAT`                 |

## Validation

`Load()` validates struct tags using `github.com/go-playground/validator/v10` internally.

| Tag | Meaning | Applied Fields |
|:----|:--------|:---------------|
| `validate:"gt=0"` | Must be greater than 0 (time.Duration compared in nanoseconds) | All timeouts and connection pool values |
| `validate:"ltefield=MaxOpenConns"` | Must be less than or equal to another field | `database.max_idle_conns` |

Error messages use `RegisterTagNameFunc` to output field names based on `mapstructure` tags and environment variable names instead of Go field names:

```
configuration validation failed:
'server.read_timeout' (env: SERVER_READ_TIMEOUT): must be greater than 0
'database.max_idle_conns' (env: DATABASE_MAX_IDLE_CONNS): must be less than or equal to MaxOpenConns
```

## Usage Example

```go
cfg, err := config.Load()
if err != nil {
    log.Fatal("failed to load config", err)
}
```

## Default Values

| Key                           | Default  |
|:------------------------------|:---------|
| `server.addr`                 | `:8080`  |
| `server.read_timeout`         | `30s`    |
| `server.write_timeout`        | `60s`    |
| `server.idle_timeout`         | `5m`     |
| `server.shutdown_timeout`     | `30s`    |
| `database.max_open_conns`     | `10`     |
| `database.max_idle_conns`     | `5`      |
| `database.conn_max_lifetime`  | `1h`     |
| `logging.level`               | `info`   |
| `logging.format`              | `json`   |
