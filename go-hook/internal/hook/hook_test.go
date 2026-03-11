package hook_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/example/go-hook/internal/hook"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseInput(t *testing.T) {
	t.Run("parses file_path", func(t *testing.T) {
		input, err := hook.ParseInput(strings.NewReader(`{"tool_input":{"file_path":"/project/.env"}}`))
		require.NoError(t, err)
		assert.Equal(t, "/project/.env", input.ToolInput.FilePath)
	})

	t.Run("parses command", func(t *testing.T) {
		input, err := hook.ParseInput(strings.NewReader(`{"tool_input":{"command":"rm -rf /"}}`))
		require.NoError(t, err)
		assert.Equal(t, "rm -rf /", input.ToolInput.Command)
	})

	t.Run("empty tool_input", func(t *testing.T) {
		input, err := hook.ParseInput(strings.NewReader(`{"tool_input":{}}`))
		require.NoError(t, err)
		assert.Empty(t, input.ToolInput.FilePath)
		assert.Empty(t, input.ToolInput.Command)
	})

	t.Run("malformed JSON returns error", func(t *testing.T) {
		_, err := hook.ParseInput(strings.NewReader(`not valid json`))
		assert.Error(t, err)
	})

	t.Run("empty object", func(t *testing.T) {
		input, err := hook.ParseInput(strings.NewReader(`{}`))
		require.NoError(t, err)
		assert.Empty(t, input.ToolInput.FilePath)
	})
}

func TestLoadGuardConfig(t *testing.T) {
	t.Run("returns nil when dir is empty", func(t *testing.T) {
		config, err := hook.LoadGuardConfig("")
		assert.NoError(t, err)
		assert.Nil(t, config)
	})

	t.Run("returns nil when guard.json is absent", func(t *testing.T) {
		config, err := hook.LoadGuardConfig(t.TempDir())
		assert.NoError(t, err)
		assert.Nil(t, config)
	})

	t.Run("loads valid guard.json", func(t *testing.T) {
		dir := t.TempDir()
		require.NoError(t, os.MkdirAll(filepath.Join(dir, ".claude"), 0o755))
		require.NoError(t, os.WriteFile(
			filepath.Join(dir, ".claude", "guard.json"),
			[]byte(`{"read":{"blockedPatterns":[".env*","*.pem"]},"bash":{"blockedCommands":["rm -rf /"],"blockedPatterns":["git push --force"]}}`),
			0o644,
		))

		config, err := hook.LoadGuardConfig(dir)
		require.NoError(t, err)
		require.NotNil(t, config)
		assert.Equal(t, []string{".env*", "*.pem"}, config.Read.BlockedPatterns)
		assert.Equal(t, []string{"rm -rf /"}, config.Bash.BlockedCommands)
	})
}

func TestLoadHooksConfig(t *testing.T) {
	t.Run("returns nil when dir is empty", func(t *testing.T) {
		config, err := hook.LoadHooksConfig("")
		assert.NoError(t, err)
		assert.Nil(t, config)
	})

	t.Run("loads valid hooks.json", func(t *testing.T) {
		dir := t.TempDir()
		require.NoError(t, os.MkdirAll(filepath.Join(dir, ".claude"), 0o755))
		require.NoError(t, os.WriteFile(
			filepath.Join(dir, ".claude", "hooks.json"),
			[]byte(`{"format":[{"extensions":[".go"],"commands":[["gofmt","-w","{{file}}"]]}],"verify":[{"name":"test","command":["go","test","./..."],"fix":"go test ./..."}]}`),
			0o644,
		))

		config, err := hook.LoadHooksConfig(dir)
		require.NoError(t, err)
		require.NotNil(t, config)
		assert.Len(t, config.Format, 1)
		assert.Equal(t, []string{".go"}, config.Format[0].Extensions)
		assert.Len(t, config.Verify, 1)
		assert.Equal(t, "test", config.Verify[0].Name)
	})
}
