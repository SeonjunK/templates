package main

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

// HookInput represents the JSON payload from Claude Code hook stdin.
type HookInput struct {
	ToolInput struct {
		FilePath string `json:"file_path"`
		Command  string `json:"command"`
	} `json:"tool_input"`
}

// GuardConfig represents the .claude/guard.json configuration.
type GuardConfig struct {
	Read struct {
		BlockedPatterns []string `json:"blockedPatterns"`
	} `json:"read"`
	Write struct {
		BlockedPatterns []string `json:"blockedPatterns"`
	} `json:"write"`
	Bash struct {
		BlockedCommands []string `json:"blockedCommands"`
		BlockedPatterns []string `json:"blockedPatterns"`
	} `json:"bash"`
}

// FormatRule defines format commands for a set of file extensions.
type FormatRule struct {
	Extensions []string   `json:"extensions"`
	Commands   [][]string `json:"commands"`
}

// VerifyStep defines a single verification step.
type VerifyStep struct {
	Name    string   `json:"name"`
	Command []string `json:"command"`
	Fix     string   `json:"fix"`
}

// HooksConfig represents the .claude/hooks.json configuration.
type HooksConfig struct {
	Format []FormatRule `json:"format"`
	Verify []VerifyStep `json:"verify"`
}

// parseInput reads and parses the JSON hook input from the given reader.
func parseInput(r io.Reader) (*HookInput, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var input HookInput
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	return &input, nil
}

// projectDir returns CLAUDE_PROJECT_DIR or empty string.
func projectDir() string {
	return os.Getenv("CLAUDE_PROJECT_DIR")
}

// loadGuardConfig loads guard.json from the project directory.
// Returns nil config (no error) if the project dir is unset or the file is absent.
func loadGuardConfig(dir string) (*GuardConfig, error) {
	if dir == "" {
		return nil, nil
	}
	path := filepath.Join(dir, ".claude", "guard.json")
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var config GuardConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// loadHooksConfig loads hooks.json from the project directory.
// Returns nil config (no error) if the project dir is unset or the file is absent.
func loadHooksConfig(dir string) (*HooksConfig, error) {
	if dir == "" {
		return nil, nil
	}
	path := filepath.Join(dir, ".claude", "hooks.json")
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var config HooksConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
