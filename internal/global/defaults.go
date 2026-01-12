package global

import "fmt"

const (
	DefaultDatacenterMoid = "datacenter-3"
	DefaultClusterMoid    = "domain-c34"
	DefaultVmName         = "asegev-ubuntu-for-workflow-runs"
	DefaultSnapshotName   = "assisted-migration-deep-inspector-snapshot"
)

var DefaultClonedVmName = fmt.Sprintf("clone-%s", DefaultVmName)
