package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchesExtension(t *testing.T) {
	assert.True(t, matchesExtension(".go", []string{".go"}))
	assert.True(t, matchesExtension(".py", []string{".go", ".py"}))
	assert.False(t, matchesExtension(".rs", []string{".go", ".py"}))
	assert.False(t, matchesExtension(".go", nil))
}

func TestExpandFileArg(t *testing.T) {
	result := expandFileArg([]string{"gofmt", "-w", "{{file}}"}, "/project/main.go")
	assert.Equal(t, []string{"gofmt", "-w", "/project/main.go"}, result)

	result = expandFileArg([]string{"echo", "hello"}, "/project/main.go")
	assert.Equal(t, []string{"echo", "hello"}, result)
}

func TestFormatCmd_SkipsWhenFilePathMissing(t *testing.T) {
	fakeStdin(t, `{}`)
	out := captureStdout(t, func() { _ = formatCmd() })
	assert.Empty(t, out)
}

func TestFormatCmd_SkipsWhenNoConfig(t *testing.T) {
	t.Setenv("CLAUDE_PROJECT_DIR", "")
	fakeStdin(t, `{"tool_input":{"file_path":"main.go"}}`)
	out := captureStdout(t, func() { _ = formatCmd() })
	assert.Empty(t, out)
}

func TestFormatCmd_SkipsNonMatchingExtension(t *testing.T) {
	dir := setupHooksEnv(t, `{"format":[{"extensions":[".go"],"commands":[["gofmt","-w","{{file}}"]]}]}`)
	_ = dir
	fakeStdin(t, `{"tool_input":{"file_path":"main.py"}}`)
	out := captureStdout(t, func() { _ = formatCmd() })
	assert.Empty(t, out)
}
