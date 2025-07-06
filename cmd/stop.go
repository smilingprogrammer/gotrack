package cmd

import (
	"gotrack/tracker"

	"github.com/spf13/cobra"
)

var stopCommand = &cobra.Command{

	Use:   "stop [task]",
	Short: "Stop the current task",
	Run: func(cmd *cobra.Command, args []string) {

		tracker.StopTask()
	},
}

func init() {

	rootCommand.AddCommand(stopCommand)
}
