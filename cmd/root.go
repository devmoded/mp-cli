package cmd

import (
	"fmt"
	"os"

	"github.com/devmoded/mp-cli/internal/download"
	"github.com/devmoded/mp-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	formatJSON bool
	filename   string
)

var rootCmd = &cobra.Command{
	Use:     "mp-cli [command]",
	Short:   "Utility for downloading and unpacking Minecraft modpacks",
	Example: "mp-cli download github.com/user/repo/ .",
}

var downloadCmd = &cobra.Command{
	Use:   "download [modpack-url] [destination]",
	Short: "Downloads a modpack to the specified directory",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		o := output.NewOutput(formatJSON)
		download.Download(o, args[0], args[1], filename)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&formatJSON, "json", false, "Print messages in JSON format")

	downloadCmd.Flags().StringVarP(&filename, "filename", "f", "", "Override filename")

	rootCmd.AddCommand(downloadCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
