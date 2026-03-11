package hook

import (
	"encoding/json"
	"os"
	"path/filepath"
)

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

// ProjectDir returns CLAUDE_PROJECT_DIR or empty string.
func ProjectDir() string {
	return os.Getenv("CLAUDE_PROJECT_DIR")
}

// LoadGuardConfig loads guard.json from the project directory.
// Returns nil config (no error) if the project dir is unset or the file is absent.
func LoadGuardConfig(dir string) (*GuardConfig, error) {
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

// LoadHooksConfig loads hooks.json from the project directory.
// Returns nil config (no error) if the project dir is unset or the file is absent.
func LoadHooksConfig(dir string) (*HooksConfig, error) {
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
