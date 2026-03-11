package verify

import (
	"testing"

	"github.com/example/go-hook/internal/hook"
	"github.com/stretchr/testify/assert"
)

func TestRunStep(t *testing.T) {
	t.Run("succeeds with true command", func(t *testing.T) {
		step := hook.VerifyStep{
			Name:    "always-pass",
			Command: []string{"true"},
		}
		assert.NoError(t, runStep(step))
	})

	t.Run("fails with false command", func(t *testing.T) {
		step := hook.VerifyStep{
			Name:    "always-fail",
			Command: []string{"false"},
		}
		assert.Error(t, runStep(step))
	})

	t.Run("includes fix hint on failure", func(t *testing.T) {
		step := hook.VerifyStep{
			Name:    "fixable",
			Command: []string{"false"},
			Fix:     "run fix command",
		}
		err := runStep(step)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "fix: run fix command")
	})

	t.Run("empty command is no-op", func(t *testing.T) {
		step := hook.VerifyStep{
			Name:    "empty",
			Command: nil,
		}
		assert.NoError(t, runStep(step))
	})
}
