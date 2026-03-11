package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// formatCmd runs configured format commands on a written/edited file.
// Format rules are read from .claude/hooks.json.
func formatCmd() error {
	input, err := parseInput(os.Stdin)
	if err != nil || input.ToolInput.FilePath == "" {
		return nil
	}

	dir := projectDir()
	config, err := loadHooksConfig(dir)
	if err != nil || config == nil {
		return nil
	}

	file := input.ToolInput.FilePath
	ext := filepath.Ext(file)

	for _, rule := range config.Format {
		if !matchesExtension(ext, rule.Extensions) {
			continue
		}
		for _, cmdTemplate := range rule.Commands {
			args := expandFileArg(cmdTemplate, file)
			if len(args) == 0 {
				continue
			}
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Dir = dir
			if err := cmd.Run(); err != nil {
				warn(fmt.Sprintf("⚠ Format failed for %s: %s", file, strings.Join(args, " ")))
				return nil
			}
		}
		return nil
	}

	return nil
}

// matchesExtension checks if ext is in the list of extensions.
func matchesExtension(ext string, extensions []string) bool {
	for _, e := range extensions {
		if e == ext {
			return true
		}
	}
	return false
}

// expandFileArg replaces "{{file}}" placeholders in a command template with the actual file path.
func expandFileArg(cmdTemplate []string, file string) []string {
	result := make([]string, len(cmdTemplate))
	for i, part := range cmdTemplate {
		result[i] = strings.ReplaceAll(part, "{{file}}", file)
	}
	return result
}
