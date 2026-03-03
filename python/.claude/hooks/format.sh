#!/bin/bash
# Format Python file after Write/Edit

cd "$CLAUDE_PROJECT_DIR"

INPUT=$(cat)
FILE=$(echo "$INPUT" | jq -r '.tool_input.file_path // empty')

if [[ "$FILE" != *.py ]]; then
  exit 0
fi

if ! uv run ruff format "$FILE" >/dev/null 2>&1; then
  jq -n -c --arg file "$FILE" '{"systemMessage": "⚠ Format failed for \($file)"}'
  exit 0
fi

exit 0
