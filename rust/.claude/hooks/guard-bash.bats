#!/usr/bin/env bats
# guard-bash.sh tests - dangerous bash command blocking hook
# Based on actual hook stdin format from logs

SCRIPT="$(cd "$(dirname "$BATS_TEST_FILENAME")" && pwd)/guard-bash.sh"

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
  local tool_input="$1"
  local permission_mode="${2:-default}"
  jq -c -n \
    --argjson tool_input "$tool_input" \
    --arg permission_mode "$permission_mode" \
    '{
      session_id: "test-session-id",
      transcript_path: "/tmp/test.jsonl",
      cwd: "/tmp/project",
      permission_mode: $permission_mode,
      hook_event_name: "PreToolUse",
      tool_name: "Bash",
      tool_input: $tool_input,
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

@test "allows when command is missing" {
  run sh "$SCRIPT" <<< '{}'
  assert_success
  refute_output
}

@test "allows when guard.json is absent" {
  rm -f "$TEST_TMPDIR/.claude/guard.json"
  run sh "$SCRIPT" <<< "$(_stdin '{"command": "rm -rf /"}')"
  assert_success
  refute_output
}

@test "blocks exact match command (default mode)" {
  _guard_config <<< '{"bash": {"blockedCommands": ["rm -rf /"], "blockedPatterns": []}}'
  run sh "$SCRIPT" <<< "$(_stdin '{"command": "rm -rf /"}' 'default')"
  assert_success
  assert_output "$(_expected_deny '⚠ Command blocked: rm -rf / is blocked by guard policy')"
}

@test "blocks exact match command (bypassPermissions mode)" {
  _guard_config <<< '{"bash": {"blockedCommands": ["rm -rf /"], "blockedPatterns": []}}'
  run sh "$SCRIPT" <<< "$(_stdin '{"command": "rm -rf /"}' 'bypassPermissions')"
  assert_success
  assert_output "$(_expected_deny '⚠ Command blocked: rm -rf / is blocked by guard policy')"
}

@test "blocks exact match command (plan mode)" {
  _guard_config <<< '{"bash": {"blockedCommands": ["rm -rf /"], "blockedPatterns": []}}'
  run sh "$SCRIPT" <<< "$(_stdin '{"command": "rm -rf /"}' 'plan')"
  assert_success
  assert_output "$(_expected_deny '⚠ Command blocked: rm -rf / is blocked by guard policy')"
}

@test "blocks exact match command (acceptEdits mode)" {
  _guard_config <<< '{"bash": {"blockedCommands": ["rm -rf /"], "blockedPatterns": []}}'
  run sh "$SCRIPT" <<< "$(_stdin '{"command": "rm -rf /"}' 'acceptEdits')"
  assert_success
  assert_output "$(_expected_deny '⚠ Command blocked: rm -rf / is blocked by guard policy')"
}

@test "allows non-exact-match command" {
  _guard_config <<< '{"bash": {"blockedCommands": ["rm -rf /"], "blockedPatterns": []}}'
  run sh "$SCRIPT" <<< "$(_stdin '{"command": "ls -la"}')"
  assert_success
  refute_output
}

@test "blocks command matching pattern with spaces" {
  _guard_config <<< '{"bash": {"blockedCommands": [], "blockedPatterns": ["git push --force"]}}'
  run sh "$SCRIPT" <<< "$(_stdin '{"command": "git push --force origin main", "description": "force push"}')"
  assert_success
  assert_output "$(_expected_deny '⚠ Command blocked: matches pattern git push --force')"
}

@test "allows command not matching pattern" {
  _guard_config <<< '{"bash": {"blockedCommands": [], "blockedPatterns": ["rm -rf"]}}'
  run sh "$SCRIPT" <<< "$(_stdin '{"command": "git status"}')"
  assert_success
  refute_output
}

@test "blocks command with description field (real format)" {
  _guard_config <<< '{"bash": {"blockedCommands": [], "blockedPatterns": ["drop table"]}}'
  run sh "$SCRIPT" <<< "$(_stdin '{"command": "echo drop table users | mysql", "description": "dangerous operation"}')"
  assert_success
  assert_output "$(_expected_deny '⚠ Command blocked: matches pattern drop table')"
}

@test "blocks with multiple blocked patterns" {
  _guard_config <<< '{"bash": {"blockedCommands": [], "blockedPatterns": ["rm -rf", "git push --force", "drop table"]}}'
  run sh "$SCRIPT" <<< "$(_stdin '{"command": "sudo rm -rf /var/log"}')"
  assert_success
  assert_output "$(_expected_deny '⚠ Command blocked: matches pattern rm -rf')"
}

@test "allows safe command with strict config" {
  _guard_config <<< '{"bash": {"blockedCommands": ["rm -rf /", ":(){ :|:& };:"], "blockedPatterns": ["git push --force", "drop table", "truncate table"]}}'
  run sh "$SCRIPT" <<< "$(_stdin '{"command": "go test ./...", "description": "run tests"}')"
  assert_success
  refute_output
}

@test "handles empty tool_input gracefully" {
  _guard_config <<< '{"bash": {"blockedCommands": ["rm -rf /"], "blockedPatterns": []}}'
  run sh "$SCRIPT" <<< '{"tool_input": {}}'
  assert_success
  refute_output
}

@test "handles malformed JSON gracefully" {
  _guard_config <<< '{"bash": {"blockedCommands": ["rm -rf /"], "blockedPatterns": []}}'
  run sh "$SCRIPT" <<< 'not valid json'
  assert_success
  refute_output
}
