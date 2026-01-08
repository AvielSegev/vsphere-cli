package inspect

import (
	"github.com/spf13/cobra"
)

// NewInspectCmd creates the inspect command
func NewInspectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inspect",
		Short: "Inspect virtual machines and resources",
		Long: `Inspect virtual machines and display detailed information.

Available subcommands:
  vm  - Display comprehensive VM information`,
	}

	cmd.AddCommand(newVMCmd())

	return cmd
}
