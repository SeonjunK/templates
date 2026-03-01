#!/bin/bash
# Run format, lint, and test on Stop

if [ ! -f "Cargo.toml" ]; then
  exit 0
fi

UNFORMATTED=$(cargo fmt -- --check 2>&1 | grep -E "^Diff in" || true)
if [ -n "$UNFORMATTED" ]; then
  jq -n -c '{"decision": "block", "reason": "Unformatted files", "systemMessage": "⚠ Format failed - unformatted files. Run `cargo fmt` to fix."}'
  exit 0
fi

if ! cargo clippy -- -D warnings >/dev/null 2>&1; then
  jq -n -c '{"decision": "block", "reason": "Clippy failed", "systemMessage": "⚠ Clippy failed. Run `cargo clippy -- -D warnings` to see details."}'
  exit 0
fi

if ! cargo test >/dev/null 2>&1; then
  jq -n -c '{"decision": "block", "reason": "Tests failed", "systemMessage": "⚠ Tests failed. Run `cargo test` to see details."}'
  exit 0
fi

jq -n -c '{"decision": "approve", "systemMessage": "✓ All checks passed (format, clippy, test)"}'
exit 0
