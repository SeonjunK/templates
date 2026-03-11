#!/bin/sh
# Guard hook for Read tool - block sensitive file access
set -e

HOOKS_DIR="$(cd "$(dirname "$0")/.." && pwd)"
. "$HOOKS_DIR/_lib/parse.sh"
. "$HOOKS_DIR/_lib/response.sh"

HOOK_INPUT=$(read_input)
FILE=$(parse_file_path "$HOOK_INPUT") || exit 0
[ -z "$FILE" ] && exit 0

CONFIG=$(guard_config_path)
[ -z "$CONFIG" ] && exit 0

BASENAME=$(basename "$FILE")
guard_patterns "$CONFIG" ".read.blockedPatterns" | while IFS= read -r pattern; do
  [ -z "$pattern" ] && continue
  case "$BASENAME" in
    $pattern)
      deny "⚠ File access blocked: $BASENAME (matched pattern: $pattern)"
      exit 0
      ;;
  esac
done
