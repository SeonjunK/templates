#!/bin/bash
# Run format, lint, and test on Stop

if [ ! -f "go.mod" ]; then
  exit 0
fi

UNFORMATTED=$(gofmt -l . 2>/dev/null | grep -v vendor/)
if [ -n "$UNFORMATTED" ]; then
  jq -n -c '{"decision": "block", "reason": "Unformatted files", "systemMessage": "⚠ Format failed - unformatted files. Run `gofmt -w .` to fix."}'
  exit 0
fi

LONG_LINES=$(go tool golines -l . 2>/dev/null | grep -v vendor/)
if [ -n "$LONG_LINES" ]; then
  jq -n -c '{"decision": "block", "reason": "Long lines detected", "systemMessage": "⚠ Long lines detected. Run `go tool golines -w .` to fix."}'
  exit 0
fi

if ! go tool golangci-lint run ./... >/dev/null 2>&1; then
  jq -n -c '{"decision": "block", "reason": "Lint failed", "systemMessage": "⚠ Lint failed. Run `go tool golangci-lint run ./...` to see details."}'
  exit 0
fi

if ! go test -race ./... >/dev/null 2>&1; then
  jq -n -c '{"decision": "block", "reason": "Tests failed", "systemMessage": "⚠ Tests failed. Run `go test -race ./...` to see details."}'
  exit 0
fi

jq -n -c '{"decision": "approve", "systemMessage": "✓ All checks passed (format, golines, lint, test)"}'
exit 0
