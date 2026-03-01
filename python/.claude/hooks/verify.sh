#!/bin/bash
# Run format, lint, type check, and test on Stop

if ! uv run ruff format . >/dev/null 2>&1; then
  jq -n -c '{"decision": "block", "reason": "Format failed", "systemMessage": "⚠ Format failed. Run `uv run ruff format .` to see details."}'
  exit 0
fi

if ! uv run ruff check . --fix >/dev/null 2>&1; then
  jq -n -c '{"decision": "block", "reason": "Lint failed", "systemMessage": "⚠ Lint failed. Run `uv run ruff check . --fix` to see details."}'
  exit 0
fi

if ! uv run mypy src >/dev/null 2>&1; then
  jq -n -c '{"decision": "block", "reason": "Type check failed", "systemMessage": "⚠ Type check failed. Run `uv run mypy src` to see details."}'
  exit 0
fi

if ! uv run pytest --cov=src --cov-report=xml >/dev/null 2>&1; then
  jq -n -c '{"decision": "block", "reason": "Tests failed", "systemMessage": "⚠ Tests failed. Run `uv run pytest` to see details."}'
  exit 0
fi

jq -n -c '{"decision": "approve", "systemMessage": "✓ All checks passed (format, lint, mypy, test)"}'
exit 0
