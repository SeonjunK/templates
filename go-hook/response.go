package main

import (
	"encoding/json"
	"fmt"
)

// PreToolResponse is the JSON output for PreToolUse hooks.
type PreToolResponse struct {
	HookSpecificOutput *HookOutput `json:"hookSpecificOutput,omitempty"`
	SystemMessage      string      `json:"systemMessage,omitempty"`
}

// HookOutput contains the permission decision for PreToolUse hooks.
type HookOutput struct {
	PermissionDecision string `json:"permissionDecision"`
}

// StopResponse is the JSON output for Stop hooks.
type StopResponse struct {
	Decision      string `json:"decision"`
	Reason        string `json:"reason,omitempty"`
	SystemMessage string `json:"systemMessage"`
}

// deny outputs a PreToolUse deny response and prints it to stdout.
func deny(message string) {
	resp := PreToolResponse{
		HookSpecificOutput: &HookOutput{PermissionDecision: "deny"},
		SystemMessage:      message,
	}
	printJSON(resp)
}

// warn outputs a PreToolUse warning (no deny) and prints it to stdout.
func warn(message string) {
	resp := PreToolResponse{
		SystemMessage: message,
	}
	printJSON(resp)
}

// block outputs a Stop block response and prints it to stdout.
func block(reason, message string) {
	resp := StopResponse{
		Decision:      "block",
		Reason:        reason,
		SystemMessage: message,
	}
	printJSON(resp)
}

// approve outputs a Stop approve response and prints it to stdout.
func approve(message string) {
	resp := StopResponse{
		Decision:      "approve",
		SystemMessage: message,
	}
	printJSON(resp)
}

func printJSON(v any) {
	data, _ := json.Marshal(v)
	fmt.Println(string(data))
}
