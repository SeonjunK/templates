# Hook CLI Architecture

## Overview

The hook CLI is a language-agnostic binary that implements Claude Code hooks for project templates. It reads JSON payloads from stdin and emits structured responses to stdout.

## Package Structure

```
cmd/hook/          → CLI entry point (thin router)
internal/
├── hook/          → Core types, input parsing, config loading
├── response/      → Hook response types and builders
├── guard/         → PreToolUse guard handlers
├── format/        → PostToolUse format handler
└── verify/        → Stop verify handler
```

## Command Routing

`cmd/hook/main.go` dispatches subcommands to internal packages:

| Command | Package | Hook Type |
|:--------|:--------|:----------|
| `guard-read` | `internal/guard` | PreToolUse (Read) |
| `guard-write` | `internal/guard` | PreToolUse (Write/Edit) |
| `guard-bash` | `internal/guard` | PreToolUse (Bash) |
| `format` | `internal/format` | PostToolUse (Write/Edit) |
| `verify` | `internal/verify` | Stop |

## Data Flow

### PreToolUse (Guard)

```
stdin (JSON) → hook.ParseInput → guard.Read/Write/Bash
                                      ↓
                              hook.LoadGuardConfig
                                      ↓
                              pattern matching
                                      ↓
                              response.Deny (stdout) or exit 0
```

### PostToolUse (Format)

```
stdin (JSON) → hook.ParseInput → format.Run
                                      ↓
                              hook.LoadHooksConfig
                                      ↓
                              extension matching → exec format commands
```

### Stop (Verify)

```
hook.LoadHooksConfig → verify.Run
                            ↓
                    sequential step execution
                            ↓
                    response.Block or response.Approve (stdout)
```

## Configuration Files

### `.claude/guard.json`

Defines blocked file patterns (read/write) and bash commands/patterns.

```json
{
  "read":  { "blockedPatterns": [".env*", "*.pem"] },
  "write": { "blockedPatterns": [".env*", "*.pem"] },
  "bash":  { "blockedCommands": ["rm -rf /"], "blockedPatterns": ["git push --force"] }
}
```

### `.claude/hooks.json`

Defines format rules (by file extension) and verify steps (sequential execution).

```json
{
  "format": [{ "extensions": [".go"], "commands": [["gofmt", "-w", "{{file}}"]] }],
  "verify": [{ "name": "test", "command": ["go", "test", "./..."], "fix": "go test ./..." }]
}
```

## Response Protocol

### PreToolUse Response

```json
{ "hookSpecificOutput": { "permissionDecision": "deny: reason" } }
```

Decision prefixes: `deny:` (blocks), `warn:` (allows with warning).

### Stop Response

```json
{ "decision": "block", "reason": "verification failed" }
{ "decision": "approve" }
```
