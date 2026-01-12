package credentials

import (
	"fmt"
	"github.com/asegev/vsphere-cli/pkg/config"

	vmware "github.com/kubev2v/assisted-migration-agent/pkg/vmware"
	"github.com/spf13/cobra"
)

const (
	defaultDatacenterMoid = "datacenter-3"
	defaultClusterMoid    = "domain-c34"
	defaultVmName         = "asegev-ubuntu-for-workflow-runs"
)

var requiredPrivileges = []string{
	"VirtualMachine.Provisioning.Clone",
	"VirtualMachine.Inventory.CreateFromExisting",
	"VirtualMachine.State.CreateSnapshot",
	"VirtualMachine.State.RemoveSnapshot",
}

func newTestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "Test vSphere connection and validate credentials",
		Long: `Tests connectivity to vSphere and validates authentication.

This command will:
  - Connect to the vCenter/ESXi host
  - Authenticate with provided credentials
  - Verify API accessibility
  - Display connection details

Exit code 0 on success, 1 on failure.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			cfg, err := config.LoadFromEnv()

			c, err := vmware.NewVsphereClient(ctx, cfg.Host, cfg.Username, cfg.Password, cfg.Insecure)
			if err != nil {
				return err
			}

			finder, err := vmware.NewFinder(ctx, c.Client, defaultDatacenterMoid)
			if err != nil {
				return err
			}

			vm, err := finder.FindVMByName(ctx, defaultVmName)
			if err != nil {
				return err
			}

			if err := vmware.ValidateUserPrivileges(ctx, c.Client, vm.Reference(), requiredPrivileges, cfg.Username); err != nil {
				return err
			}

			fmt.Println("Successfully validated credentials.")

			return nil
		},
	}

	return cmd
}
