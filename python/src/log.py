from __future__ import annotations

import logging

import logfire
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.trace.export import BatchSpanProcessor

from config import AppConfig, OtelConfig


def setup(app: AppConfig, otel: OtelConfig) -> None:
    """Configure Logfire and route stdlib logging through it."""
    additional_span_processors: list[BatchSpanProcessor] = []
    if otel.endpoint:
        additional_span_processors.append(
            BatchSpanProcessor(
                OTLPSpanExporter(
                    endpoint=f"{otel.endpoint.rstrip('/')}/v1/traces",
                    headers=otel.all_headers,
                )
            )
        )

    logfire.configure(
        service_name=app.name,
        environment=app.env,
        send_to_logfire=otel.send_to_logfire,
        additional_span_processors=additional_span_processors,
    )
    logging.basicConfig(
        level=app.log_level,
        handlers=[logfire.LogfireLoggingHandler()],
    )
