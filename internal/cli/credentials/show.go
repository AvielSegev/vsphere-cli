package credentials

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func newShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Display current credential configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			// 1. Fetch values
			host := os.Getenv("VCLI_HOST")
			user := os.Getenv("VCLI_USERNAME")
			pass := os.Getenv("VCLI_PASSWORD")

			// 2. Initialize TabWriter
			// minwidth, tabwidth, padding, padchar, flags
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

			fmt.Fprintln(w, "KEY\tVALUE\tSOURCE")
			fmt.Fprintln(w, "---\t-----\t------")

			// Helper to print rows with masked password logic
			printRow := func(w *tabwriter.Writer, key, val, envName string, mask bool) {
				source := "Manual/Flag"
				if val == os.Getenv(envName) && val != "" {
					source = fmt.Sprintf("Env (%s)", envName)
				}

				displayVal := val
				if mask && val != "" {
					displayVal = "********"
				} else if val == "" {
					displayVal = "<not set>"
				}

				fmt.Fprintf(w, "%s\t%s\t%s\n", key, displayVal, source)
			}

			// 3. Print the data
			printRow(w, "Host", host, "VCLI_HOST", false)
			printRow(w, "Username", user, "VCLI_USERNAME", false)
			printRow(w, "Password", pass, "VCLI_PASSWORD", true)

			return w.Flush()
		},
	}

	return cmd
}
