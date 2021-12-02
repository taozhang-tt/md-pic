package cmd

import (
	"mdpic/manage"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "mdpic",
		Short: "Oss management util",
		Run:   f,
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func f(cmd *cobra.Command, args []string) {
	manage.UploadFromClipboard()
}
