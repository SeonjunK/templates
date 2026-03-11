#!/bin/sh
# Guard hook for Bash tool - delegates to prebuilt hook binary
HOOK_BIN="${CLAUDE_PROJECT_DIR}/bin/hook"
[ -f "$HOOK_BIN" ] || exit 0
exec "$HOOK_BIN" guard-bash
