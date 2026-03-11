package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// guardRead checks if a file read should be blocked by guard policy.
func guardRead() error {
	input, err := parseInput(os.Stdin)
	if err != nil || input.ToolInput.FilePath == "" {
		return nil
	}

	config, err := loadGuardConfig(projectDir())
	if err != nil || config == nil {
		return nil
	}

	basename := filepath.Base(input.ToolInput.FilePath)
	for _, pattern := range config.Read.BlockedPatterns {
		if matched, _ := filepath.Match(pattern, basename); matched {
			deny(fmt.Sprintf("⚠ File access blocked: %s (matched pattern: %s)", basename, pattern))
			return nil
		}
	}
	return nil
}

// guardWrite checks if a file write should be blocked by guard policy.
func guardWrite() error {
	input, err := parseInput(os.Stdin)
	if err != nil || input.ToolInput.FilePath == "" {
		return nil
	}

	config, err := loadGuardConfig(projectDir())
	if err != nil || config == nil {
		return nil
	}

	basename := filepath.Base(input.ToolInput.FilePath)
	for _, pattern := range config.Write.BlockedPatterns {
		if matched, _ := filepath.Match(pattern, basename); matched {
			deny(fmt.Sprintf("⚠ File write blocked: %s (matched pattern: %s)", basename, pattern))
			return nil
		}
	}
	return nil
}

// guardBash checks if a bash command should be blocked by guard policy.
func guardBash() error {
	input, err := parseInput(os.Stdin)
	if err != nil || input.ToolInput.Command == "" {
		return nil
	}

	config, err := loadGuardConfig(projectDir())
	if err != nil || config == nil {
		return nil
	}

	cmd := input.ToolInput.Command

	for _, blocked := range config.Bash.BlockedCommands {
		if cmd == blocked {
			deny(fmt.Sprintf("⚠ Command blocked: %s is blocked by guard policy", cmd))
			return nil
		}
	}

	for _, pattern := range config.Bash.BlockedPatterns {
		if strings.Contains(cmd, pattern) {
			deny(fmt.Sprintf("⚠ Command blocked: matches pattern %s", pattern))
			return nil
		}
	}
	return nil
}
