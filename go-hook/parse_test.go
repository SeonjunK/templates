package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseInput(t *testing.T) {
	t.Run("parses file_path", func(t *testing.T) {
		r := strings.NewReader(`{"tool_input":{"file_path":"/project/.env"}}`)
		input, err := parseInput(r)
		require.NoError(t, err)
		assert.Equal(t, "/project/.env", input.ToolInput.FilePath)
	})

	t.Run("parses command", func(t *testing.T) {
		r := strings.NewReader(`{"tool_input":{"command":"rm -rf /"}}`)
		input, err := parseInput(r)
		require.NoError(t, err)
		assert.Equal(t, "rm -rf /", input.ToolInput.Command)
	})

	t.Run("empty tool_input", func(t *testing.T) {
		r := strings.NewReader(`{"tool_input":{}}`)
		input, err := parseInput(r)
		require.NoError(t, err)
		assert.Empty(t, input.ToolInput.FilePath)
		assert.Empty(t, input.ToolInput.Command)
	})

	t.Run("malformed JSON returns error", func(t *testing.T) {
		r := strings.NewReader(`not valid json`)
		_, err := parseInput(r)
		assert.Error(t, err)
	})

	t.Run("empty object", func(t *testing.T) {
		r := strings.NewReader(`{}`)
		input, err := parseInput(r)
		require.NoError(t, err)
		assert.Empty(t, input.ToolInput.FilePath)
	})
}

func TestLoadGuardConfig(t *testing.T) {
	t.Run("returns nil when projectDir is empty", func(t *testing.T) {
		config, err := loadGuardConfig("")
		assert.NoError(t, err)
		assert.Nil(t, config)
	})

	t.Run("returns nil when guard.json is absent", func(t *testing.T) {
		dir := t.TempDir()
		config, err := loadGuardConfig(dir)
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

		config, err := loadGuardConfig(dir)
		require.NoError(t, err)
		require.NotNil(t, config)
		assert.Equal(t, []string{".env*", "*.pem"}, config.Read.BlockedPatterns)
		assert.Equal(t, []string{"rm -rf /"}, config.Bash.BlockedCommands)
		assert.Equal(t, []string{"git push --force"}, config.Bash.BlockedPatterns)
	})

	t.Run("empty guard.json object", func(t *testing.T) {
		dir := t.TempDir()
		require.NoError(t, os.MkdirAll(filepath.Join(dir, ".claude"), 0o755))
		require.NoError(t, os.WriteFile(
			filepath.Join(dir, ".claude", "guard.json"),
			[]byte(`{}`),
			0o644,
		))

		config, err := loadGuardConfig(dir)
		require.NoError(t, err)
		require.NotNil(t, config)
		assert.Empty(t, config.Read.BlockedPatterns)
	})
}

func TestLoadHooksConfig(t *testing.T) {
	t.Run("returns nil when projectDir is empty", func(t *testing.T) {
		config, err := loadHooksConfig("")
		assert.NoError(t, err)
		assert.Nil(t, config)
	})

	t.Run("returns nil when hooks.json is absent", func(t *testing.T) {
		dir := t.TempDir()
		config, err := loadHooksConfig(dir)
		assert.NoError(t, err)
		assert.Nil(t, config)
	})

	t.Run("loads valid hooks.json", func(t *testing.T) {
		dir := t.TempDir()
		require.NoError(t, os.MkdirAll(filepath.Join(dir, ".claude"), 0o755))
		require.NoError(t, os.WriteFile(
			filepath.Join(dir, ".claude", "hooks.json"),
			[]byte(`{
				"format": [{"extensions": [".go"], "commands": [["gofmt", "-w", "{{file}}"]]}],
				"verify": [{"name": "test", "command": ["go", "test", "./..."], "fix": "go test ./..."}]
			}`),
			0o644,
		))

		config, err := loadHooksConfig(dir)
		require.NoError(t, err)
		require.NotNil(t, config)
		assert.Len(t, config.Format, 1)
		assert.Equal(t, []string{".go"}, config.Format[0].Extensions)
		assert.Len(t, config.Verify, 1)
		assert.Equal(t, "test", config.Verify[0].Name)
	})
}
