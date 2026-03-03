#!/bin/sh
# Guard hook for Bash tool - block dangerous commands

cd "$CLAUDE_PROJECT_DIR"
set -e

INPUT=$(cat)
CMD=$(echo "$INPUT" | jq -r '.tool_input.command // empty' 2>/dev/null) || exit 0
GUARD_CONFIG="${CLAUDE_PROJECT_DIR:-}/.claude/guard.json"

# Exit early if no command or no config
if [ -z "$CMD" ] || [ -z "${CLAUDE_PROJECT_DIR:-}" ] || [ ! -f "$GUARD_CONFIG" ]; then
  exit 0
fi

# Temporary file for pattern matching
TMP_FILE=$(mktemp)
trap 'rm -f "$TMP_FILE"' EXIT

# Check exact blocked commands
jq -r '.bash.blockedCommands[] // empty' "$GUARD_CONFIG" 2>/dev/null > "$TMP_FILE"
while IFS= read -r blocked; do
  [ -z "$blocked" ] && continue
  if [ "$CMD" = "$blocked" ]; then
    jq -n -c \
      --arg cmd "$CMD" \
      '{
        hookSpecificOutput: {
          permissionDecision: "deny"
        },
        systemMessage: "⚠ Command blocked: \($cmd) is blocked by guard policy"
      }'
    exit 0
  fi
done < "$TMP_FILE"

# Check pattern matches (substring)
jq -r '.bash.blockedPatterns[] // empty' "$GUARD_CONFIG" 2>/dev/null > "$TMP_FILE"
while IFS= read -r pattern; do
  [ -z "$pattern" ] && continue
  case "$CMD" in
    *"$pattern"*)
      jq -n -c \
        --arg pattern "$pattern" \
        '{
          hookSpecificOutput: {
            permissionDecision: "deny"
          },
          systemMessage: "⚠ Command blocked: matches pattern \($pattern)"
        }'
      exit 0
      ;;
  esac
done < "$TMP_FILE"

exit 0
