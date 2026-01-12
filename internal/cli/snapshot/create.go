package snapshot

import (
	"fmt"
	"github.com/asegev/vsphere-cli/internal/global"
	"github.com/asegev/vsphere-cli/pkg/config"
	vmware "github.com/kubev2v/assisted-migration-agent/pkg/vmware"

	"github.com/spf13/cobra"
)

var (
	vmName            string
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
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			cfg, err := config.LoadFromEnv()

			c, err := vmware.NewVsphereClient(ctx, cfg.Host, cfg.Username, cfg.Password, cfg.Insecure)
			if err != nil {
				return err
			}

			dcm := vmware.NewVMManager(c)
			if err != nil {
				return err
			}

			finder, err := vmware.NewDatacenterFinder(ctx, c.Client, global.DefaultDatacenterMoid)
			if err != nil {
				return err
			}

			vm, err := finder.FindVMByName(ctx, vmName)
			if err != nil {
				return err
			}

			req := vmware.CreateSnapshotRequest{
				VmMoid:       vm.Reference().Value,
				SnapshotName: createName,
				Description:  createDescription,
				Memory:       createMemory,
				Quiesce:      createQuiesce,
			}

			if err := dcm.CreateSnapshot(ctx, req); err != nil {
				return err
			}

			fmt.Printf("Snapshot %s created\n", vm.Reference().Value)

			return nil
		},
	}

	cmd.Flags().StringVar(&vmName, "vmName", global.DefaultVmName, "Snapshot name (from defaults.go if omitted)")
	cmd.Flags().StringVar(&createName, "name", global.DefaultSnapshotName, "Snapshot name (from defaults.go if omitted)")
	cmd.Flags().StringVar(&createDescription, "description", "", "Snapshot description")
	cmd.Flags().BoolVar(&createMemory, "memory", false, "Include VM memory state")
	cmd.Flags().BoolVar(&createQuiesce, "quiesce", false, "Quiesce filesystem (requires VMware Tools)")

	return cmd
}
