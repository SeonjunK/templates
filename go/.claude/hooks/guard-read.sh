#!/bin/sh
# Guard hook for Read tool - block sensitive file access

set -e

INPUT=$(cat)
FILE=$(echo "$INPUT" | jq -r '.tool_input.file_path // empty' 2>/dev/null) || exit 0
GUARD_CONFIG="${CLAUDE_PROJECT_DIR:-}/.claude/guard.json"

# Exit early if no file or no config
if [ -z "$FILE" ] || [ -z "${CLAUDE_PROJECT_DIR:-}" ] || [ ! -f "$GUARD_CONFIG" ]; then
  exit 0
fi

BASENAME=$(basename "$FILE")

# Temporary file for pattern matching
TMP_FILE=$(mktemp)
trap 'rm -f "$TMP_FILE"' EXIT

# Check blocked patterns
jq -r '.read.blockedPatterns[]? // empty' "$GUARD_CONFIG" 2>/dev/null > "$TMP_FILE"
while IFS= read -r pattern; do
  [ -z "$pattern" ] && continue
  case "$BASENAME" in
    $pattern)
      jq -n -c \
        --arg basename "$BASENAME" \
        --arg pattern "$pattern" \
        --arg file "$FILE" \
        '{
          hookSpecificOutput: {
            permissionDecision: "deny"
          },
          systemMessage: "⚠ File access blocked: \($basename) (matched pattern: \($pattern))"
        }'
      exit 0
      ;;
  esac
done < "$TMP_FILE"

exit 0
