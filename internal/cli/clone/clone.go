package clone

import (
	"github.com/spf13/cobra"
)

// NewCloneCmd creates the clone command
func NewCloneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone virtual machines",
		Long: `Clone virtual machines with basic cloning capabilities.

Available subcommands:
  create  - Create a full clone of a VM
  list    - List cloned VMs`,
	}

	cmd.AddCommand(newCreateCmd())

	return cmd
}
