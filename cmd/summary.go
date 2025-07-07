package cmd

import (
	"gotrack/tracker"

	"github.com/spf13/cobra"
)

var summaryCommand = &cobra.Command{
	Use:   "summary",
	Short: "Show productivity summary",
	Run: func(cmd *cobra.Command, args []string) {

		tracker.ShowSummary()
	},
}

func init() {

	rootCommand.AddCommand(summaryCommand)
}
