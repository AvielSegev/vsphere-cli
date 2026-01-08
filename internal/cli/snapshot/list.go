package snapshot

import (
	"errors"

	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <vm>",
		Short: "List VM snapshots",
		Long: `Lists all snapshots for a virtual machine in chronological order.

Output includes:
  - Snapshot name
  - Creation date
  - Description
  - Size
  - Current indicator

Examples:
  vcli snapshot list my-vm
  vcli snapshot list my-vm -o json`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	return cmd
}
