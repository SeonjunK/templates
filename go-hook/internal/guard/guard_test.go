package guard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchFilePattern(t *testing.T) {
	patterns := []string{".env*", "*.pem", "*.key"}

	t.Run("blocks .env", func(t *testing.T) {
		assert.NotEmpty(t, matchFilePattern("/project/.env", patterns))
	})

	t.Run("blocks .env.local", func(t *testing.T) {
		assert.NotEmpty(t, matchFilePattern("/project/.env.local", patterns))
	})

	t.Run("blocks .pem files", func(t *testing.T) {
		assert.NotEmpty(t, matchFilePattern("/project/certs/server.pem", patterns))
	})

	t.Run("blocks .key files", func(t *testing.T) {
		assert.NotEmpty(t, matchFilePattern("/project/private.key", patterns))
	})

	t.Run("allows regular files", func(t *testing.T) {
		assert.Empty(t, matchFilePattern("/project/main.go", patterns))
	})

	t.Run("allows empty file path", func(t *testing.T) {
		assert.Empty(t, matchFilePattern("", patterns))
	})

	t.Run("allows empty patterns", func(t *testing.T) {
		assert.Empty(t, matchFilePattern("/project/.env", nil))
	})

	t.Run("skips invalid pattern", func(t *testing.T) {
		assert.Empty(t, matchFilePattern("/project/file.go", []string{"["}))
	})
}
