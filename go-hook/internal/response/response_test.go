package response_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/example/go-hook/internal/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeny(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, response.Deny(&buf, "blocked file"))

	var resp response.PreToolResponse
	require.NoError(t, json.Unmarshal(buf.Bytes(), &resp))
	assert.Equal(t, "deny: blocked file", resp.HookSpecificOutput.PermissionDecision)
}

func TestWarn(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, response.Warn(&buf, "sensitive path"))

	var resp response.PreToolResponse
	require.NoError(t, json.Unmarshal(buf.Bytes(), &resp))
	assert.Equal(t, "warn: sensitive path", resp.HookSpecificOutput.PermissionDecision)
}

func TestBlock(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, response.Block(&buf, "tests failing"))

	var resp response.StopResponse
	require.NoError(t, json.Unmarshal(buf.Bytes(), &resp))
	assert.Equal(t, "block", resp.Decision)
	assert.Equal(t, "tests failing", resp.Reason)
}

func TestApprove(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, response.Approve(&buf))

	var resp response.StopResponse
	require.NoError(t, json.Unmarshal(buf.Bytes(), &resp))
	assert.Equal(t, "approve", resp.Decision)
	assert.Empty(t, resp.Reason)
}
