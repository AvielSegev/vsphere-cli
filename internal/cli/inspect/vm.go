package inspect

import (
	"errors"

	"github.com/spf13/cobra"
)

func newVMCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vm <vm-name>",
		Short: "Display comprehensive VM information",
		Long: `Displays comprehensive information about a virtual machine.

Information displayed:
  - General: Name, UUID, power state, guest OS, VMware Tools status
  - Hardware: CPU count, memory, NICs, disks
  - Compute: Host, cluster, resource pool
  - Storage: Datastores, disk sizes, provisioned vs used space
  - Network: Network adapters, MAC addresses, IP addresses
  - Snapshots: Snapshot count, total size
  - Metadata: Creation date, last modified, annotations

Examples:
  vcli inspect vm my-vm
  vcli inspect vm my-vm -o json`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	return cmd
}
