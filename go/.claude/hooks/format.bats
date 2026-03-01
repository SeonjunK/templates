#!/usr/bin/env bats
# format.sh tests - Go file auto-format hook

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

_mock_gofmt() {
  printf '#!/bin/bash\nexit %d\n' "$1" > "$MOCK_BIN/gofmt"
  chmod +x "$MOCK_BIN/gofmt"
}

_mock_go() {
  printf '#!/bin/bash\nexit %d\n' "$1" > "$MOCK_BIN/go"
  chmod +x "$MOCK_BIN/go"
}

@test "skips when file_path is missing" {
  run bash "$SCRIPT" <<< '{}'
  assert_success
  refute_output
}

@test "skips non-.go files" {
  run bash "$SCRIPT" <<< '{"tool_input": {"file_path": "main.py"}}'
  assert_success
  refute_output
}

@test "reports when gofmt fails on .go file" {
  _mock_gofmt 1
  run bash "$SCRIPT" <<< '{"tool_input": {"file_path": "main.go"}}'
  assert_success
  assert_output "$(_expected_message '⚠ Format failed for main.go')"
}

@test "reports when golines fails on .go file" {
  _mock_gofmt 0
  _mock_go 1
  run bash "$SCRIPT" <<< '{"tool_input": {"file_path": "main.go"}}'
  assert_success
  assert_output "$(_expected_message '⚠ golines failed for main.go')"
}

@test "produces no output when all formatters succeed on .go file" {
  _mock_gofmt 0
  _mock_go 0
  run bash "$SCRIPT" <<< '{"tool_input": {"file_path": "main.go"}}'
  assert_success
  refute_output
}
