#!/bin/bash
# Run format check, lint, type check, and test on Stop

HOOKS_DIR="$(cd "$(dirname "$0")/.." && pwd)"
. "$HOOKS_DIR/_lib/response.sh"

cd "$CLAUDE_PROJECT_DIR"

if ! uv run ruff format --check . >/dev/null 2>&1; then
  block "Format failed" "⚠ Format check failed. Run \`uv run ruff format .\` to fix."
  exit 0
fi

if ! uv run ruff check . >/dev/null 2>&1; then
  block "Lint failed" "⚠ Lint failed. Run \`uv run ruff check . --fix\` to see details."
  exit 0
fi

if ! uv run mypy src >/dev/null 2>&1; then
  block "Type check failed" "⚠ Type check failed. Run \`uv run mypy src\` to see details."
  exit 0
fi

if ! uv run pytest --cov=src --cov-report=xml >/dev/null 2>&1; then
  block "Tests failed" "⚠ Tests failed. Run \`uv run pytest\` to see details."
  exit 0
fi

approve "✓ All checks passed (format, lint, mypy, test)"
