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
  list         - List snapshots for a VM
  tree         - Display snapshot hierarchy
  delete       - Delete a specific snapshot
  delete-all   - Delete all snapshots
  revert       - Revert to a snapshot
  consolidate  - Consolidate snapshot disks`,
	}

	cmd.AddCommand(newCreateCmd())
	cmd.AddCommand(newTreeCmd())
	cmd.AddCommand(newDeleteCmd())
	cmd.AddCommand(newConsolidateCmd())

	return cmd
}
