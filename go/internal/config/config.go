package config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	enLocale "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	"github.com/spf13/viper"
)

const (
	defaultMaxOpenConns = 10
	defaultMaxIdleConns = 5
)

// Config holds the complete application configuration.
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Logging  LoggingConfig  `mapstructure:"logging"`
}

// ServerConfig holds HTTP server configuration.
type ServerConfig struct {
	Addr            string        `mapstructure:"addr"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"     validate:"gt=0"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"    validate:"gt=0"`
	IdleTimeout     time.Duration `mapstructure:"idle_timeout"     validate:"gt=0"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout" validate:"gt=0"`
}

// DatabaseConfig holds database connection configuration.
type DatabaseConfig struct {
	DSN             string        `mapstructure:"dsn"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"    validate:"gt=0"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"    validate:"gt=0,ltefield=MaxOpenConns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" validate:"gt=0"`
}

// LoggingConfig holds logging configuration.
type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"` // "json" or "text"
}

// Load reads the configuration from file and environment variables.
//
// Configuration is resolved in priority order (lowest to highest):
//  1. Built-in defaults
//  2. Configuration file (./configs/config.yaml or ./config.yaml)
//  3. Environment variables (e.g. SERVER_ADDR overrides server.addr)
func Load() (*Config, error) {
	vpr := viper.New()

	vpr.SetConfigName("config")
	vpr.SetConfigType("yaml")
	vpr.AddConfigPath("./configs/")
	vpr.AddConfigPath(".")

	vpr.SetDefault("server.addr", ":8080")
	vpr.SetDefault("server.read_timeout", "30s")
	vpr.SetDefault("server.write_timeout", "60s")
	vpr.SetDefault("server.idle_timeout", "5m")
	vpr.SetDefault("server.shutdown_timeout", "30s")
	vpr.SetDefault("database.dsn", "")
	vpr.SetDefault("database.max_open_conns", defaultMaxOpenConns)
	vpr.SetDefault("database.max_idle_conns", defaultMaxIdleConns)
	vpr.SetDefault("database.conn_max_lifetime", "1h")
	vpr.SetDefault("logging.level", "info")
	vpr.SetDefault("logging.format", "json")

	vpr.AutomaticEnv()
	vpr.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := vpr.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var cfg Config
	if err := vpr.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode configuration: %w", err)
	}

	locale := enLocale.New()
	uni := ut.New(locale, locale)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		if name := fld.Tag.Get("mapstructure"); name != "" && name != "-" {
			return name
		}

		return fld.Name
	})

	_ = enTrans.RegisterDefaultTranslations(validate, trans)

	if err := validate.Struct(&cfg); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errs := make([]error, 0, len(ve))

			for _, e := range ve {
				path, _ := strings.CutPrefix(e.Namespace(), "Config.")
				errs = append(errs, fmt.Errorf("'%s' (env: %s): %s",
					path,
					strings.ToUpper(strings.ReplaceAll(path, ".", "_")),
					strings.TrimPrefix(e.Translate(trans), e.Field()+" "),
				))
			}

			return nil, fmt.Errorf("configuration validation failed:\n%w", errors.Join(errs...))
		}

		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return &cfg, nil
}
