package snapshot

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	createName        string
	createDescription string
	createMemory      bool
	createQuiesce     bool
)

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <vm>",
		Short: "Create a VM snapshot",
		Long: `Creates a snapshot of the specified virtual machine.

The snapshot name is auto-generated if not provided (snapshot-YYYY-MM-DD-HHMMSS).

Examples:
  vcli snapshot create my-vm
  vcli snapshot create my-vm --name "before-upgrade"
  vcli snapshot create my-vm --memory --description "With memory state"`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	cmd.Flags().StringVar(&createName, "name", "", "Snapshot name (auto-generated if omitted)")
	cmd.Flags().StringVar(&createDescription, "description", "", "Snapshot description")
	cmd.Flags().BoolVar(&createMemory, "memory", false, "Include VM memory state")
	cmd.Flags().BoolVar(&createQuiesce, "quiesce", false, "Quiesce filesystem (requires VMware Tools)")

	return cmd
}
