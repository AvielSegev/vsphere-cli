package clone

import (
	"errors"

	"github.com/spf13/cobra"
)

var createSnapshot string

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <source-vm> <new-name>",
		Short: "Create a full clone of a VM",
		Long: `Creates a full independent clone of the specified virtual machine.

The clone is created in the same folder, resource pool, and datastore as
the source VM. The cloned VM starts in a powered-off state.

Examples:
  vcli clone create web-server-01 web-server-02
  vcli clone create web-server-01 web-server-02 --snapshot before-upgrade`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	cmd.Flags().StringVar(&createSnapshot, "snapshot", "", "Clone from specific snapshot")

	return cmd
}
