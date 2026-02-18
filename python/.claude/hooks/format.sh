#!/bin/bash
# Format Python file after Write/Edit

INPUT=$(cat)
FILE=$(echo "$INPUT" | jq -r '.tool_input.file_path // empty')

if [[ "$FILE" != *.py ]]; then
  exit 0
fi

if ! uv run ruff format "$FILE" >/dev/null 2>&1; then
  echo "{\"systemMessage\": \"⚠ Format failed for $FILE\", \"decision\": \"block\", \"reason\": \"Format failed. Run \`uv run ruff format $FILE\` to see details.\"}"
  exit 0
fi

exit 0
