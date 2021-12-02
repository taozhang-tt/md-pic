package remove

import (
	"mdpic/cmd"
	"mdpic/manage"

	"github.com/spf13/cobra"
)

var (
	keys  = []string{}
	rmCmd = &cobra.Command{
		Use:   "rm",
		Short: "Delete object",
		Args: func(cmd *cobra.Command, args []string) error {
			keys = args
			return nil
		},
		Run: runRemove,
	}
)

func init() {
	cmd.RootCmd.AddCommand(rmCmd)

}

func runRemove(cmd *cobra.Command, args []string) {
	manage.DeleteObject(keys)
}
