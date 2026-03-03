package config

import "context"

// contextKey is an unexported type for context keys scoped to this package.
type contextKey struct{}

// WithConfig returns a new context derived from ctx with cfg stored inside.
func WithConfig(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, contextKey{}, cfg)
}

// FromContext retrieves the Config value stored by WithConfig.
// It panics if no Config is present, indicating a programming error.
func FromContext(ctx context.Context) *Config {
	cfg, ok := ctx.Value(contextKey{}).(*Config)
	if !ok || cfg == nil {
		panic("config: no Config in context; ensure PersistentPreRunE has run")
	}

	return cfg
}
