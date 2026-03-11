#!/usr/bin/env bats
# verify.sh tests - Python session-end quality verification hook

SCRIPT="$(cd "$(dirname "$BATS_TEST_FILENAME")" && pwd)/../stop/verify.sh"

bats_load_library bats-support
bats_load_library bats-assert

setup() {
  TEST_TMPDIR="$(mktemp -d)"
  MOCK_BIN="$TEST_TMPDIR/bin"
  mkdir -p "$MOCK_BIN"
  export PATH="$MOCK_BIN:$PATH"
  export CLAUDE_PROJECT_DIR="$TEST_TMPDIR"
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

_mock_uv_pass() {
  printf '#!/bin/bash\nexit 0\n' > "$MOCK_BIN/uv"
  chmod +x "$MOCK_BIN/uv"
}

@test "blocks when ruff format fails" {
  printf '#!/bin/bash\nexit 1\n' > "$MOCK_BIN/uv"
  chmod +x "$MOCK_BIN/uv"
  run bash "$SCRIPT"
  assert_success
  assert_output "$(_expected_block 'Format failed' '⚠ Format check failed. Run `uv run ruff format .` to fix.')"
}

@test "blocks when ruff check fails" {
  cat > "$MOCK_BIN/uv" <<'EOF'
#!/bin/bash
if [[ "$3" == "format" && "$4" == "--check" ]]; then exit 0; fi
exit 1
EOF
  chmod +x "$MOCK_BIN/uv"
  run bash "$SCRIPT"
  assert_success
  assert_output "$(_expected_block 'Lint failed' '⚠ Lint failed. Run `uv run ruff check . --fix` to see details.')"
}

@test "blocks when mypy fails" {
  cat > "$MOCK_BIN/uv" <<'EOF'
#!/bin/bash
if [[ "$2" == "mypy" ]]; then exit 1; fi
exit 0
EOF
  chmod +x "$MOCK_BIN/uv"
  run bash "$SCRIPT"
  assert_success
  assert_output "$(_expected_block 'Type check failed' '⚠ Type check failed. Run `uv run mypy src` to see details.')"
}

@test "blocks when pytest fails" {
  cat > "$MOCK_BIN/uv" <<'EOF'
#!/bin/bash
if [[ "$2" == "pytest" ]]; then exit 1; fi
exit 0
EOF
  chmod +x "$MOCK_BIN/uv"
  run bash "$SCRIPT"
  assert_success
  assert_output "$(_expected_block 'Tests failed' '⚠ Tests failed. Run `uv run pytest` to see details.')"
}

@test "approves when all checks pass" {
  _mock_uv_pass
  run bash "$SCRIPT"
  assert_success
  assert_output "$(_expected_approve '✓ All checks passed (format, lint, mypy, test)')"
}
