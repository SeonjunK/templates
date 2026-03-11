package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// captureStdout captures stdout output during fn execution.
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	old := os.Stdout
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	return buf.String()
}

func TestDeny(t *testing.T) {
	out := captureStdout(t, func() {
		deny("⚠ File access blocked: .env (matched pattern: .env)")
	})

	var resp PreToolResponse
	require.NoError(t, json.Unmarshal([]byte(out), &resp))
	assert.Equal(t, "deny", resp.HookSpecificOutput.PermissionDecision)
	assert.Contains(t, resp.SystemMessage, "File access blocked")
}

func TestWarn(t *testing.T) {
	out := captureStdout(t, func() {
		warn("⚠ Format failed for main.go")
	})

	var resp PreToolResponse
	require.NoError(t, json.Unmarshal([]byte(out), &resp))
	assert.Nil(t, resp.HookSpecificOutput)
	assert.Contains(t, resp.SystemMessage, "Format failed")
}

func TestBlock(t *testing.T) {
	out := captureStdout(t, func() {
		block("Tests failed", "⚠ Tests failed. Run `go test` to see details.")
	})

	var resp StopResponse
	require.NoError(t, json.Unmarshal([]byte(out), &resp))
	assert.Equal(t, "block", resp.Decision)
	assert.Equal(t, "Tests failed", resp.Reason)
	assert.Contains(t, resp.SystemMessage, "Tests failed")
}

func TestApprove(t *testing.T) {
	out := captureStdout(t, func() {
		approve("✓ All checks passed")
	})

	var resp StopResponse
	require.NoError(t, json.Unmarshal([]byte(out), &resp))
	assert.Equal(t, "approve", resp.Decision)
	assert.Contains(t, resp.SystemMessage, "All checks passed")
}
