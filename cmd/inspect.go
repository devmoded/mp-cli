package cmd

import (
	"path/filepath"

	"github.com/devmoded/mp-cli/internal/archive"
	"github.com/devmoded/mp-cli/internal/output"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect [modpack-path]",
	Short: "Inspect a modpack",
	Args:  cobra.ExactArgs(1),
	Run:   inspect,
}

func inspect(cmd *cobra.Command, args []string) {
	path, err := filepath.Abs(args[0])
	if err != nil {
		out.Println(output.Message{Event: "error", Message: err.Error()})
	}
	a := archive.Open(path)

	meta, entries, err := a.Inspect()
	if err != nil {
		out.Println(output.Message{Event: "error", Message: err.Error()})
	}
	out.Println(output.Message{Event: "message", Message: meta.Sprint()})
	for _, entry := range entries {
		out.Println(output.Message{Event: "message", Message: entry.Path})
	}
}
