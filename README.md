# vsphere-cli

This project is not affiliated with, endorsed by, or sponsored by VMware, Inc.
VMware, vSphere, and vCenter are registered trademarks of VMware, Inc.

## Building

### Prerequisites

- Go 1.21 or later

### Build

```bash
go build -o vcli ./cmd/vcli
```

### Install

```bash
go install ./cmd/vcli
```

## Usage

### Environment Variables

vcli uses environment variables for authentication:

```bash
export VCLI_HOST=vcenter.example.com
export VCLI_USERNAME=administrator@vsphere.local
export VCLI_PASSWORD=your-password
export VCLI_INSECURE=false  # optional
```

### Commands

```bash
# Display help
vcli --help

# Credentials
vcli credentials test
vcli credentials show

# Snapshots
vcli snapshot create <vm>
vcli snapshot list <vm>
vcli snapshot tree <vm>
vcli snapshot delete <vm> <snapshot-name>
vcli snapshot delete-all <vm> --confirm
vcli snapshot revert <vm> <snapshot-name>
vcli snapshot consolidate <vm>

# Cloning
vcli clone create <source-vm> <new-name>
vcli clone list

# Inspection
vcli inspect vm <vm-name>
```

### Global Flags

- `--host` - Override VCLI_HOST
- `--username` - Override VCLI_USERNAME
- `--password` - Override VCLI_PASSWORD
- `--insecure` - Skip TLS verification
- `--output, -o` - Output format (table, json, yaml)
- `--verbose, -v` - Verbose logging

## Current Status

This is a CLI skeleton with all commands defined but not yet implemented. Each command returns "not implemented" errors. The vSphere integration is stubbed and ready for implementation.
