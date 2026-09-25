package cmd

import (
	"path/filepath"

	"github.com/devmoded/mp-cli/internal/modpack"
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
		out.Println(output.Message{Event: output.EventError, Message: err.Error()})
		return
	}
	a, err := modpack.Open(path)
	if err != nil {
		out.Println(output.Message{Event: output.EventError, Message: err.Error()})
		return
	}

	i, err := modpack.Inspect(a)
	if err != nil {
		out.Println(output.Message{Event: output.EventError, Message: err.Error()})
		return
	}
	out.Println(output.Message{Event: output.EventInspect, Message: i})
}
