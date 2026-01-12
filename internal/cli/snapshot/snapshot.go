package snapshot

import (
	"github.com/spf13/cobra"
)

// NewSnapshotCmd creates the snapshot command
func NewSnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snapshot",
		Short: "Manage VM snapshots",
		Long: `Manage VM snapshots with comprehensive snapshot management features.

Available subcommands:
  create       - Create a new snapshot
  tree         - Display snapshot hierarchy
  delete       - Delete a specific snapshot`,
	}

	cmd.AddCommand(newCreateCmd())
	cmd.AddCommand(newTreeCmd())
	cmd.AddCommand(newDeleteCmd())

	return cmd
}
