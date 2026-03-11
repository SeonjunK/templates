// Package hook provides core types and parsing for Claude Code hook payloads.
package hook

import (
	"encoding/json"
	"io"
)

// Input represents the JSON payload from Claude Code hook stdin.
type Input struct {
	ToolInput struct {
		FilePath string `json:"file_path"`
		Command  string `json:"command"`
	} `json:"tool_input"`
}

// ParseInput reads and parses the JSON hook input from the given reader.
func ParseInput(r io.Reader) (*Input, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var input Input
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}

	return &input, nil
}
