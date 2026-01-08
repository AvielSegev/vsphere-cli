package cli

import (
	"fmt"

	"github.com/asegev/vsphere-cli/pkg/config"
	"github.com/asegev/vsphere-cli/pkg/output"
	"github.com/spf13/cobra"
)

var (
	// Global flags
	flagHost     string
	flagUsername string
	flagPassword string
	flagInsecure bool
	flagOutput   string
	flagVerbose  bool

	// Global config
	globalConfig *config.Config
	globalFormat output.Format
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "vcli",
	Short: "A CLI tool for interacting with VMware vSphere/ESXi",
	Long: `vcli is a command-line tool for managing VMware vSphere environments.

It provides commands for snapshot management, VM cloning, VM inspection,
and credential validation.

Authentication is configured via environment variables:
  VCLI_HOST      - vCenter/ESXi host address
  VCLI_USERNAME  - Authentication username
  VCLI_PASSWORD  - Authentication password
  VCLI_INSECURE  - Skip TLS verification (optional, default: false)`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Skip validation for commands that don't need it
		if cmd.Name() == "help" || cmd.Name() == "vcli" {
			return nil
		}

		// Load config from environment
		cfg, err := config.LoadFromEnv()
		if err != nil {
			return err
		}

		// Override with flags if provided
		if flagHost != "" {
			cfg.Host = flagHost
		}
		if flagUsername != "" {
			cfg.Username = flagUsername
		}
		if flagPassword != "" {
			cfg.Password = flagPassword
		}
		if cmd.Flags().Changed("insecure") {
			cfg.Insecure = flagInsecure
		}

		// Validate config (skip for credentials show command)
		if cmd.Name() != "show" {
			if err := cfg.Validate(); err != nil {
				return fmt.Errorf("configuration error: %w\nSet environment variables or use flags", err)
			}
		}

		globalConfig = cfg

		// Set output format
		switch flagOutput {
		case "table", "json", "yaml":
			globalFormat = output.Format(flagOutput)
		default:
			return fmt.Errorf("invalid output format: %s (must be table, json, or yaml)", flagOutput)
		}

		return nil
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVar(&flagHost, "host", "", "vCenter/ESXi host (overrides VCLI_HOST)")
	rootCmd.PersistentFlags().StringVar(&flagUsername, "username", "", "Username (overrides VCLI_USERNAME)")
	rootCmd.PersistentFlags().StringVar(&flagPassword, "password", "", "Password (overrides VCLI_PASSWORD)")
	rootCmd.PersistentFlags().BoolVar(&flagInsecure, "insecure", false, "Skip TLS verification (overrides VCLI_INSECURE)")
	rootCmd.PersistentFlags().StringVarP(&flagOutput, "output", "o", "table", "Output format (table, json, yaml)")
	rootCmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "Verbose output")
}

// GetConfig returns the global config
func GetConfig() *config.Config {
	return globalConfig
}

// GetFormat returns the global output format
func GetFormat() output.Format {
	return globalFormat
}
