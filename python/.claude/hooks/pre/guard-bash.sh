#!/bin/sh
# Guard hook for Bash tool - block dangerous commands
set -e

HOOKS_DIR="$(cd "$(dirname "$0")/.." && pwd)"
. "$HOOKS_DIR/_lib/parse.sh"
. "$HOOKS_DIR/_lib/response.sh"

HOOK_INPUT=$(read_input)
CMD=$(parse_command "$HOOK_INPUT") || exit 0
[ -z "$CMD" ] && exit 0

CONFIG=$(guard_config_path)
[ -z "$CONFIG" ] && exit 0

# Check exact blocked commands
guard_patterns "$CONFIG" ".bash.blockedCommands" | while IFS= read -r blocked; do
  [ -z "$blocked" ] && continue
  if [ "$CMD" = "$blocked" ]; then
    deny "⚠ Command blocked: $CMD is blocked by guard policy"
    exit 0
  fi
done

# Check pattern matches (substring)
guard_patterns "$CONFIG" ".bash.blockedPatterns" | while IFS= read -r pattern; do
  [ -z "$pattern" ] && continue
  case "$CMD" in
    *"$pattern"*)
      deny "⚠ Command blocked: matches pattern $pattern"
      exit 0
      ;;
  esac
done
