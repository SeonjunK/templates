#!/usr/bin/env bats
# verify.sh tests - Go session-end quality verification hook

SCRIPT="$(cd "$(dirname "$BATS_TEST_FILENAME")" && pwd)/verify.sh"

bats_load_library bats-support
bats_load_library bats-assert

setup() {
  TEST_TMPDIR="$(mktemp -d)"
  MOCK_BIN="$TEST_TMPDIR/bin"
  WORK_DIR="$TEST_TMPDIR/work"
  mkdir -p "$MOCK_BIN" "$WORK_DIR"
  export PATH="$MOCK_BIN:$PATH"
}

teardown() {
  rm -rf "$TEST_TMPDIR"
}

_expected_block() {
  local reason="$1"
  local message="$2"
  jq -c -n --arg reason "$reason" --arg message "$message" \
    '{"decision": "block", "reason": $reason, "systemMessage": $message}'
}

_expected_approve() {
  local message="$1"
  jq -c -n --arg message "$message" '{"decision": "approve", "systemMessage": $message}'
}

_mock_gofmt_pass() {
  printf '#!/bin/bash\nexit 0\n' > "$MOCK_BIN/gofmt"
  chmod +x "$MOCK_BIN/gofmt"
}

_mock_gofmt_unformatted() {
  printf '#!/bin/bash\necho "main.go"\nexit 0\n' > "$MOCK_BIN/gofmt"
  chmod +x "$MOCK_BIN/gofmt"
}

_mock_go_pass() {
  printf '#!/bin/bash\nexit 0\n' > "$MOCK_BIN/go"
  chmod +x "$MOCK_BIN/go"
}

_mock_all_pass() {
  _mock_gofmt_pass
  _mock_go_pass
}

@test "exits silently when go.mod is absent" {
  run bash -c "cd '$WORK_DIR' && bash '$SCRIPT'"
  assert_success
  refute_output
}

@test "blocks when unformatted files exist" {
  touch "$WORK_DIR/go.mod"
  _mock_gofmt_unformatted
  run bash -c "cd '$WORK_DIR' && bash '$SCRIPT'"
  assert_success
  assert_output "$(_expected_block 'Unformatted files' '⚠ Format failed - unformatted files. Run `gofmt -w .` to fix.')"
}

@test "blocks when long lines are detected" {
  touch "$WORK_DIR/go.mod"
  _mock_gofmt_pass
  cat > "$MOCK_BIN/go" <<'EOF'
#!/bin/bash
if [[ "$1" == "tool" && "$2" == "golines" && "$3" == "-l" ]]; then
  echo "main.go"
fi
exit 0
EOF
  chmod +x "$MOCK_BIN/go"
  run bash -c "cd '$WORK_DIR' && bash '$SCRIPT'"
  assert_success
  assert_output "$(_expected_block 'Long lines detected' '⚠ Long lines detected. Run `go tool golines -w .` to fix.')"
}

@test "blocks when lint fails" {
  touch "$WORK_DIR/go.mod"
  _mock_gofmt_pass
  cat > "$MOCK_BIN/go" <<'EOF'
#!/bin/bash
if [[ "$1" == "tool" && "$2" == "golangci-lint" ]]; then
  exit 1
fi
exit 0
EOF
  chmod +x "$MOCK_BIN/go"
  run bash -c "cd '$WORK_DIR' && bash '$SCRIPT'"
  assert_success
  assert_output "$(_expected_block 'Lint failed' '⚠ Lint failed. Run `go tool golangci-lint run ./...` to see details.')"
}

@test "blocks when tests fail" {
  touch "$WORK_DIR/go.mod"
  _mock_gofmt_pass
  cat > "$MOCK_BIN/go" <<'EOF'
#!/bin/bash
if [[ "$1" == "test" ]]; then
  exit 1
fi
exit 0
EOF
  chmod +x "$MOCK_BIN/go"
  run bash -c "cd '$WORK_DIR' && bash '$SCRIPT'"
  assert_success
  assert_output "$(_expected_block 'Tests failed' '⚠ Tests failed. Run `go test -race ./...` to see details.')"
}

@test "approves when all checks pass" {
  touch "$WORK_DIR/go.mod"
  _mock_all_pass
  run bash -c "cd '$WORK_DIR' && bash '$SCRIPT'"
  assert_success
  assert_output "$(_expected_approve '✓ All checks passed (format, golines, lint, test)')"
}
