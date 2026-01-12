package clone

import (
	"fmt"
	"github.com/asegev/vsphere-cli/internal/global"
	"github.com/asegev/vsphere-cli/pkg/config"
	vmware "github.com/kubev2v/assisted-migration-agent/pkg/vmware"

	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <source-vm> <new-name>",
		Short: "delete a linked clone of a VM",
		Args:  cobra.NoArgs,
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

			vm, err := finder.FindVMByName(ctx, cloneName)
			if err != nil {
				return err
			}

			req := vmware.RemoveLinkedCloneRequest{
				VmMoid: vm.Reference().Value,
			}

			if err := dcm.RemoveLinkedClone(ctx, req); err != nil {
				return err
			}

			fmt.Printf("Successfuly delete VM %s\n", vm.Reference().Value)

			return nil

		},
	}

	cmd.Flags().StringVar(&vmName, "vmName", global.DefaultVmName, "from defaults.go if omitted")
	cmd.Flags().StringVar(&snapshotName, "snapshotName", global.DefaultSnapshotName, "from defaults.go if omitted")
	cmd.Flags().StringVar(&cloneName, "cloneName", global.DefaultClonedVmName, "from defaults.go if omitted")

	return cmd
}
