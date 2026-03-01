#!/usr/bin/env bats
# format.sh tests - Rust file auto-format hook

SCRIPT="$(cd "$(dirname "$BATS_TEST_FILENAME")" && pwd)/format.sh"

bats_load_library bats-support
bats_load_library bats-assert

setup() {
  TEST_TMPDIR="$(mktemp -d)"
  MOCK_BIN="$TEST_TMPDIR/bin"
  mkdir -p "$MOCK_BIN"
  export PATH="$MOCK_BIN:$PATH"
}

teardown() {
  rm -rf "$TEST_TMPDIR"
}

_expected_message() {
  local message="$1"
  jq -c -n --arg message "$message" '{"systemMessage": $message}'
}

_mock_rustfmt() {
  printf '#!/bin/bash\nexit %d\n' "$1" > "$MOCK_BIN/rustfmt"
  chmod +x "$MOCK_BIN/rustfmt"
}

@test "skips when file_path is missing" {
  run bash "$SCRIPT" <<< '{}'
  assert_success
  refute_output
}

@test "skips non-.rs files" {
  run bash "$SCRIPT" <<< '{"tool_input": {"file_path": "main.go"}}'
  assert_success
  refute_output
}

@test "reports when rustfmt fails on .rs file" {
  _mock_rustfmt 1
  run bash "$SCRIPT" <<< '{"tool_input": {"file_path": "src/main.rs"}}'
  assert_success
  assert_output "$(_expected_message '⚠ Format failed for src/main.rs')"
}

@test "produces no output when rustfmt succeeds on .rs file" {
  _mock_rustfmt 0
  run bash "$SCRIPT" <<< '{"tool_input": {"file_path": "src/main.rs"}}'
  assert_success
  refute_output
}
