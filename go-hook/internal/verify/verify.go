// Package verify implements the Stop verify hook that runs sequential verification steps.
package verify

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/example/go-hook/internal/hook"
	"github.com/example/go-hook/internal/response"
)

// Run loads the hooks config and executes each verify step sequentially.
// If any step fails, it blocks the agent from stopping and reports which step failed.
func Run() error {
	config, err := hook.LoadHooksConfig(hook.ProjectDir())
	if err != nil {
		return fmt.Errorf("loading hooks config: %w", err)
	}

	if config == nil {
		return nil
	}

	var failures []string

	for _, step := range config.Verify {
		if err := runStep(step); err != nil {
			failures = append(failures, fmt.Sprintf("%s: %v", step.Name, err))
		}
	}

	if len(failures) > 0 {
		reason := fmt.Sprintf("verification failed:\n%s", strings.Join(failures, "\n"))

		return response.Block(os.Stdout, reason)
	}

	return response.Approve(os.Stdout)
}

func runStep(step hook.VerifyStep) error {
	if len(step.Command) == 0 {
		return nil
	}

	cmd := exec.Command(step.Command[0], step.Command[1:]...) //nolint:gosec // commands come from trusted hooks.json config
	cmd.Dir = hook.ProjectDir()
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if step.Fix != "" {
			return fmt.Errorf("%w (fix: %s)", err, step.Fix)
		}

		return err
	}

	return nil
}
