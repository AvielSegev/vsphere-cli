package snapshot

import (
	"errors"

	"github.com/spf13/cobra"
)

func newConsolidateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "consolidate <vm>",
		Short: "Consolidate snapshot disks",
		Long: `Consolidates snapshot disk files.

Useful when snapshot deletion leaves orphaned disk files.
No-op if consolidation is not needed.

Examples:
  vcli snapshot consolidate my-vm`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	return cmd
}
