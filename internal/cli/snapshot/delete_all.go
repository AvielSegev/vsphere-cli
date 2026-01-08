package snapshot

import (
	"errors"

	"github.com/spf13/cobra"
)

var deleteAllConfirm bool

func newDeleteAllCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-all <vm>",
		Short: "Delete all snapshots from a VM",
		Long: `Removes all snapshots from a virtual machine.

Requires --confirm flag to prevent accidents.
Useful for cleanup operations.

Examples:
  vcli snapshot delete-all my-vm --confirm`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	cmd.Flags().BoolVar(&deleteAllConfirm, "confirm", false, "Confirm deletion of all snapshots (required)")
	cmd.MarkFlagRequired("confirm")

	return cmd
}
