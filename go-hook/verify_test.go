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

// setupHooksEnv creates a temp dir with hooks.json and sets CLAUDE_PROJECT_DIR.
func setupHooksEnv(t *testing.T, hooksJSON string) string {
	t.Helper()
	dir := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(dir, ".claude"), 0o755))
	if hooksJSON != "" {
		require.NoError(t, os.WriteFile(
			filepath.Join(dir, ".claude", "hooks.json"),
			[]byte(hooksJSON),
			0o644,
		))
	}
	t.Setenv("CLAUDE_PROJECT_DIR", dir)
	return dir
}

func TestVerifyCmd_NoConfigReturnsNil(t *testing.T) {
	t.Setenv("CLAUDE_PROJECT_DIR", "")
	out := captureStdout(t, func() { _ = verifyCmd() })
	assert.Empty(t, out)
}

func TestVerifyCmd_EmptyStepsReturnsNil(t *testing.T) {
	setupHooksEnv(t, `{"verify":[]}`)
	out := captureStdout(t, func() { _ = verifyCmd() })
	assert.Empty(t, out)
}

func TestVerifyCmd_AllPassApproves(t *testing.T) {
	dir := setupHooksEnv(t, `{"verify":[
		{"name":"check","command":["true"],"fix":"true"}
	]}`)
	_ = dir
	out := captureStdout(t, func() { _ = verifyCmd() })

	var resp StopResponse
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &resp))
	assert.Equal(t, "approve", resp.Decision)
	assert.Contains(t, resp.SystemMessage, "check")
}

func TestVerifyCmd_FailureBlocks(t *testing.T) {
	setupHooksEnv(t, `{"verify":[
		{"name":"failing step","command":["false"],"fix":"fix it"}
	]}`)
	out := captureStdout(t, func() { _ = verifyCmd() })

	var resp StopResponse
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &resp))
	assert.Equal(t, "block", resp.Decision)
	assert.Equal(t, "failing step", resp.Reason)
	assert.Contains(t, resp.SystemMessage, "fix it")
}

func TestVerifyCmd_StopsAtFirstFailure(t *testing.T) {
	setupHooksEnv(t, `{"verify":[
		{"name":"pass","command":["true"],"fix":"true"},
		{"name":"fail","command":["false"],"fix":"fix"},
		{"name":"skip","command":["true"],"fix":"true"}
	]}`)
	out := captureStdout(t, func() { _ = verifyCmd() })

	var resp StopResponse
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &resp))
	assert.Equal(t, "block", resp.Decision)
	assert.Equal(t, "fail", resp.Reason)
}

func TestVerifyCmd_MultiplePassShowsAllNames(t *testing.T) {
	setupHooksEnv(t, `{"verify":[
		{"name":"fmt","command":["true"]},
		{"name":"lint","command":["true"]},
		{"name":"test","command":["true"]}
	]}`)
	out := captureStdout(t, func() { _ = verifyCmd() })

	var resp StopResponse
	require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(out)), &resp))
	assert.Equal(t, "approve", resp.Decision)
	assert.Contains(t, resp.SystemMessage, "fmt, lint, test")
}
