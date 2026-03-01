#!/usr/bin/env bats
# verify.sh tests - Rust session-end quality verification hook

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

_mock_cargo_pass() {
  printf '#!/bin/bash\nexit 0\n' > "$MOCK_BIN/cargo"
  chmod +x "$MOCK_BIN/cargo"
}

@test "exits silently when Cargo.toml is absent" {
  run bash -c "cd '$WORK_DIR' && bash '$SCRIPT'"
  assert_success
  refute_output
}

@test "blocks when cargo fmt detects unformatted files" {
  touch "$WORK_DIR/Cargo.toml"
  cat > "$MOCK_BIN/cargo" <<'EOF'
#!/bin/bash
if [[ "$1" == "fmt" ]]; then
  echo "Diff in src/main.rs:"
  exit 1
fi
exit 0
EOF
  chmod +x "$MOCK_BIN/cargo"
  run bash -c "cd '$WORK_DIR' && bash '$SCRIPT'"
  assert_success
  assert_output "$(_expected_block 'Unformatted files' '⚠ Format failed - unformatted files. Run `cargo fmt` to fix.')"
}

@test "blocks when cargo clippy fails" {
  touch "$WORK_DIR/Cargo.toml"
  cat > "$MOCK_BIN/cargo" <<'EOF'
#!/bin/bash
if [[ "$1" == "clippy" ]]; then exit 1; fi
exit 0
EOF
  chmod +x "$MOCK_BIN/cargo"
  run bash -c "cd '$WORK_DIR' && bash '$SCRIPT'"
  assert_success
  assert_output "$(_expected_block 'Clippy failed' '⚠ Clippy failed. Run `cargo clippy -- -D warnings` to see details.')"
}

@test "blocks when cargo test fails" {
  touch "$WORK_DIR/Cargo.toml"
  cat > "$MOCK_BIN/cargo" <<'EOF'
#!/bin/bash
if [[ "$1" == "test" ]]; then exit 1; fi
exit 0
EOF
  chmod +x "$MOCK_BIN/cargo"
  run bash -c "cd '$WORK_DIR' && bash '$SCRIPT'"
  assert_success
  assert_output "$(_expected_block 'Tests failed' '⚠ Tests failed. Run `cargo test` to see details.')"
}

@test "approves when all checks pass" {
  touch "$WORK_DIR/Cargo.toml"
  _mock_cargo_pass
  run bash -c "cd '$WORK_DIR' && bash '$SCRIPT'"
  assert_success
  assert_output "$(_expected_approve '✓ All checks passed (format, clippy, test)')"
}
