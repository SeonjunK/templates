#!/bin/bash
# Run format, lint, and test on Stop

if [ ! -f "Cargo.toml" ]; then
  exit 0
fi

UNFORMATTED=$(cargo fmt -- --check 2>&1 | grep -E "^Diff in" || true)
if [ -n "$UNFORMATTED" ]; then
  echo '{"systemMessage": "⚠ Format failed - unformatted files", "decision": "block", "reason": "Run `cargo fmt` to fix formatting."}'
  exit 0
fi

if ! cargo clippy -- -D warnings >/dev/null 2>&1; then
  echo '{"systemMessage": "⚠ Clippy failed", "decision": "block", "reason": "Clippy failed. Run `cargo clippy -- -D warnings` to see details."}'
  exit 0
fi

if ! cargo test >/dev/null 2>&1; then
  echo '{"systemMessage": "⚠ Tests failed", "decision": "block", "reason": "Tests failed. Run `cargo test` to see details."}'
  exit 0
fi

echo '{"systemMessage": "✓ All checks passed (format, clippy, test)"}'
exit 0
