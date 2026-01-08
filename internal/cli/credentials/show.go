package credentials

import (
	"errors"

	"github.com/spf13/cobra"
)

func newShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Display current credential configuration",
		Long: `Displays the current credential configuration.

This command shows:
  - vSphere host
  - Username
  - Password (masked)
  - Insecure flag
  - Source of each value (environment variable or flag)

This does NOT test connectivity. Use 'credentials test' for that.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	return cmd
}
