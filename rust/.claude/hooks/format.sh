#!/bin/bash
# Format Rust file after Write/Edit

INPUT=$(cat)
FILE=$(echo "$INPUT" | jq -r '.tool_input.file_path // empty')

if [[ "$FILE" != *.rs ]]; then
  exit 0
fi

if ! rustfmt "$FILE" >/dev/null 2>&1; then
  jq -n -c --arg file "$FILE" '{"systemMessage": "⚠ Format failed for \($file)"}'
  exit 0
fi

exit 0
