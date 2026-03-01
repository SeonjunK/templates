#!/usr/bin/env bats
# format.sh tests - Python file auto-format hook

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

_mock_uv() {
  printf '#!/bin/bash\nexit %d\n' "$1" > "$MOCK_BIN/uv"
  chmod +x "$MOCK_BIN/uv"
}

@test "skips when file_path is missing" {
  run bash "$SCRIPT" <<< '{}'
  assert_success
  refute_output
}

@test "skips non-.py files" {
  run bash "$SCRIPT" <<< '{"tool_input": {"file_path": "main.go"}}'
  assert_success
  refute_output
}

@test "reports when ruff format fails on .py file" {
  _mock_uv 1
  run bash "$SCRIPT" <<< '{"tool_input": {"file_path": "main.py"}}'
  assert_success
  assert_output "$(_expected_message '⚠ Format failed for main.py')"
}

@test "produces no output when ruff format succeeds on .py file" {
  _mock_uv 0
  run bash "$SCRIPT" <<< '{"tool_input": {"file_path": "main.py"}}'
  assert_success
  refute_output
}
