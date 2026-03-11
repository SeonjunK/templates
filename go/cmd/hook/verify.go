package main

import (
	"os"
	"os/exec"
	"strings"
)

// verify runs all quality checks (format, golines, lint, test) before session stop.
func verify() error {
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		return nil
	}

	unformatted, _ := exec.Command("gofmt", "-l", ".").Output()
	// Filter out vendor/ lines
	var files []string
	for _, line := range strings.Split(strings.TrimSpace(string(unformatted)), "\n") {
		if line != "" && !strings.HasPrefix(line, "vendor/") {
			files = append(files, line)
		}
	}
	if len(files) > 0 {
		block("Unformatted files", "⚠ Format failed - unformatted files. Run `gofmt -w .` to fix.")
		return nil
	}

	longLines, _ := exec.Command("go", "tool", "golines", "-l", ".").Output()
	var longFiles []string
	for _, line := range strings.Split(strings.TrimSpace(string(longLines)), "\n") {
		if line != "" && !strings.HasPrefix(line, "vendor/") {
			longFiles = append(longFiles, line)
		}
	}
	if len(longFiles) > 0 {
		block("Long lines detected", "⚠ Long lines detected. Run `go tool golines -w .` to fix.")
		return nil
	}

	if err := exec.Command("go", "tool", "golangci-lint", "run", "./...").Run(); err != nil {
		block("Lint failed", "⚠ Lint failed. Run `go tool golangci-lint run ./...` to see details.")
		return nil
	}

	if err := exec.Command("go", "test", "-race", "./...").Run(); err != nil {
		block("Tests failed", "⚠ Tests failed. Run `go test -race ./...` to see details.")
		return nil
	}

	approve("✓ All checks passed (format, golines, lint, test)")
	return nil
}
