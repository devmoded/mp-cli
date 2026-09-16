package cmd

import (
	"path/filepath"

	"github.com/devmoded/mp-cli/internal/archive"
	"github.com/devmoded/mp-cli/internal/output"
	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:   "extract [modpack-path] [destination]",
	Short: "Extract files from a modpack",
	Example: `# Extract all files
mp-cli extract modpack.zip .

# Extract specific entry
mp-cli extract modpack.zip . --entry info.json --entry mods
`,
	Args: cobra.ExactArgs(2),
	Run:  extract,
}

func extract(cmd *cobra.Command, args []string) {
	path, err := filepath.Abs(args[0])
	if err != nil {
		out.Println(output.Message{Event: output.EventError, Message: err.Error()})
		return
	}
	dest, err := filepath.Abs(args[0])
	if err != nil {
		out.Println(output.Message{Event: output.EventError, Message: err.Error()})
		return
	}

	a, err := archive.Open(path)
	if err != nil {
		out.Println(output.Message{Event: output.EventError, Message: err.Error()})
		return
	}

	if extractEntries == nil {
		err := a.ExtractAll(dest)
		if err != nil {
			out.Println(output.Message{Event: output.EventError, Message: err.Error()})
			return
		}
	}
}
