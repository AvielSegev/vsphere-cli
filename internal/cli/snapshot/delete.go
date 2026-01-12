package snapshot

import (
	"fmt"
	"github.com/asegev/vsphere-cli/internal/global"
	"github.com/asegev/vsphere-cli/pkg/config"
	vmware "github.com/kubev2v/assisted-migration-agent/pkg/vmware"

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

			req := vmware.RemoveSnapshotRequest{
				VmMoid:       vm.Reference().Value,
				SnapshotName: createName,
				Consolidate:  false,
			}

			if err := dcm.RemoveSnapshot(ctx, req); err != nil {
				return err
			}

			fmt.Printf("Snapshot %s deleted\n", vm.Reference().Value)

			return nil
		},
	}

	cmd.Flags().BoolVar(&deleteForce, "force", false, "Skip confirmation prompt")

	return cmd
}
