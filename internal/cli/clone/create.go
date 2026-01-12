package clone

import (
	"fmt"
	"github.com/asegev/vsphere-cli/internal/global"
	"github.com/asegev/vsphere-cli/pkg/config"
	vmware "github.com/kubev2v/assisted-migration-agent/pkg/vmware"

	"github.com/spf13/cobra"
)

var (
	vmName       string
	snapshotName string
)

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <source-vm> <new-name>",
		Short: "Create a full clone of a VM",
		Long: `Creates a full independent clone of the specified virtual machine.

The clone is created in the same folder, resource pool, and datastore as
the source VM. The cloned VM starts in a powered-off state.

Examples:
  vcli clone create web-server-01 web-server-02
  vcli clone create web-server-01 web-server-02 --snapshot before-upgrade`,
		Args: cobra.NoArgs,
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

			snapshotRef, err := vm.FindSnapshot(ctx, snapshotName)
			if err != nil {
				return err
			}

			req := vmware.CreateLinkedCloneRequest{
				VmMoid:      vm.Reference().Value,
				SnapshotRef: snapshotRef,
				CloneName:   fmt.Sprintf("clone-%s", vmName),
			}

			if err := dcm.CreateLinkedClone(ctx, req); err != nil {
				return err
			}

			fmt.Printf("Successfuly cloned VM %s\n", vm.Reference().Value)

			return nil

		},
	}

	cmd.Flags().StringVar(&vmName, "vmName", global.DefaultVmName, "Snapshot name (from defaults.go if omitted)")
	cmd.Flags().StringVar(&snapshotName, "name", global.DefaultSnapshotName, "Snapshot name (from defaults.go if omitted)")

	return cmd
}
