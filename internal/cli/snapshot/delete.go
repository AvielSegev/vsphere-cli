package snapshot

import (
	"errors"

	"github.com/spf13/cobra"
)

var deleteForce bool

func newDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <vm> <snapshot-name>",
		Short: "Delete a specific snapshot",
		Long: `Deletes a specific snapshot by name.

Prompts for confirmation unless --force flag is used.
Consolidates disks after deletion.

Examples:
  vcli snapshot delete my-vm snapshot-2024-01-01
  vcli snapshot delete my-vm snapshot-2024-01-01 --force`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	cmd.Flags().BoolVar(&deleteForce, "force", false, "Skip confirmation prompt")

	return cmd
}
