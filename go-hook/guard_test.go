package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupGuardEnv creates a temp dir with guard.json and sets CLAUDE_PROJECT_DIR.
func setupGuardEnv(t *testing.T, guardJSON string) string {
	t.Helper()
	dir := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(dir, ".claude"), 0o755))
	if guardJSON != "" {
		require.NoError(t, os.WriteFile(
			filepath.Join(dir, ".claude", "guard.json"),
			[]byte(guardJSON),
			0o644,
		))
	}
	t.Setenv("CLAUDE_PROJECT_DIR", dir)
	return dir
}

// fakeStdin replaces os.Stdin with data and restores it after the test.
func fakeStdin(t *testing.T, data string) {
	t.Helper()
	r, w, err := os.Pipe()
	require.NoError(t, err)
	_, err = w.WriteString(data)
	require.NoError(t, err)
	w.Close()
	old := os.Stdin
	os.Stdin = r
	t.Cleanup(func() { os.Stdin = old })
}

// --- guard-read tests ---

func TestGuardRead_AllowsWhenFilePathMissing(t *testing.T) {
	setupGuardEnv(t, `{"read":{"blockedPatterns":[".env"]}}`)
	fakeStdin(t, `{}`)
	out := captureStdout(t, func() { _ = guardRead() })
	assert.Empty(t, out)
}

func TestGuardRead_AllowsWhenProjectDirUnset(t *testing.T) {
	t.Setenv("CLAUDE_PROJECT_DIR", "")
	fakeStdin(t, `{"tool_input":{"file_path":"/project/.env"}}`)
	out := captureStdout(t, func() { _ = guardRead() })
	assert.Empty(t, out)
}

func TestGuardRead_AllowsWhenGuardJSONAbsent(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("CLAUDE_PROJECT_DIR", dir)
	fakeStdin(t, `{"tool_input":{"file_path":"/project/.env"}}`)
	out := captureStdout(t, func() { _ = guardRead() })
	assert.Empty(t, out)
}

func TestGuardRead_AllowsWhenNoReadKey(t *testing.T) {
	setupGuardEnv(t, `{}`)
	fakeStdin(t, `{"tool_input":{"file_path":"/project/.env"}}`)
	out := captureStdout(t, func() { _ = guardRead() })
	assert.Empty(t, out)
}

func TestGuardRead_BlocksMatchingPattern(t *testing.T) {
	setupGuardEnv(t, `{"read":{"blockedPatterns":[".env"]}}`)
	fakeStdin(t, `{"tool_input":{"file_path":"/project/.env"}}`)
	out := captureStdout(t, func() { _ = guardRead() })

	var resp PreToolResponse
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &resp))
	assert.Equal(t, "deny", resp.HookSpecificOutput.PermissionDecision)
	assert.Contains(t, resp.SystemMessage, "File access blocked: .env")
}

func TestGuardRead_BlocksWildcardPattern(t *testing.T) {
	setupGuardEnv(t, `{"read":{"blockedPatterns":["*.key"]}}`)
	fakeStdin(t, `{"tool_input":{"file_path":"/project/private.key"}}`)
	out := captureStdout(t, func() { _ = guardRead() })

	var resp PreToolResponse
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &resp))
	assert.Equal(t, "deny", resp.HookSpecificOutput.PermissionDecision)
	assert.Contains(t, resp.SystemMessage, "private.key")
}

func TestGuardRead_AllowsNonMatchingFile(t *testing.T) {
	setupGuardEnv(t, `{"read":{"blockedPatterns":["*.env"]}}`)
	fakeStdin(t, `{"tool_input":{"file_path":"/project/main.go"}}`)
	out := captureStdout(t, func() { _ = guardRead() })
	assert.Empty(t, out)
}

func TestGuardRead_BlocksMultiplePatterns(t *testing.T) {
	setupGuardEnv(t, `{"read":{"blockedPatterns":[".env*","*.pem","*.key"]}}`)
	fakeStdin(t, `{"tool_input":{"file_path":"/project/.env.production"}}`)
	out := captureStdout(t, func() { _ = guardRead() })

	var resp PreToolResponse
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &resp))
	assert.Contains(t, resp.SystemMessage, ".env.production")
}

func TestGuardRead_AllowsSimilarButNotMatching(t *testing.T) {
	setupGuardEnv(t, `{"read":{"blockedPatterns":[".env"]}}`)
	fakeStdin(t, `{"tool_input":{"file_path":"/project/.envrc"}}`)
	out := captureStdout(t, func() { _ = guardRead() })
	assert.Empty(t, out)
}

// --- guard-write tests ---

func TestGuardWrite_BlocksMatchingPattern(t *testing.T) {
	setupGuardEnv(t, `{"write":{"blockedPatterns":[".env"]}}`)
	fakeStdin(t, `{"tool_input":{"file_path":"/project/.env"}}`)
	out := captureStdout(t, func() { _ = guardWrite() })

	var resp PreToolResponse
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &resp))
	assert.Equal(t, "deny", resp.HookSpecificOutput.PermissionDecision)
	assert.Contains(t, resp.SystemMessage, "File write blocked: .env")
}

func TestGuardWrite_BlocksWildcardPattern(t *testing.T) {
	setupGuardEnv(t, `{"write":{"blockedPatterns":["*.pem"]}}`)
	fakeStdin(t, `{"tool_input":{"file_path":"/project/cert.pem"}}`)
	out := captureStdout(t, func() { _ = guardWrite() })

	var resp PreToolResponse
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &resp))
	assert.Contains(t, resp.SystemMessage, "cert.pem")
}

func TestGuardWrite_AllowsNonMatchingFile(t *testing.T) {
	setupGuardEnv(t, `{"write":{"blockedPatterns":["*.env"]}}`)
	fakeStdin(t, `{"tool_input":{"file_path":"/project/main.go"}}`)
	out := captureStdout(t, func() { _ = guardWrite() })
	assert.Empty(t, out)
}

// --- guard-bash tests ---

func TestGuardBash_AllowsWhenCommandMissing(t *testing.T) {
	setupGuardEnv(t, `{"bash":{"blockedCommands":["rm -rf /"],"blockedPatterns":[]}}`)
	fakeStdin(t, `{}`)
	out := captureStdout(t, func() { _ = guardBash() })
	assert.Empty(t, out)
}

func TestGuardBash_AllowsWhenGuardJSONAbsent(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("CLAUDE_PROJECT_DIR", dir)
	fakeStdin(t, `{"tool_input":{"command":"rm -rf /"}}`)
	out := captureStdout(t, func() { _ = guardBash() })
	assert.Empty(t, out)
}

func TestGuardBash_BlocksExactMatchCommand(t *testing.T) {
	setupGuardEnv(t, `{"bash":{"blockedCommands":["rm -rf /"],"blockedPatterns":[]}}`)
	fakeStdin(t, `{"tool_input":{"command":"rm -rf /"}}`)
	out := captureStdout(t, func() { _ = guardBash() })

	var resp PreToolResponse
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &resp))
	assert.Equal(t, "deny", resp.HookSpecificOutput.PermissionDecision)
	assert.Contains(t, resp.SystemMessage, "rm -rf /")
}

func TestGuardBash_AllowsNonExactMatchCommand(t *testing.T) {
	setupGuardEnv(t, `{"bash":{"blockedCommands":["rm -rf /"],"blockedPatterns":[]}}`)
	fakeStdin(t, `{"tool_input":{"command":"ls -la"}}`)
	out := captureStdout(t, func() { _ = guardBash() })
	assert.Empty(t, out)
}

func TestGuardBash_BlocksPatternMatch(t *testing.T) {
	setupGuardEnv(t, `{"bash":{"blockedCommands":[],"blockedPatterns":["git push --force"]}}`)
	fakeStdin(t, `{"tool_input":{"command":"git push --force origin main"}}`)
	out := captureStdout(t, func() { _ = guardBash() })

	var resp PreToolResponse
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &resp))
	assert.Contains(t, resp.SystemMessage, "git push --force")
}

func TestGuardBash_AllowsCommandNotMatchingPattern(t *testing.T) {
	setupGuardEnv(t, `{"bash":{"blockedCommands":[],"blockedPatterns":["rm -rf"]}}`)
	fakeStdin(t, `{"tool_input":{"command":"git status"}}`)
	out := captureStdout(t, func() { _ = guardBash() })
	assert.Empty(t, out)
}

func TestGuardBash_BlocksMultiplePatterns(t *testing.T) {
	setupGuardEnv(t, `{"bash":{"blockedCommands":[],"blockedPatterns":["rm -rf","git push --force","drop table"]}}`)
	fakeStdin(t, `{"tool_input":{"command":"sudo rm -rf /var/log"}}`)
	out := captureStdout(t, func() { _ = guardBash() })

	var resp PreToolResponse
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &resp))
	assert.Contains(t, resp.SystemMessage, "rm -rf")
}

func TestGuardBash_AllowsSafeCommandWithStrictConfig(t *testing.T) {
	setupGuardEnv(t, `{"bash":{"blockedCommands":["rm -rf /"],"blockedPatterns":["git push --force","drop table","truncate table"]}}`)
	fakeStdin(t, `{"tool_input":{"command":"go test ./..."}}`)
	out := captureStdout(t, func() { _ = guardBash() })
	assert.Empty(t, out)
}

func TestGuardBash_HandlesEmptyToolInput(t *testing.T) {
	setupGuardEnv(t, `{"bash":{"blockedCommands":["rm -rf /"],"blockedPatterns":[]}}`)
	fakeStdin(t, `{"tool_input":{}}`)
	out := captureStdout(t, func() { _ = guardBash() })
	assert.Empty(t, out)
}

func TestGuardBash_HandlesMalformedJSON(t *testing.T) {
	setupGuardEnv(t, `{"bash":{"blockedCommands":["rm -rf /"],"blockedPatterns":[]}}`)
	fakeStdin(t, `not valid json`)
	out := captureStdout(t, func() { _ = guardBash() })
	assert.Empty(t, out)
}
