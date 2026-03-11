// Package response provides types and builders for Claude Code hook responses.
package response

import (
	"encoding/json"
	"fmt"
	"io"
)

// PreToolResponse is emitted by PreToolUse hooks (guard-read, guard-write, guard-bash).
type PreToolResponse struct {
	HookSpecificOutput *HookSpecificOutput `json:"hookSpecificOutput,omitempty"`
}

// HookSpecificOutput contains the permission decision for PreToolUse hooks.
type HookSpecificOutput struct {
	PermissionDecision string `json:"permissionDecision"`
}

// StopResponse is emitted by Stop hooks (verify).
type StopResponse struct {
	Decision string `json:"decision,omitempty"`
	Reason   string `json:"reason,omitempty"`
}

// Deny writes a PreToolResponse that blocks the tool call with the given reason.
func Deny(w io.Writer, reason string) error {
	resp := PreToolResponse{
		HookSpecificOutput: &HookSpecificOutput{
			PermissionDecision: fmt.Sprintf("deny: %s", reason),
		},
	}

	return json.NewEncoder(w).Encode(resp)
}

// Warn writes a PreToolResponse that allows the tool call but shows a warning.
func Warn(w io.Writer, reason string) error {
	resp := PreToolResponse{
		HookSpecificOutput: &HookSpecificOutput{
			PermissionDecision: fmt.Sprintf("warn: %s", reason),
		},
	}

	return json.NewEncoder(w).Encode(resp)
}

// Block writes a StopResponse that prevents the agent from stopping.
func Block(w io.Writer, reason string) error {
	resp := StopResponse{
		Decision: "block",
		Reason:   reason,
	}

	return json.NewEncoder(w).Encode(resp)
}

// Approve writes a StopResponse that allows the agent to stop.
func Approve(w io.Writer) error {
	resp := StopResponse{
		Decision: "approve",
	}

	return json.NewEncoder(w).Encode(resp)
}
