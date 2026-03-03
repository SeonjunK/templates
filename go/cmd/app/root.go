package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/example/template/internal/config"
)

//nolint:gochecknoglobals // rootCmd is package-level per cobra convention.
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Template application.",
	Long:  "Template application demonstrating Cobra and Viper integration.",
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}

		cmd.SetContext(config.WithConfig(cmd.Context(), cfg))

		return nil
	},
	RunE: func(cmd *cobra.Command, _ []string) error {
		cfg := config.FromContext(cmd.Context())

		_, err := fmt.Fprintf(cmd.OutOrStdout(), "server addr: %s\n", cfg.Server.Addr)
		if err != nil {
			return fmt.Errorf("writing output: %w", err)
		}

		return nil
	},
}
