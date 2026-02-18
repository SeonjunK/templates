#!/bin/bash
# Run format, lint, and test on Stop

if [ ! -f "go.mod" ]; then
  exit 0
fi

UNFORMATTED=$(gofmt -l . 2>/dev/null | grep -v vendor/)
if [ -n "$UNFORMATTED" ]; then
  echo '{"systemMessage": "⚠ Format failed - unformatted files", "decision": "block", "reason": "Run `gofmt -w .` to fix formatting."}'
  exit 0
fi

LONG_LINES=$(go tool golines -l . 2>/dev/null | grep -v vendor/)
if [ -n "$LONG_LINES" ]; then
  echo '{"systemMessage": "⚠ Long lines detected", "decision": "block", "reason": "Run `go tool golines -w .` to fix long lines."}'
  exit 0
fi

if ! go tool golangci-lint run ./... >/dev/null 2>&1; then
  echo '{"systemMessage": "⚠ Lint failed", "decision": "block", "reason": "Lint failed. Run `go tool golangci-lint run ./...` to see details."}'
  exit 0
fi

if ! go test -race ./... >/dev/null 2>&1; then
  echo '{"systemMessage": "⚠ Tests failed", "decision": "block", "reason": "Tests failed. Run `go test -race ./...` to see details."}'
  exit 0
fi

echo '{"systemMessage": "✓ All checks passed (format, golines, lint, test)"}'
exit 0
