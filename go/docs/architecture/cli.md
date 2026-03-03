# CLI (cmd/app)

## Overview

`cmd/app` is the application entry point using [Cobra](https://github.com/spf13/cobra).

## Directory Structure

```
cmd/app/
├── main.go   # Entry point: calls rootCmd.Execute()
└── root.go   # Root command definition and PersistentPreRunE
```

## Config Integration

Config is loaded once in `PersistentPreRunE` and passed to subcommands via context:

1. `PersistentPreRunE` calls `config.Load()`
2. Stores result with `config.WithConfig(cmd.Context(), cfg)`
3. Subcommands retrieve it with `config.FromContext(cmd.Context())`

## Adding a Subcommand

1. Create `cmd/app/<name>.go` in `package main`
2. Define `var <name>Cmd = &cobra.Command{...}`
3. Register in `init()`: `rootCmd.AddCommand(<name>Cmd)`
4. Access config: `cfg := config.FromContext(cmd.Context())`
