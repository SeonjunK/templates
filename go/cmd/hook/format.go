package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// format runs gofmt and golines on a written/edited Go file.
func format() error {
	input, err := parseInput(os.Stdin)
	if err != nil || input.ToolInput.FilePath == "" {
		return nil
	}

	if filepath.Ext(input.ToolInput.FilePath) != ".go" {
		return nil
	}

	file := input.ToolInput.FilePath

	if err := exec.Command("gofmt", "-w", file).Run(); err != nil {
		warn(fmt.Sprintf("⚠ Format failed for %s", file))
		return nil
	}

	if err := exec.Command("go", "tool", "golines", "-w", file).Run(); err != nil {
		warn(fmt.Sprintf("⚠ golines failed for %s", file))
		return nil
	}

	return nil
}
