#!/bin/sh
# Guard hook for Bash tool - delegates to compiled hook binary
HOOK_BIN="${CLAUDE_PROJECT_DIR}/bin/hook"
[ -f "$HOOK_BIN" ] || go build -o "$HOOK_BIN" "${CLAUDE_PROJECT_DIR}/cmd/hook" 2>/dev/null || exit 0
exec "$HOOK_BIN" guard-bash
