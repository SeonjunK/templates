#!/bin/bash
# Run format, lint, type check, and test on Stop

if ! uv run ruff format . >/dev/null 2>&1; then
  echo '{"systemMessage": "⚠ Format failed", "decision": "block", "reason": "Format failed. Run `uv run ruff format .` to see details."}'
  exit 0
fi

if ! uv run ruff check . --fix >/dev/null 2>&1; then
  echo '{"systemMessage": "⚠ Lint failed", "decision": "block", "reason": "Lint failed. Run `uv run ruff check . --fix` to see details."}'
  exit 0
fi

if ! uv run mypy src >/dev/null 2>&1; then
  echo '{"systemMessage": "⚠ Type check failed - Run `uv run mypy src` to see details", "decision": "block", "reason": "Type check failed. Run `uv run mypy src` to see details."}'
  exit 0
fi

if ! uv run pytest --cov=src --cov-report=xml >/dev/null 2>&1; then
  echo '{"systemMessage": "⚠ Tests failed - Run `uv run pytest` to see details", "decision": "block", "reason": "Tests failed. Run `uv run pytest` to see details."}'
  exit 0
fi

echo '{"systemMessage": "✓ All checks passed (format, lint, mypy, test)"}'
exit 0
