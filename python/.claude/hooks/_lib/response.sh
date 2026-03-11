#!/bin/sh
# Shared response helpers for Claude Code hooks.
# Source this file: . "$HOOKS_DIR/lib/response.sh"

# Output a PreToolUse deny response.
# Usage: deny "message"
deny() {
  jq -n -c --arg msg "$1" '{
    hookSpecificOutput: { permissionDecision: "deny" },
    systemMessage: $msg
  }'
}

# Output a PreToolUse/PostToolUse warning (no deny).
# Usage: warn "message"
warn() {
  jq -n -c --arg msg "$1" '{ systemMessage: $msg }'
}

# Output a Stop block response.
# Usage: block "reason" "message"
block() {
  jq -n -c --arg reason "$1" --arg msg "$2" '{
    decision: "block",
    reason: $reason,
    systemMessage: $msg
  }'
}

# Output a Stop approve response.
# Usage: approve "message"
approve() {
  jq -n -c --arg msg "$1" '{
    decision: "approve",
    systemMessage: $msg
  }'
}
