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
	table := tablewriter.NewTable(f.writer,
		tablewriter.WithHeader(headers),
	)

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
