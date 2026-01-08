# vcli - VMware CLI Tool Design

**Date:** 2026-01-08
**Status:** Approved

## Overview

`vcli` is a command-line tool for interacting with VMware vSphere/ESXi environments. It provides snapshot management, VM cloning, VM inspection, and credential validation capabilities.

## Command Structure

The tool uses a feature-based command organization:

```
vcli snapshot <subcommand>
vcli clone <subcommand>
vcli inspect <subcommand>
vcli credentials <subcommand>
```

## Authentication

All authentication uses environment variables:

- `VCLI_HOST` - vCenter/ESXi host address
- `VCLI_USERNAME` - Authentication username
- `VCLI_PASSWORD` - Authentication password
- `VCLI_INSECURE` - Optional: skip TLS verification (default: false)

### Global Flags

Available on all commands:

- `--host` - Override VCLI_HOST
- `--username` - Override VCLI_USERNAME
- `--password` - Override VCLI_PASSWORD
- `--insecure` - Override VCLI_INSECURE
- `--output, -o` - Output format: `table` (default), `json`, `yaml`
- `--verbose, -v` - Verbose logging

### Validation

The root command validates environment variables are set (or overridden via flags) before executing any subcommand.

## Output Philosophy

- Default to human-readable table format
- Support JSON/YAML for scripting/automation
- Use color coding for status indicators (success=green, warning=yellow, error=red)
- Progress indicators for long-running operations

## Credentials Commands

### `vcli credentials test`

Validates credentials and tests vSphere connectivity.

**Behavior:**
- Attempts to connect to vCenter/ESXi using current credentials
- Validates authentication and basic API connectivity
- Reports connection success/failure with details
- Exit code 0 on success, 1 on failure

**Output Example:**
```
Testing connection to vcenter.example.com...
✓ Connection successful
✓ Authenticated as: administrator@vsphere.local
✓ vCenter version: 7.0.3
✓ API accessible
```

### `vcli credentials show`

Displays current credential configuration.

**Behavior:**
- Displays current credential configuration (from env vars or flags)
- Masks password (shows only first 2 and last 2 characters)
- Shows which source each value came from (env var vs flag)
- Does NOT test connectivity (use `test` for that)

**Output Example:**
```
Current Configuration:
  Host:     vcenter.example.com (from VCLI_HOST)
  Username: administrator@vsphere.local (from VCLI_USERNAME)
  Password: pa••••••rd (from VCLI_PASSWORD)
  Insecure: false (default)
```

## Snapshot Commands

### `vcli snapshot create <vm> [flags]`

Creates a snapshot of the specified VM.

**Flags:**
- `--name <name>` - Optional snapshot name (auto-generated if omitted: `snapshot-YYYY-MM-DD-HHMMSS`)
- `--description <desc>` - Optional description
- `--memory` - Include VM memory state (default: false)
- `--quiesce` - Quiesce filesystem before snapshot (requires VMware Tools, default: false)

### `vcli snapshot list <vm>`

Lists all snapshots for a VM in chronological order.

**Output:**
- Shows: name, creation date, description, size, current indicator
- Table format by default, supports JSON/YAML

### `vcli snapshot tree <vm>`

Displays snapshot hierarchy as ASCII tree.

**Output:**
- Shows parent-child relationships
- Indicates current snapshot with marker
- Useful for VMs with complex snapshot trees

### `vcli snapshot delete <vm> <snapshot-name>`

Deletes a specific snapshot by name.

**Behavior:**
- Prompts for confirmation unless `--force` flag used
- Consolidates disks after deletion

### `vcli snapshot delete-all <vm>`

Removes all snapshots from a VM.

**Behavior:**
- Requires `--confirm` flag to prevent accidents
- Useful for cleanup operations

### `vcli snapshot revert <vm> <snapshot-name>`

Reverts VM to specified snapshot state.

**Behavior:**
- Prompts for confirmation unless `--force` flag used
- Can revert to older snapshots (not just current)

### `vcli snapshot consolidate <vm>`

Consolidates snapshot disk files.

**Behavior:**
- Useful when snapshot deletion leaves orphaned disks
- No-op if consolidation not needed

## Clone Commands

### `vcli clone create <source-vm> <new-name> [flags]`

Creates a full clone of the specified VM.

**Arguments:**
- `<source-vm>` - Name of the VM to clone
- `<new-name>` - Name for the new cloned VM

**Flags:**
- `--snapshot <name>` - Optional: clone from a specific snapshot instead of current state

**Behavior:**
- Creates a full independent clone (not linked)
- Clone is created in same folder as source VM
- Clone is created in same resource pool as source VM
- Clone is created in same datastore as source VM
- Clone starts in powered-off state
- All disks are cloned (full copy)
- Shows progress during cloning operation

**Output Example:**
```
Cloning VM 'web-server-01' to 'web-server-02'...
⠋ Creating clone... 45% (copying disk 1 of 2)
✓ Clone created successfully
  Name: web-server-02
  State: Powered Off
  Disks: 2 (50 GB total)
```

### `vcli clone list`

Lists VMs that were created as clones.

**Output:**
- Shows: clone name, source VM, creation date, state
- Useful for tracking cloned VMs
- Note: This tracks clones created via vcli or relies on vSphere clone metadata

## Inspect Commands

### `vcli inspect vm <vm-name>`

Displays comprehensive information about a VM.

**Information Displayed:**
- **General:** Name, UUID, power state, guest OS, VMware Tools status
- **Hardware:** CPU count, memory (MB), number of NICs, number of disks
- **Compute:** Host, cluster, resource pool
- **Storage:** Datastore(s), disk sizes, provisioned vs used space
- **Network:** Network adapter details, MAC addresses, IP addresses (if Tools running)
- **Snapshots:** Snapshot count, total snapshot size
- **Metadata:** Creation date, last modified, annotation/notes

**Output Example (Table Format):**
```
VM: web-server-01
================================================================================
General Information:
  Power State:    Powered On
  Guest OS:       Ubuntu Linux (64-bit)
  VMware Tools:   Running (version 11.3.5)
  UUID:           420d3c4f-8c2e-4b9a-a3d1-9f8e7c6b5a4d

Hardware Configuration:
  CPUs:           4
  Memory:         8192 MB
  NICs:           1
  Disks:          2

Storage:
  Datastore:      datastore1
  Disk 1:         50 GB (thin provisioned, 23 GB used)
  Disk 2:         100 GB (thick provisioned)
  Snapshots:      3 snapshots (12 GB total)

Network:
  NIC 1:          VM Network (vmxnet3)
    MAC:          00:50:56:9a:12:34
    IP:           192.168.1.100
```

**JSON Output:**
- Structured data for programmatic access
- All fields available for scripting/automation

## Project Structure

```
vcli/
├── cmd/
│   └── vcli/
│       └── main.go              # Entry point only
├── internal/
│   └── cli/
│       ├── root.go              # Root command setup
│       ├── credentials/
│       │   ├── credentials.go   # Credentials command group
│       │   ├── test.go
│       │   └── show.go
│       ├── snapshot/
│       │   ├── snapshot.go      # Snapshot command group
│       │   ├── create.go
│       │   ├── delete.go
│       │   ├── list.go
│       │   ├── revert.go
│       │   ├── tree.go
│       │   ├── delete_all.go
│       │   └── consolidate.go
│       ├── clone/
│       │   ├── clone.go         # Clone command group
│       │   ├── create.go
│       │   └── list.go
│       └── inspect/
│           ├── inspect.go       # Inspect command group
│           └── vm.go
├── pkg/
│   ├── config/
│   │   └── config.go            # Environment variable handling
│   ├── vsphere/
│   │   └── client.go            # vSphere client wrapper (stub for now)
│   └── output/
│       └── formatter.go         # Output formatting (table/json/yaml)
├── go.mod
└── go.sum
```

## Dependencies

- `github.com/spf13/cobra` - CLI framework
- `github.com/spf13/viper` - Configuration management (for flags/env vars)
- `github.com/vmware/govmomi` - vSphere API client (implementation phase)
- `github.com/olekukonko/tablewriter` - Table formatting
- `gopkg.in/yaml.v3` - YAML output

## Implementation Notes

### Skeleton Phase

- Each command file will have a `RunE` function that returns `errors.New("not implemented")`
- vSphere client package will have interface definitions but no actual implementation
- This allows the CLI structure to be tested without vSphere connectivity

### Error Handling

- All commands return errors via Cobra's `RunE`
- Connection errors clearly distinguished from command errors
- Exit codes:
  - 0 = success
  - 1 = error
  - 2 = invalid arguments
