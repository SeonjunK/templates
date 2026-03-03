# Observability

Integrates OpenTelemetry tracing and stdlib logging through Logfire.

## Logging Flow

```
logging (stdlib)
    └── LogfireLoggingHandler
            └── Logfire SDK
                    ├── BatchSpanProcessor → OTLPSpanExporter  (when otel.endpoint is set)
                    └── send_to_logfire                        (when otel.send_to_logfire=true)
```

---

## Configuration Parameters

| Field | Default | Description |
|:------|:--------|:------------|
| `otel.endpoint` | `""` | OTLP HTTP endpoint (empty = console output only) |
| `otel.headers` | `{}` | Non-sensitive OTLP headers (e.g. `X-Tenant-ID`) |
| `otel.secret_headers` | `{}` | Sensitive OTLP headers — values masked in logs (e.g. `Authorization`) |
| `otel.send_to_logfire` | `false` | Whether to send traces to Logfire platform |
| `app.log_level` | `"INFO"` | Standard logging level (`DEBUG`/`INFO`/`WARNING`/`ERROR`/`CRITICAL`) |

### `OtelConfig.all_headers`

`log.py` passes `otel.all_headers` (not `otel.headers`) to the OTLP exporter. This property merges `headers` and `secret_headers` into a single `dict[str, str]`, with `secret_headers` values taking precedence on key conflict. Callers never handle `SecretStr` directly.

---

## `otel.send_to_logfire` Behavior

When `false` (default), Logfire SDK runs in local mode — spans are not sent to the Logfire cloud platform. When `true`, the SDK authenticates with the Logfire platform using the `LOGFIRE_TOKEN` environment variable and streams traces there.

Setting `otel.endpoint` and `otel.send_to_logfire=true` simultaneously is supported — traces are sent to both destinations.
