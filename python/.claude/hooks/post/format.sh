#!/bin/bash
# Format Python file after Write/Edit

HOOKS_DIR="$(cd "$(dirname "$0")/.." && pwd)"
. "$HOOKS_DIR/_lib/parse.sh"
. "$HOOKS_DIR/_lib/response.sh"

HOOK_INPUT=$(read_input)
FILE=$(parse_file_path "$HOOK_INPUT")

if [[ "$FILE" != *.py ]]; then
  exit 0
fi

cd "$CLAUDE_PROJECT_DIR"

if ! uv run ruff format "$FILE" >/dev/null 2>&1; then
  warn "⚠ Format failed for $FILE"
  exit 0
fi
