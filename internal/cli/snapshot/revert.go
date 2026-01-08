package snapshot

import (
	"errors"

	"github.com/spf13/cobra"
)

var revertForce bool

func newRevertCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revert <vm> <snapshot-name>",
		Short: "Revert VM to a snapshot",
		Long: `Reverts a virtual machine to the specified snapshot state.

Prompts for confirmation unless --force flag is used.
Can revert to any snapshot in the tree, not just the current one.

Examples:
  vcli snapshot revert my-vm snapshot-2024-01-01
  vcli snapshot revert my-vm snapshot-2024-01-01 --force`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	cmd.Flags().BoolVar(&revertForce, "force", false, "Skip confirmation prompt")

	return cmd
}
