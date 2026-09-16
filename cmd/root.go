package cmd

import (
	"fmt"
	"os"

	"github.com/devmoded/mp-cli/internal/output"
	"github.com/devmoded/mp-cli/internal/version"
	"github.com/spf13/cobra"
)

var out output.Output

var (
	formatJSON       bool
	downloadFilename string
	extractEntries   []string
)

var rootCmd = &cobra.Command{
	Use:     "mp-cli [command]",
	Short:   "Utility for downloading and unpacking Minecraft modpacks",
	Example: "mp-cli download github.com/user/repo/releases/latest/download/modpack.zip .",
	Version: version.Version,
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&formatJSON, "json", false, "Print messages in JSON format")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		out = output.NewOutput(formatJSON)
	}

	downloadCmd.Flags().StringVarP(&downloadFilename, "filename", "f", "", "Override filename")
	extractCmd.Flags().StringArrayVarP(&extractEntries, "entry", "e", nil, "Extract only the specified entries")

	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(inspectCmd)
	rootCmd.AddCommand(extractCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
