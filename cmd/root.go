package cmd

import (
	"mdpic/upload"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "mdpic",
		Short: "A upload util",
		Run:   f,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func f(cmd *cobra.Command, args []string) {
	upload.UploadFromClipboard()
}
