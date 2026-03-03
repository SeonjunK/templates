package config

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Run("error for unreadable config file", func(t *testing.T) {
		tempDir := t.TempDir()
		t.Chdir(tempDir)
		require.NoError(t, os.WriteFile(
			filepath.Join(tempDir, "config.yaml"),
			[]byte(""),
			0000,
		))

		cfg, err := Load()
		assert.Contains(t, err.Error(), "error reading config file")
		assert.Nil(t, cfg)
	})

	t.Run("error for invalid duration in config file", func(t *testing.T) {
		tempDir := t.TempDir()
		t.Chdir(tempDir)
		require.NoError(t, os.WriteFile(
			filepath.Join(tempDir, "config.yaml"),
			[]byte("server:\n  read_timeout: not-a-duration\n"),
			0600,
		))

		cfg, err := Load()
		assert.Contains(t, err.Error(), "unable to decode configuration")
		assert.Nil(t, cfg)
	})

	validationTests := []struct {
		yaml string
		err  error
	}{
		{
			yaml: "server:\n  read_timeout: 0s\n",
			err: errors.New(
				"configuration validation failed:\n'server.read_timeout' (env: SERVER_READ_TIMEOUT): must be greater than 0",
			),
		},
		{
			yaml: "server:\n  write_timeout: 0s\n",
			err: errors.New(
				"configuration validation failed:\n'server.write_timeout' (env: SERVER_WRITE_TIMEOUT): must be greater than 0",
			),
		},
		{
			yaml: "server:\n  idle_timeout: 0s\n",
			err: errors.New(
				"configuration validation failed:\n'server.idle_timeout' (env: SERVER_IDLE_TIMEOUT): must be greater than 0",
			),
		},
		{
			yaml: "server:\n  shutdown_timeout: 0s\n",
			err: errors.New(
				"configuration validation failed:\n'server.shutdown_timeout' (env: SERVER_SHUTDOWN_TIMEOUT): must be greater than 0",
			),
		},
		{
			yaml: "database:\n  dsn: \"postgres://localhost/mydb\"\n  max_open_conns: 5\n  max_idle_conns: 10\n",
			err: errors.New(
				"configuration validation failed:\n'database.max_idle_conns' (env: DATABASE_MAX_IDLE_CONNS): must be less than or equal to MaxOpenConns",
			),
		},
	}
	for _, tt := range validationTests {
		t.Run(tt.yaml, func(t *testing.T) {
			tempDir := t.TempDir()
			t.Chdir(tempDir)
			require.NoError(t, os.WriteFile(
				filepath.Join(tempDir, "config.yaml"),
				[]byte(tt.yaml),
				0600,
			))

			cfg, err := Load()
			require.EqualError(t, err, tt.err.Error())
			assert.Nil(t, cfg)
		})
	}

	t.Run("success with defaults", func(t *testing.T) {
		cfg, err := Load()
		require.NoError(t, err)
		assert.Equal(t, Config{
			Server: ServerConfig{
				Addr:            ":8080",
				ReadTimeout:     30 * time.Second,
				WriteTimeout:    60 * time.Second,
				IdleTimeout:     5 * time.Minute,
				ShutdownTimeout: 30 * time.Second,
			},
			Database: DatabaseConfig{
				DSN:             "",
				MaxOpenConns:    10,
				MaxIdleConns:    5,
				ConnMaxLifetime: time.Hour,
			},
			Logging: LoggingConfig{
				Level:  "info",
				Format: "json",
			},
		}, *cfg)
	})

	t.Run("success with all configurations", func(t *testing.T) {
		tempDir := t.TempDir()
		t.Chdir(tempDir)

		err := os.WriteFile(filepath.Join(tempDir, "config.yaml"), []byte(`
server:
  addr: ":9090"
  read_timeout: 15s
  write_timeout: 30s
  idle_timeout: 2m
  shutdown_timeout: 1m

database:
  dsn: "postgres://user:pass@localhost:5432/mydb?sslmode=disable"
  max_open_conns: 20
  max_idle_conns: 10
  conn_max_lifetime: 30m

logging:
  level: "debug"
  format: "text"
`), 0600)
		require.NoError(t, err)

		cfg, err := Load()
		require.NoError(t, err)
		assert.Equal(t, Config{
			Server: ServerConfig{
				Addr:            ":9090",
				ReadTimeout:     15 * time.Second,
				WriteTimeout:    30 * time.Second,
				IdleTimeout:     2 * time.Minute,
				ShutdownTimeout: time.Minute,
			},
			Database: DatabaseConfig{
				DSN:             "postgres://user:pass@localhost:5432/mydb?sslmode=disable",
				MaxOpenConns:    20,
				MaxIdleConns:    10,
				ConnMaxLifetime: 30 * time.Minute,
			},
			Logging: LoggingConfig{
				Level:  "debug",
				Format: "text",
			},
		}, *cfg)
	})

	t.Run("success with all environment variables", func(t *testing.T) {
		tempDir := t.TempDir()
		t.Chdir(tempDir)

		t.Setenv("SERVER_ADDR", ":9090")
		t.Setenv("SERVER_READ_TIMEOUT", "15s")
		t.Setenv("SERVER_WRITE_TIMEOUT", "30s")
		t.Setenv("SERVER_IDLE_TIMEOUT", "2m")
		t.Setenv("SERVER_SHUTDOWN_TIMEOUT", "1m")
		t.Setenv("DATABASE_DSN", "postgres://env-user:env-pass@localhost/envdb")
		t.Setenv("DATABASE_MAX_OPEN_CONNS", "20")
		t.Setenv("DATABASE_MAX_IDLE_CONNS", "10")
		t.Setenv("DATABASE_CONN_MAX_LIFETIME", "30m")
		t.Setenv("LOGGING_LEVEL", "warn")
		t.Setenv("LOGGING_FORMAT", "text")

		cfg, err := Load()
		require.NoError(t, err)
		assert.Equal(t, Config{
			Server: ServerConfig{
				Addr:            ":9090",
				ReadTimeout:     15 * time.Second,
				WriteTimeout:    30 * time.Second,
				IdleTimeout:     2 * time.Minute,
				ShutdownTimeout: time.Minute,
			},
			Database: DatabaseConfig{
				DSN:             "postgres://env-user:env-pass@localhost/envdb",
				MaxOpenConns:    20,
				MaxIdleConns:    10,
				ConnMaxLifetime: 30 * time.Minute,
			},
			Logging: LoggingConfig{
				Level:  "warn",
				Format: "text",
			},
		}, *cfg)
	})

	t.Run("success with environment variables overriding config file", func(t *testing.T) {
		tempDir := t.TempDir()
		t.Chdir(tempDir)

		err := os.WriteFile(filepath.Join(tempDir, "config.yaml"), []byte(`
server:
  addr: ":8080"
  read_timeout: 15s
database:
  dsn: "postgres://file-user:file-pass@localhost/filedb"
logging:
  level: "debug"
`), 0600)
		require.NoError(t, err)

		t.Setenv("SERVER_ADDR", ":9999")
		t.Setenv("DATABASE_DSN", "postgres://env-user:env-pass@localhost/envdb")

		cfg, err := Load()
		require.NoError(t, err)
		assert.Equal(t, Config{
			Server: ServerConfig{
				Addr:            ":9999",
				ReadTimeout:     15 * time.Second,
				WriteTimeout:    60 * time.Second,
				IdleTimeout:     5 * time.Minute,
				ShutdownTimeout: 30 * time.Second,
			},
			Database: DatabaseConfig{
				DSN:             "postgres://env-user:env-pass@localhost/envdb",
				MaxOpenConns:    10,
				MaxIdleConns:    5,
				ConnMaxLifetime: time.Hour,
			},
			Logging: LoggingConfig{
				Level:  "debug",
				Format: "json",
			},
		}, *cfg)
	})
}
