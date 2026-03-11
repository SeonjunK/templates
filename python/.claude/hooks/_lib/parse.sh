#!/bin/sh
# Shared parsing helpers for Claude Code hooks.
# Source this file: . "$HOOKS_DIR/lib/parse.sh"

# Read stdin JSON and extract a field from tool_input.
# Usage: HOOK_INPUT=$(read_input)
read_input() {
  cat
}

# Extract file_path from hook input JSON.
# Usage: FILE=$(parse_file_path "$HOOK_INPUT")
parse_file_path() {
  echo "$1" | jq -r '.tool_input.file_path // empty' 2>/dev/null
}

# Extract command from hook input JSON.
# Usage: CMD=$(parse_command "$HOOK_INPUT")
parse_command() {
  echo "$1" | jq -r '.tool_input.command // empty' 2>/dev/null
}

# Load guard.json config path. Returns empty if unavailable.
# Usage: CONFIG=$(guard_config_path)
guard_config_path() {
  local config="${CLAUDE_PROJECT_DIR:-}/.claude/guard.json"
  if [ -z "${CLAUDE_PROJECT_DIR:-}" ] || [ ! -f "$config" ]; then
    echo ""
    return
  fi
  echo "$config"
}

# Read blocked patterns from guard.json for a given section.
# Usage: patterns=$(guard_patterns "$config" ".read.blockedPatterns")
guard_patterns() {
  jq -r "$2[]? // empty" "$1" 2>/dev/null
}
