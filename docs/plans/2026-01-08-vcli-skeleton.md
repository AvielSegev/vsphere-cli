# vcli CLI Skeleton Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Create a complete CLI skeleton for vcli with all commands, flags, and help text defined, returning "not implemented" stubs.

**Architecture:** Feature-based Cobra CLI with modular command structure. Commands organized in internal/cli with pkg containing reusable logic for config, vSphere client interface, and output formatting.

**Tech Stack:** Go 1.21+, Cobra (CLI), Viper (config), tablewriter (output), govmomi (vSphere SDK stub)

---

## Task 1: Initialize Go Module and Dependencies

**Files:**
- Create: `go.mod`
- Create: `go.sum`

**Step 1: Initialize Go module**

Run: `go mod init github.com/asegev/vsphere-cli`

Expected: Creates go.mod file

**Step 2: Add dependencies**

Run:
```bash
go get github.com/spf13/cobra@latest
go get github.com/spf13/viper@latest
go get github.com/olekukonko/tablewriter@latest
go get gopkg.in/yaml.v3@latest
go get github.com/vmware/govmomi@latest
```

Expected: Dependencies added to go.mod and go.sum

**Step 3: Verify dependencies**

Run: `go mod tidy`

Expected: All dependencies resolved

**Step 4: Commit**

```bash
git add go.mod go.sum
git commit -m "chore: initialize Go module and add dependencies"
```

---

## Task 2: Create Directory Structure

**Files:**
- Create: Directory structure as per design

**Step 1: Create directory structure**

Run:
```bash
mkdir -p cmd/vcli
mkdir -p internal/cli/credentials
mkdir -p internal/cli/snapshot
mkdir -p internal/cli/clone
mkdir -p internal/cli/inspect
mkdir -p pkg/config
mkdir -p pkg/vsphere
mkdir -p pkg/output
```

Expected: All directories created

**Step 2: Verify structure**

Run: `tree -L 3 -d`

Expected: Directory structure matches design

**Step 3: Commit**

```bash
git add .
git commit -m "chore: create project directory structure"
```

---

## Task 3: Config Package - Environment Variable Handling

**Files:**
- Create: `pkg/config/config.go`

**Step 1: Create config struct**

```go
package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds the vSphere connection configuration
type Config struct {
	Host     string
	Username string
	Password string
	Insecure bool
}

// Source tracks where a config value came from
type Source string

const (
	SourceEnv     Source = "env"
	SourceFlag    Source = "flag"
	SourceDefault Source = "default"
)

// ConfigWithSource holds config values with their sources
type ConfigWithSource struct {
	Host     string
	HostSrc  Source
	Username string
	UserSrc  Source
	Password string
	PassSrc  Source
	Insecure bool
	InsecSrc Source
}

// LoadFromEnv loads configuration from environment variables
func LoadFromEnv() (*Config, error) {
	cfg := &Config{
		Host:     os.Getenv("VCLI_HOST"),
		Username: os.Getenv("VCLI_USERNAME"),
		Password: os.Getenv("VCLI_PASSWORD"),
		Insecure: false,
	}

	if insecure := os.Getenv("VCLI_INSECURE"); insecure != "" {
		val, err := strconv.ParseBool(insecure)
		if err != nil {
			return nil, fmt.Errorf("invalid VCLI_INSECURE value: %w", err)
		}
		cfg.Insecure = val
	}

	return cfg, nil
}

// Validate checks if required configuration is present
func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("VCLI_HOST is required")
	}
	if c.Username == "" {
		return fmt.Errorf("VCLI_USERNAME is required")
	}
	if c.Password == "" {
		return fmt.Errorf("VCLI_PASSWORD is required")
	}
	return nil
}

// MaskPassword returns password with only first 2 and last 2 chars visible
func MaskPassword(password string) string {
	if len(password) <= 4 {
		return "••••"
	}
	return password[:2] + "••••••" + password[len(password)-2:]
}
```

**Step 2: Verify code compiles**

Run: `go build ./pkg/config`

Expected: No errors

**Step 3: Commit**

```bash
git add pkg/config/config.go
git commit -m "feat: add config package for environment variable handling"
```

---

## Task 4: vSphere Client Interface (Stub)

**Files:**
- Create: `pkg/vsphere/client.go`

**Step 1: Create client interface**

```go
package vsphere

import (
	"context"
	"errors"
)

// ErrNotImplemented is returned by stub methods
var ErrNotImplemented = errors.New("not implemented")

// Client interface for vSphere operations
type Client interface {
	// Connection
	Login(ctx context.Context) error
	Logout(ctx context.Context) error

	// Info
	Version(ctx context.Context) (string, error)
	CurrentUser(ctx context.Context) (string, error)
}

// StubClient is a stub implementation that returns not implemented errors
type StubClient struct{}

// NewStubClient creates a new stub client
func NewStubClient() *StubClient {
	return &StubClient{}
}

func (c *StubClient) Login(ctx context.Context) error {
	return ErrNotImplemented
}

func (c *StubClient) Logout(ctx context.Context) error {
	return ErrNotImplemented
}

func (c *StubClient) Version(ctx context.Context) (string, error) {
	return "", ErrNotImplemented
}

func (c *StubClient) CurrentUser(ctx context.Context) (string, error) {
	return "", ErrNotImplemented
}
```

**Step 2: Verify code compiles**

Run: `go build ./pkg/vsphere`

Expected: No errors

**Step 3: Commit**

```bash
git add pkg/vsphere/client.go
git commit -m "feat: add vSphere client interface stub"
```

---

## Task 5: Output Formatter Package

**Files:**
- Create: `pkg/output/formatter.go`

**Step 1: Create formatter interface**

```go
package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v3"
)

// Format represents output format
type Format string

const (
	FormatTable Format = "table"
	FormatJSON  Format = "json"
	FormatYAML  Format = "yaml"
)

// Formatter handles output formatting
type Formatter struct {
	format Format
	writer io.Writer
}

// NewFormatter creates a new formatter
func NewFormatter(format Format) *Formatter {
	return &Formatter{
		format: format,
		writer: os.Stdout,
	}
}

// SetWriter sets the output writer (useful for testing)
func (f *Formatter) SetWriter(w io.Writer) {
	f.writer = w
}

// PrintJSON outputs data as JSON
func (f *Formatter) PrintJSON(data interface{}) error {
	encoder := json.NewEncoder(f.writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// PrintYAML outputs data as YAML
func (f *Formatter) PrintYAML(data interface{}) error {
	encoder := yaml.NewEncoder(f.writer)
	defer encoder.Close()
	return encoder.Encode(data)
}

// PrintTable outputs data as a table
func (f *Formatter) PrintTable(headers []string, rows [][]string) {
	table := tablewriter.NewWriter(f.writer)
	table.SetHeader(headers)
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	for _, row := range rows {
		table.Append(row)
	}
	table.Render()
}

// Print outputs data based on the configured format
func (f *Formatter) Print(data interface{}, headers []string, rowFunc func(interface{}) [][]string) error {
	switch f.format {
	case FormatJSON:
		return f.PrintJSON(data)
	case FormatYAML:
		return f.PrintYAML(data)
	case FormatTable:
		if rowFunc != nil {
			rows := rowFunc(data)
			f.PrintTable(headers, rows)
		}
		return nil
	default:
		return fmt.Errorf("unknown format: %s", f.format)
	}
}
```

**Step 2: Verify code compiles**

Run: `go build ./pkg/output`

Expected: No errors

**Step 3: Commit**

```bash
git add pkg/output/formatter.go
git commit -m "feat: add output formatter package"
```

---

## Task 6: Root Command

**Files:**
- Create: `internal/cli/root.go`

**Step 1: Create root command**

```go
package cli

import (
	"fmt"
	"os"

	"github.com/asegev/vsphere-cli/pkg/config"
	"github.com/asegev/vsphere-cli/pkg/output"
	"github.com/spf13/cobra"
)

var (
	// Global flags
	flagHost     string
	flagUsername string
	flagPassword string
	flagInsecure bool
	flagOutput   string
	flagVerbose  bool

	// Global config
	globalConfig *config.Config
	globalFormat output.Format
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "vcli",
	Short: "A CLI tool for interacting with VMware vSphere/ESXi",
	Long: `vcli is a command-line tool for managing VMware vSphere environments.

It provides commands for snapshot management, VM cloning, VM inspection,
and credential validation.

Authentication is configured via environment variables:
  VCLI_HOST      - vCenter/ESXi host address
  VCLI_USERNAME  - Authentication username
  VCLI_PASSWORD  - Authentication password
  VCLI_INSECURE  - Skip TLS verification (optional, default: false)`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Skip validation for commands that don't need it
		if cmd.Name() == "help" || cmd.Name() == "vcli" {
			return nil
		}

		// Load config from environment
		cfg, err := config.LoadFromEnv()
		if err != nil {
			return err
		}

		// Override with flags if provided
		if flagHost != "" {
			cfg.Host = flagHost
		}
		if flagUsername != "" {
			cfg.Username = flagUsername
		}
		if flagPassword != "" {
			cfg.Password = flagPassword
		}
		if cmd.Flags().Changed("insecure") {
			cfg.Insecure = flagInsecure
		}

		// Validate config (skip for credentials show command)
		if cmd.Name() != "show" {
			if err := cfg.Validate(); err != nil {
				return fmt.Errorf("configuration error: %w\nSet environment variables or use flags", err)
			}
		}

		globalConfig = cfg

		// Set output format
		switch flagOutput {
		case "table", "json", "yaml":
			globalFormat = output.Format(flagOutput)
		default:
			return fmt.Errorf("invalid output format: %s (must be table, json, or yaml)", flagOutput)
		}

		return nil
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVar(&flagHost, "host", "", "vCenter/ESXi host (overrides VCLI_HOST)")
	rootCmd.PersistentFlags().StringVar(&flagUsername, "username", "", "Username (overrides VCLI_USERNAME)")
	rootCmd.PersistentFlags().StringVar(&flagPassword, "password", "", "Password (overrides VCLI_PASSWORD)")
	rootCmd.PersistentFlags().BoolVar(&flagInsecure, "insecure", false, "Skip TLS verification (overrides VCLI_INSECURE)")
	rootCmd.PersistentFlags().StringVarP(&flagOutput, "output", "o", "table", "Output format (table, json, yaml)")
	rootCmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "Verbose output")
}

// GetConfig returns the global config
func GetConfig() *config.Config {
	return globalConfig
}

// GetFormat returns the global output format
func GetFormat() output.Format {
	return globalFormat
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/root.go
git commit -m "feat: add root command with global flags"
```

---

## Task 7: Main Entry Point

**Files:**
- Create: `cmd/vcli/main.go`

**Step 1: Create main.go**

```go
package main

import (
	"os"

	"github.com/asegev/vsphere-cli/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
```

**Step 2: Build the binary**

Run: `go build -o vcli ./cmd/vcli`

Expected: Binary created successfully

**Step 3: Test running vcli**

Run: `./vcli --help`

Expected: Help text displays

**Step 4: Clean up binary**

Run: `rm vcli`

**Step 5: Commit**

```bash
git add cmd/vcli/main.go
git commit -m "feat: add main entry point"
```

---

## Task 8: Credentials Command Group

**Files:**
- Create: `internal/cli/credentials/credentials.go`

**Step 1: Create credentials command group**

```go
package credentials

import (
	"github.com/spf13/cobra"
)

// NewCredentialsCmd creates the credentials command
func NewCredentialsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "credentials",
		Short: "Validate and display credential configuration",
		Long: `The credentials command provides validation and visibility into
the current authentication configuration.

Available subcommands:
  test  - Test connection and validate credentials
  show  - Display current configuration (masked)`,
	}

	cmd.AddCommand(newTestCmd())
	cmd.AddCommand(newShowCmd())

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/credentials`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/credentials/credentials.go
git commit -m "feat: add credentials command group"
```

---

## Task 9: Credentials Test Command

**Files:**
- Create: `internal/cli/credentials/test.go`

**Step 1: Create test command**

```go
package credentials

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

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
			return errors.New("not implemented")
		},
	}

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/credentials`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/credentials/test.go
git commit -m "feat: add credentials test command stub"
```

---

## Task 10: Credentials Show Command

**Files:**
- Create: `internal/cli/credentials/show.go`

**Step 1: Create show command**

```go
package credentials

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func newShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Display current credential configuration",
		Long: `Displays the current credential configuration.

This command shows:
  - vSphere host
  - Username
  - Password (masked)
  - Insecure flag
  - Source of each value (environment variable or flag)

This does NOT test connectivity. Use 'credentials test' for that.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/credentials`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/credentials/show.go
git commit -m "feat: add credentials show command stub"
```

---

## Task 11: Snapshot Command Group

**Files:**
- Create: `internal/cli/snapshot/snapshot.go`

**Step 1: Create snapshot command group**

```go
package snapshot

import (
	"github.com/spf13/cobra"
)

// NewSnapshotCmd creates the snapshot command
func NewSnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snapshot",
		Short: "Manage VM snapshots",
		Long: `Manage VM snapshots with comprehensive snapshot management features.

Available subcommands:
  create       - Create a new snapshot
  list         - List snapshots for a VM
  tree         - Display snapshot hierarchy
  delete       - Delete a specific snapshot
  delete-all   - Delete all snapshots
  revert       - Revert to a snapshot
  consolidate  - Consolidate snapshot disks`,
	}

	cmd.AddCommand(newCreateCmd())
	cmd.AddCommand(newListCmd())
	cmd.AddCommand(newTreeCmd())
	cmd.AddCommand(newDeleteCmd())
	cmd.AddCommand(newDeleteAllCmd())
	cmd.AddCommand(newRevertCmd())
	cmd.AddCommand(newConsolidateCmd())

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/snapshot`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/snapshot/snapshot.go
git commit -m "feat: add snapshot command group"
```

---

## Task 12: Snapshot Create Command

**Files:**
- Create: `internal/cli/snapshot/create.go`

**Step 1: Create create command**

```go
package snapshot

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
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
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	cmd.Flags().StringVar(&createName, "name", "", "Snapshot name (auto-generated if omitted)")
	cmd.Flags().StringVar(&createDescription, "description", "", "Snapshot description")
	cmd.Flags().BoolVar(&createMemory, "memory", false, "Include VM memory state")
	cmd.Flags().BoolVar(&createQuiesce, "quiesce", false, "Quiesce filesystem (requires VMware Tools)")

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/snapshot`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/snapshot/create.go
git commit -m "feat: add snapshot create command stub"
```

---

## Task 13: Snapshot List Command

**Files:**
- Create: `internal/cli/snapshot/list.go`

**Step 1: Create list command**

```go
package snapshot

import (
	"errors"

	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <vm>",
		Short: "List VM snapshots",
		Long: `Lists all snapshots for a virtual machine in chronological order.

Output includes:
  - Snapshot name
  - Creation date
  - Description
  - Size
  - Current indicator

Examples:
  vcli snapshot list my-vm
  vcli snapshot list my-vm -o json`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/snapshot`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/snapshot/list.go
git commit -m "feat: add snapshot list command stub"
```

---

## Task 14: Snapshot Tree Command

**Files:**
- Create: `internal/cli/snapshot/tree.go`

**Step 1: Create tree command**

```go
package snapshot

import (
	"errors"

	"github.com/spf13/cobra"
)

func newTreeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tree <vm>",
		Short: "Display snapshot hierarchy as a tree",
		Long: `Displays the snapshot hierarchy as an ASCII tree.

Shows parent-child relationships and indicates the current snapshot.
Useful for VMs with complex snapshot trees.

Examples:
  vcli snapshot tree my-vm`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/snapshot`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/snapshot/tree.go
git commit -m "feat: add snapshot tree command stub"
```

---

## Task 15: Snapshot Delete Command

**Files:**
- Create: `internal/cli/snapshot/delete.go`

**Step 1: Create delete command**

```go
package snapshot

import (
	"errors"

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
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	cmd.Flags().BoolVar(&deleteForce, "force", false, "Skip confirmation prompt")

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/snapshot`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/snapshot/delete.go
git commit -m "feat: add snapshot delete command stub"
```

---

## Task 16: Snapshot Delete-All Command

**Files:**
- Create: `internal/cli/snapshot/delete_all.go`

**Step 1: Create delete-all command**

```go
package snapshot

import (
	"errors"

	"github.com/spf13/cobra"
)

var deleteAllConfirm bool

func newDeleteAllCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-all <vm>",
		Short: "Delete all snapshots from a VM",
		Long: `Removes all snapshots from a virtual machine.

Requires --confirm flag to prevent accidents.
Useful for cleanup operations.

Examples:
  vcli snapshot delete-all my-vm --confirm`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	cmd.Flags().BoolVar(&deleteAllConfirm, "confirm", false, "Confirm deletion of all snapshots (required)")
	cmd.MarkFlagRequired("confirm")

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/snapshot`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/snapshot/delete_all.go
git commit -m "feat: add snapshot delete-all command stub"
```

---

## Task 17: Snapshot Revert Command

**Files:**
- Create: `internal/cli/snapshot/revert.go`

**Step 1: Create revert command**

```go
package snapshot

import (
	"errors"

	"github.com/spf13/cobra"
)

var revertForce bool

func newRevertCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revert <vm> <snapshot-name>",
		Short: "Revert VM to a snapshot",
		Long: `Reverts a virtual machine to the specified snapshot state.

Prompts for confirmation unless --force flag is used.
Can revert to any snapshot in the tree, not just the current one.

Examples:
  vcli snapshot revert my-vm snapshot-2024-01-01
  vcli snapshot revert my-vm snapshot-2024-01-01 --force`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	cmd.Flags().BoolVar(&revertForce, "force", false, "Skip confirmation prompt")

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/snapshot`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/snapshot/revert.go
git commit -m "feat: add snapshot revert command stub"
```

---

## Task 18: Snapshot Consolidate Command

**Files:**
- Create: `internal/cli/snapshot/consolidate.go`

**Step 1: Create consolidate command**

```go
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
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/snapshot`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/snapshot/consolidate.go
git commit -m "feat: add snapshot consolidate command stub"
```

---

## Task 19: Clone Command Group

**Files:**
- Create: `internal/cli/clone/clone.go`

**Step 1: Create clone command group**

```go
package clone

import (
	"github.com/spf13/cobra"
)

// NewCloneCmd creates the clone command
func NewCloneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone virtual machines",
		Long: `Clone virtual machines with basic cloning capabilities.

Available subcommands:
  create  - Create a full clone of a VM
  list    - List cloned VMs`,
	}

	cmd.AddCommand(newCreateCmd())
	cmd.AddCommand(newListCmd())

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/clone`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/clone/clone.go
git commit -m "feat: add clone command group"
```

---

## Task 20: Clone Create Command

**Files:**
- Create: `internal/cli/clone/create.go`

**Step 1: Create create command**

```go
package clone

import (
	"errors"

	"github.com/spf13/cobra"
)

var createSnapshot string

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
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	cmd.Flags().StringVar(&createSnapshot, "snapshot", "", "Clone from specific snapshot")

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/clone`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/clone/create.go
git commit -m "feat: add clone create command stub"
```

---

## Task 21: Clone List Command

**Files:**
- Create: `internal/cli/clone/list.go`

**Step 1: Create list command**

```go
package clone

import (
	"errors"

	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List cloned VMs",
		Long: `Lists virtual machines that were created as clones.

Output includes:
  - Clone name
  - Source VM
  - Creation date
  - Power state

Examples:
  vcli clone list
  vcli clone list -o json`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("not implemented")
		},
	}

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/clone`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/clone/list.go
git commit -m "feat: add clone list command stub"
```

---

## Task 22: Inspect Command Group

**Files:**
- Create: `internal/cli/inspect/inspect.go`

**Step 1: Create inspect command group**

```go
package inspect

import (
	"github.com/spf13/cobra"
)

// NewInspectCmd creates the inspect command
func NewInspectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inspect",
		Short: "Inspect virtual machines and resources",
		Long: `Inspect virtual machines and display detailed information.

Available subcommands:
  vm  - Display comprehensive VM information`,
	}

	cmd.AddCommand(newVMCmd())

	return cmd
}
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/inspect`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/inspect/inspect.go
git commit -m "feat: add inspect command group"
```

---

## Task 23: Inspect VM Command

**Files:**
- Create: `internal/cli/inspect/vm.go`

**Step 1: Create vm command**

```go
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
```

**Step 2: Verify code compiles**

Run: `go build ./internal/cli/inspect`

Expected: No errors

**Step 3: Commit**

```bash
git add internal/cli/inspect/vm.go
git commit -m "feat: add inspect vm command stub"
```

---

## Task 24: Wire Up All Commands to Root

**Files:**
- Modify: `internal/cli/root.go`

**Step 1: Import command packages**

Add imports to `internal/cli/root.go` after the existing imports:

```go
import (
	// ... existing imports ...
	"github.com/asegev/vsphere-cli/internal/cli/clone"
	"github.com/asegev/vsphere-cli/internal/cli/credentials"
	"github.com/asegev/vsphere-cli/internal/cli/inspect"
	"github.com/asegev/vsphere-cli/internal/cli/snapshot"
)
```

**Step 2: Add commands to root in init()**

Add to the `init()` function in `internal/cli/root.go` after the flag definitions:

```go
func init() {
	// ... existing flag definitions ...

	// Add command groups
	rootCmd.AddCommand(credentials.NewCredentialsCmd())
	rootCmd.AddCommand(snapshot.NewSnapshotCmd())
	rootCmd.AddCommand(clone.NewCloneCmd())
	rootCmd.AddCommand(inspect.NewInspectCmd())
}
```

**Step 3: Verify code compiles**

Run: `go build ./internal/cli`

Expected: No errors

**Step 4: Build and test binary**

Run: `go build -o vcli ./cmd/vcli && ./vcli --help`

Expected: All commands visible in help output

**Step 5: Test each command group**

Run:
```bash
./vcli credentials --help
./vcli snapshot --help
./vcli clone --help
./vcli inspect --help
```

Expected: Help text displays for each command

**Step 6: Clean up binary**

Run: `rm vcli`

**Step 7: Commit**

```bash
git add internal/cli/root.go
git commit -m "feat: wire up all command groups to root command"
```

---

## Task 25: Add .gitignore

**Files:**
- Create: `.gitignore`

**Step 1: Create .gitignore**

```
# Binaries
vcli
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binaries
*.test

# Output of go coverage
*.out

# Dependency directories
vendor/

# Go workspace file
go.work

# IDE
.vscode/
.idea/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db
```

**Step 2: Commit**

```bash
git add .gitignore
git commit -m "chore: add .gitignore"
```

---

## Task 26: Add README Build Instructions

**Files:**
- Modify: `README.md`

**Step 1: Update README with build instructions**

Append to `README.md`:

```markdown

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
```

**Step 2: Commit**

```bash
git add README.md
git commit -m "docs: add build instructions and usage to README"
```

---

## Task 27: Verify Complete Build

**Files:**
- None (verification only)

**Step 1: Clean any previous builds**

Run: `go clean`

**Step 2: Run go mod tidy**

Run: `go mod tidy`

Expected: No changes needed

**Step 3: Build the binary**

Run: `go build -o vcli ./cmd/vcli`

Expected: Build succeeds with no errors

**Step 4: Test all commands show help**

Run:
```bash
./vcli --help
./vcli credentials --help
./vcli credentials test --help
./vcli credentials show --help
./vcli snapshot --help
./vcli snapshot create --help
./vcli snapshot list --help
./vcli snapshot tree --help
./vcli snapshot delete --help
./vcli snapshot delete-all --help
./vcli snapshot revert --help
./vcli snapshot consolidate --help
./vcli clone --help
./vcli clone create --help
./vcli clone list --help
./vcli inspect --help
./vcli inspect vm --help
```

Expected: All commands display help text

**Step 5: Test not-implemented errors**

Run:
```bash
export VCLI_HOST=test.local
export VCLI_USERNAME=test
export VCLI_PASSWORD=test
./vcli credentials test
./vcli snapshot list test-vm
./vcli clone create vm1 vm2
./vcli inspect vm test-vm
```

Expected: Each command returns "not implemented" error

**Step 6: Clean up**

Run:
```bash
rm vcli
unset VCLI_HOST VCLI_USERNAME VCLI_PASSWORD
```

**Step 7: Commit final verification**

```bash
git add -A
git commit -m "chore: verify complete CLI skeleton build" --allow-empty
```

---

## Completion

The vcli CLI skeleton is now complete with:

- ✅ Full directory structure
- ✅ Go module and dependencies configured
- ✅ Root command with global flags and config validation
- ✅ All 4 command groups (credentials, snapshot, clone, inspect)
- ✅ All 13 subcommands defined with proper help text
- ✅ Config package for environment variable handling
- ✅ vSphere client interface (stub)
- ✅ Output formatter supporting table/json/yaml
- ✅ Proper error handling structure
- ✅ Build and usage documentation

**Next Steps:**
- Implement vSphere client using govmomi
- Implement each command's RunE function
- Add integration tests
- Add error handling and validation
