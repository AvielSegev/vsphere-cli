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
