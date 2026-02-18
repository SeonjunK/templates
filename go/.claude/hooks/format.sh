#!/bin/bash
# Format Go file after Write/Edit

INPUT=$(cat)
FILE=$(echo "$INPUT" | jq -r '.tool_input.file_path // empty')

if [[ "$FILE" != *.go ]]; then
  exit 0
fi

if ! gofmt -w "$FILE" >/dev/null 2>&1; then
  echo "{\"systemMessage\": \"⚠ Format failed for $FILE\", \"decision\": \"block\", \"reason\": \"Format failed. Run \`gofmt -w $FILE\` to see details.\"}"
  exit 0
fi

if ! go tool golines -w "$FILE" >/dev/null 2>&1; then
  echo "{\"systemMessage\": \"⚠ golines failed for $FILE\", \"decision\": \"block\", \"reason\": \"golines failed. Run \`go tool golines -w $FILE\` to see details.\"}"
  exit 0
fi

exit 0
