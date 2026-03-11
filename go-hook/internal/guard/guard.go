// Package guard implements PreToolUse guard hooks that block disallowed operations.
package guard

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/example/go-hook/internal/hook"
	"github.com/example/go-hook/internal/response"
)

// Read guards the Read tool against blocked file patterns.
func Read() error {
	input, err := hook.ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("parsing input: %w", err)
	}

	config, err := hook.LoadGuardConfig(hook.ProjectDir())
	if err != nil {
		return fmt.Errorf("loading guard config: %w", err)
	}

	if config == nil {
		return nil
	}

	if reason := matchFilePattern(input.ToolInput.FilePath, config.Read.BlockedPatterns); reason != "" {
		return response.Deny(os.Stdout, reason)
	}

	return nil
}

// Write guards the Write/Edit tools against blocked file patterns.
func Write() error {
	input, err := hook.ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("parsing input: %w", err)
	}

	config, err := hook.LoadGuardConfig(hook.ProjectDir())
	if err != nil {
		return fmt.Errorf("loading guard config: %w", err)
	}

	if config == nil {
		return nil
	}

	if reason := matchFilePattern(input.ToolInput.FilePath, config.Write.BlockedPatterns); reason != "" {
		return response.Deny(os.Stdout, reason)
	}

	return nil
}

// Bash guards the Bash tool against blocked commands and patterns.
func Bash() error {
	input, err := hook.ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("parsing input: %w", err)
	}

	config, err := hook.LoadGuardConfig(hook.ProjectDir())
	if err != nil {
		return fmt.Errorf("loading guard config: %w", err)
	}

	if config == nil {
		return nil
	}

	cmd := input.ToolInput.Command

	for _, blocked := range config.Bash.BlockedCommands {
		if cmd == blocked {
			return response.Deny(os.Stdout, fmt.Sprintf("blocked command: %s", blocked))
		}
	}

	for _, pattern := range config.Bash.BlockedPatterns {
		if strings.Contains(cmd, pattern) {
			return response.Deny(os.Stdout, fmt.Sprintf("command matches blocked pattern: %s", pattern))
		}
	}

	return nil
}

// matchFilePattern checks if a file path matches any blocked pattern.
// Returns the reason string if blocked, empty string if allowed.
func matchFilePattern(filePath string, patterns []string) string {
	if filePath == "" {
		return ""
	}

	base := filepath.Base(filePath)

	for _, pattern := range patterns {
		matched, err := filepath.Match(pattern, base)
		if err != nil {
			continue
		}

		if matched {
			return fmt.Sprintf("file matches blocked pattern: %s", pattern)
		}
	}

	return ""
}
