package credentials

import (
	"github.com/spf13/cobra"
)

// NewCredentialsCmd creates the credentials command
func NewCredentialsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "credentials",
		Short: "Validate and display credential configuration",
		Long: `The credentials command provides validation and visibility into
the current authentication configuration.

Available subcommands:
  test  - Test connection and validate credentials
  show  - Display current configuration (masked)`,
	}

	cmd.AddCommand(newTestCmd())
	cmd.AddCommand(newShowCmd())

	return cmd
}
