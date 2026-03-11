package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchesExtension(t *testing.T) {
	extensions := []string{".go", ".mod"}

	t.Run("matches .go", func(t *testing.T) {
		assert.True(t, matchesExtension(".go", extensions))
	})

	t.Run("matches .mod", func(t *testing.T) {
		assert.True(t, matchesExtension(".mod", extensions))
	})

	t.Run("does not match .py", func(t *testing.T) {
		assert.False(t, matchesExtension(".py", extensions))
	})

	t.Run("does not match empty", func(t *testing.T) {
		assert.False(t, matchesExtension("", extensions))
	})

	t.Run("empty extensions list", func(t *testing.T) {
		assert.False(t, matchesExtension(".go", nil))
	})
}
