#!/usr/bin/env bats
# guard-write.sh tests - sensitive file write blocking hook
# Based on actual hook stdin format from logs

SCRIPT="$(cd "$(dirname "$BATS_TEST_FILENAME")" && pwd)/guard-write.sh"

bats_load_library bats-support
bats_load_library bats-assert

setup() {
  TEST_TMPDIR="$(mktemp -d)"
  mkdir -p "$TEST_TMPDIR/.claude"
  export CLAUDE_PROJECT_DIR="$TEST_TMPDIR"
}

teardown() {
  rm -rf "$TEST_TMPDIR"
}

_guard_config() {
  cat > "$TEST_TMPDIR/.claude/guard.json"
}

# Real stdin format from actual logs
_stdin() {
  local file_path="$1"
  local permission_mode="${2:-default}"
  jq -c -n \
    --arg file_path "$file_path" \
    --arg permission_mode "$permission_mode" \
    '{
      session_id: "test-session-id",
      transcript_path: "/tmp/test.jsonl",
      cwd: "/tmp/project",
      permission_mode: $permission_mode,
      hook_event_name: "PreToolUse",
      tool_name: "Write",
      tool_input: {file_path: $file_path},
      tool_use_id: "test-use-id"
    }'
}

# Expected deny response
_expected_deny() {
  local message="$1"
  jq -c -n \
    --arg message "$message" \
    '{
      hookSpecificOutput: {
        permissionDecision: "deny"
      },
      systemMessage: $message
    }'
}

@test "allows when file_path is missing" {
  run sh "$SCRIPT" <<< '{}'
  assert_success
  refute_output
}

@test "allows when CLAUDE_PROJECT_DIR is unset" {
  run env -u CLAUDE_PROJECT_DIR sh "$SCRIPT" <<< "$(_stdin '/project/.env')"
  assert_success
  refute_output
}

@test "allows when guard.json has no write key" {
  _guard_config <<< '{}'
  run sh "$SCRIPT" <<< "$(_stdin '/project/.env')"
  assert_success
  refute_output
}

@test "allows when guard.json is absent" {
  rm -f "$TEST_TMPDIR/.claude/guard.json"
  run sh "$SCRIPT" <<< "$(_stdin '/project/.env')"
  assert_success
  refute_output
}

@test "blocks file write matching blocked pattern (default mode)" {
  _guard_config <<< '{"write": {"blockedPatterns": [".env"]}}'
  run sh "$SCRIPT" <<< "$(_stdin '/project/.env' 'default')"
  assert_success
  assert_output "$(_expected_deny '⚠ File write blocked: .env (matched pattern: .env)')"
}

@test "blocks file write matching blocked pattern (bypassPermissions mode)" {
  _guard_config <<< '{"write": {"blockedPatterns": [".env"]}}'
  run sh "$SCRIPT" <<< "$(_stdin '/project/.env' 'bypassPermissions')"
  assert_success
  assert_output "$(_expected_deny '⚠ File write blocked: .env (matched pattern: .env)')"
}

@test "blocks file write matching blocked pattern (plan mode)" {
  _guard_config <<< '{"write": {"blockedPatterns": [".env"]}}'
  run sh "$SCRIPT" <<< "$(_stdin '/project/.env' 'plan')"
  assert_success
  assert_output "$(_expected_deny '⚠ File write blocked: .env (matched pattern: .env)')"
}

@test "blocks file write matching blocked pattern (acceptEdits mode)" {
  _guard_config <<< '{"write": {"blockedPatterns": [".env"]}}'
  run sh "$SCRIPT" <<< "$(_stdin '/project/.env' 'acceptEdits')"
  assert_success
  assert_output "$(_expected_deny '⚠ File write blocked: .env (matched pattern: .env)')"
}

@test "blocks file write matching wildcard pattern" {
  _guard_config <<< '{"write": {"blockedPatterns": ["*.key"]}}'
  run sh "$SCRIPT" <<< "$(_stdin '/project/private.key')"
  assert_success
  assert_output "$(_expected_deny '⚠ File write blocked: private.key (matched pattern: *.key)')"
}

@test "allows file write not matching any blocked pattern" {
  _guard_config <<< '{"write": {"blockedPatterns": ["*.env"]}}'
  run sh "$SCRIPT" <<< "$(_stdin '/project/main.go')"
  assert_success
  refute_output
}

@test "blocks multiple patterns (.env*, *.pem, *.key)" {
  _guard_config <<< '{"write": {"blockedPatterns": [".env*", "*.pem", "*.key"]}}'
  run sh "$SCRIPT" <<< "$(_stdin '/project/cert.pem')"
  assert_success
  assert_output "$(_expected_deny '⚠ File write blocked: cert.pem (matched pattern: *.pem)')"
}

@test "allows similar but not matching pattern" {
  _guard_config <<< '{"write": {"blockedPatterns": [".env"]}}'
  run sh "$SCRIPT" <<< "$(_stdin '/project/.envrc')"
  assert_success
  refute_output
}

@test "handles empty tool_input gracefully" {
  _guard_config <<< '{"write": {"blockedPatterns": [".env"]}}'
  run sh "$SCRIPT" <<< '{"tool_input": {}}'
  assert_success
  refute_output
}
