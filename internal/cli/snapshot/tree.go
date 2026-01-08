package snapshot

import (
	"errors"

	"github.com/spf13/cobra"
)

func newTreeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tree <vm>",
		Short: "Display snapshot hierarchy as a tree",
		Long: `Displays the snapshot hierarchy as an ASCII tree.

Shows parent-child relationships and indicates the current snapshot.
Useful for VMs with complex snapshot trees.

Examples:
  vcli snapshot tree my-vm`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	return cmd
}
