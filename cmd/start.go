package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCommand = &cobra.Command{

	Use:   "start [task]",
	Short: "Start a new task",

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {

			fmt.Println("Task Name required")
			return
		}

	},
}

func init() {
	rootCommand.AddCommand(startCommand)
}
