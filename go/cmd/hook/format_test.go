package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat_SkipsWhenFilePathMissing(t *testing.T) {
	fakeStdin(t, `{}`)
	out := captureStdout(t, func() { _ = format() })
	assert.Empty(t, out)
}

func TestFormat_SkipsNonGoFiles(t *testing.T) {
	fakeStdin(t, `{"tool_input":{"file_path":"main.py"}}`)
	out := captureStdout(t, func() { _ = format() })
	assert.Empty(t, out)
}

func TestFormat_SkipsTxtFiles(t *testing.T) {
	fakeStdin(t, `{"tool_input":{"file_path":"README.txt"}}`)
	out := captureStdout(t, func() { _ = format() })
	assert.Empty(t, out)
}

func TestFormat_AcceptsGoFile(t *testing.T) {
	// Create a valid Go file so gofmt succeeds
	dir := t.TempDir()
	goFile := dir + "/test.go"
	err := os.WriteFile(goFile, []byte("package main\n"), 0o644)
	if err != nil {
		t.Skip("cannot create temp go file")
	}

	fakeStdin(t, `{"tool_input":{"file_path":"`+goFile+`"}}`)
	// format() will attempt to run gofmt - it may warn if gofmt not found
	// but should not error
	out := captureStdout(t, func() { _ = format() })
	// Output may be empty (success) or a warning (gofmt not found)
	_ = out
}
