package cmd

import (
	"gotrack/tracker"

	"github.com/spf13/cobra"
)

var logCommand = &cobra.Command{

	Use:   "log",
	Short: "Show today's task log",
	Run: func(cmd *cobra.Command, args []string) {

		tracker.ShowLog()
	},
}

func init() {

	rootCommand.AddCommand(logCommand)
}
