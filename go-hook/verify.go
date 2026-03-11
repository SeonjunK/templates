package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// verifyCmd runs all configured verification steps before session stop.
// Verify steps are read from .claude/hooks.json.
func verifyCmd() error {
	dir := projectDir()
	config, err := loadHooksConfig(dir)
	if err != nil || config == nil {
		return nil
	}

	if len(config.Verify) == 0 {
		return nil
	}

	var passed []string
	for _, step := range config.Verify {
		if len(step.Command) == 0 {
			continue
		}
		cmd := exec.Command(step.Command[0], step.Command[1:]...)
		cmd.Dir = dir
		if err := cmd.Run(); err != nil {
			fix := step.Fix
			if fix == "" {
				fix = strings.Join(step.Command, " ")
			}
			block(step.Name, fmt.Sprintf("⚠ %s. Run `%s` to see details.", step.Name, fix))
			return nil
		}
		passed = append(passed, step.Name)
	}

	approve(fmt.Sprintf("✓ All checks passed (%s)", strings.Join(passed, ", ")))
	return nil
}
