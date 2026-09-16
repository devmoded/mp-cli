package cmd

import (
	"github.com/devmoded/mp-cli/internal/download"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download [modpack-url] [destination]",
	Short: "Download a modpack to the specified directory",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		download.Download(out, args[0], args[1], downloadFilename)
	},
}
