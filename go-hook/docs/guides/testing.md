# Testing Guide

## Running Tests

```bash
go test -race ./...               # all tests
go test -v -race ./...            # verbose output
go test -race ./internal/hook/    # specific package
```

## File Structure

```
internal/
└── <package>/
    ├── <package>.go
    └── <package>_test.go
```

Test files live alongside the source file they test. Use the same package name.

## Test Organisation

Use table-driven tests for multiple input/output scenarios:

```go
tests := []struct {
    name string
    yaml string
    err  error
}{
    {name: "zero read_timeout", yaml: "server:\n  read_timeout: 0s\n", err: ...},
}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        ...
    })
}
```

## Assertions

Use `require` when test execution should stop on failure, `assert` otherwise:

| Use | When |
|:----|:-----|
| `require.NoError(t, err)` | Error check before dereferencing the result |
| `require.EqualError(t, err, ...)` | Error message must match to continue |
| `assert.Equal(t, ...)` | Non-fatal value comparison |
| `assert.Nil(t, cfg)` | Non-fatal nil check |

## Filesystem Isolation

Use `t.TempDir()` for temporary files and `t.Chdir()` to set the working
directory. Both are cleaned up automatically after the test.

```go
tempDir := t.TempDir()
t.Chdir(tempDir)
require.NoError(t, os.WriteFile(
    filepath.Join(tempDir, "config.yaml"),
    []byte("server:\n  read_timeout: 15s\n"),
    0600,
))
```

File permissions: use `0600` for config files written in tests.

## Environment Variables

Use `t.Setenv()` — changes are automatically reverted after the test:

```go
t.Setenv("SERVER_ADDR", ":9090")
cfg, err := Load()
require.NoError(t, err)
assert.Equal(t, ":9090", cfg.Server.Addr)
```

## Expected Errors

Use `require.EqualError` to assert the exact error message:

```go
cfg, err := Load()
require.EqualError(t, err,
    "configuration validation failed:\n'server.read_timeout' (env: SERVER_READ_TIMEOUT): must be greater than 0",
)
assert.Nil(t, cfg)
```

Use `assert.Contains` when only a substring needs to match:

```go
cfg, err := Load()
assert.Contains(t, err.Error(), "unable to decode configuration")
assert.Nil(t, cfg)
```

## Static Error Strings

Test expected error messages use `errors.New`, not `fmt.Errorf`:

```go
// good
err: errors.New("configuration validation failed:\n'server.read_timeout' ..."),

// bad
err: fmt.Errorf("configuration validation failed:\n'server.read_timeout' ..."),
```
