# Coding Conventions

## Formatting

- **Formatter**: gofmt + golines (`go fmt ./...`, `go tool golines -w .`)
- **Line length**: 120 characters (enforced by golines)
- **Import organisation**: goimports (applied on save in VS Code)

## Linting

- **Linter**: golangci-lint (`go tool golangci-lint run ./...`)
- All issues must be resolved before stopping; the Stop hook blocks on lint failures.
- Disabled linters and their rationale are documented in `.golangci.yml`.

## Error Handling

### Wrapping

Always wrap errors from external packages with `%w`:

```go
// ✅
return nil, fmt.Errorf("loading config: %w", err)

// ❌
return nil, err
return nil, fmt.Errorf("loading config: %s", err)
```

### Ignoring errors

Use `_ =` when ignoring is intentional:

```go
// ✅
_ = enTrans.RegisterDefaultTranslations(validate, trans)

// ❌ (silently dropped)
enTrans.RegisterDefaultTranslations(validate, trans)
```

### Static vs dynamic errors

- Static messages (no format args) use `errors.New`
- `fmt.Errorf` only when wrapping with `%w` or building dynamic messages

```go
// ✅
return errors.New("not found")
return fmt.Errorf("query %s: %w", id, ErrNotFound)

// ❌
return fmt.Errorf("not found")
```

### Error type checks

Use `errors.As` over direct type assertions — direct assertions break error chains:

```go
// ✅
var notFound viper.ConfigFileNotFoundError
if !errors.As(err, &notFound) { ... }

// ❌
if _, ok := err.(viper.ConfigFileNotFoundError); !ok { ... }
```

## Naming

### Variables

- Narrow scope: short names (`err`, `ok`, `cfg`)
- Package-level or function arguments: descriptive names
- Initialisms use consistent casing: `dsn`/`DSN`, `url`/`URL`, `id`/`ID`

### Constants

Define named constants for magic numbers:

```go
// ✅
const defaultMaxOpenConns = 10
viper.SetDefault("database.max_open_conns", defaultMaxOpenConns)

// ❌
viper.SetDefault("database.max_open_conns", 10)
```

## Godoc Comments

All exported types, functions, and methods require godoc comments:

```go
// Config holds the complete application configuration.
type Config struct { ... }

// Load reads the configuration from file and environment variables.
func Load() (*Config, error) { ... }
```

- First word is the identifier name
- Ends with a period

## Import Groups

goimports enforces this automatically. Three groups separated by blank lines:

```go
import (
    "errors"  // 1. standard library
    "fmt"

    "github.com/spf13/viper"  // 2. external packages

    "github.com/example/template/internal/config"  // 3. internal packages
)
```

## Struct Tags

Use `mapstructure` tag names on config structs so validation errors show
human-readable field names. Align tags when multiple are present:

```go
// ✅
ReadTimeout time.Duration `mapstructure:"read_timeout" validate:"gt=0"`

// ❌ (no tag → error message shows "ReadTimeout")
ReadTimeout time.Duration `validate:"gt=0"`
```
