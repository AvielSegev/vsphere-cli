package clone

import (
	"errors"

	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List cloned VMs",
		Long: `Lists virtual machines that were created as clones.

Output includes:
  - Clone name
  - Source VM
  - Creation date
  - Power state

Examples:
  vcli clone list
  vcli clone list -o json`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	return cmd
}
