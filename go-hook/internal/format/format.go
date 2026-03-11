// Package format implements the PostToolUse format hook that runs formatters on changed files.
package format

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/example/go-hook/internal/hook"
)

// Run reads the hook input, matches the file extension against configured format rules,
// and executes matching format commands with {{file}} expanded to the actual file path.
func Run() error {
	input, err := hook.ParseInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("parsing input: %w", err)
	}

	filePath := input.ToolInput.FilePath
	if filePath == "" {
		return nil
	}

	config, err := hook.LoadHooksConfig(hook.ProjectDir())
	if err != nil {
		return fmt.Errorf("loading hooks config: %w", err)
	}

	if config == nil {
		return nil
	}

	ext := filepath.Ext(filePath)

	for _, rule := range config.Format {
		if !matchesExtension(ext, rule.Extensions) {
			continue
		}

		for _, cmdTemplate := range rule.Commands {
			if err := runCommand(cmdTemplate, filePath); err != nil {
				return err
			}
		}
	}

	return nil
}

func matchesExtension(ext string, extensions []string) bool {
	for _, e := range extensions {
		if e == ext {
			return true
		}
	}

	return false
}

func runCommand(cmdTemplate []string, filePath string) error {
	if len(cmdTemplate) == 0 {
		return nil
	}

	args := make([]string, len(cmdTemplate))
	for i, arg := range cmdTemplate {
		args[i] = strings.ReplaceAll(arg, "{{file}}", filePath)
	}

	cmd := exec.Command(args[0], args[1:]...) //nolint:gosec // commands come from trusted hooks.json config
	cmd.Dir = hook.ProjectDir()
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("running %s: %w", args[0], err)
	}

	return nil
}
