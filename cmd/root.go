package cmd

import "github.com/spf13/cobra"

var rootCommand = &cobra.Command{

	Use:   "go track",
	Short: "CLI Time Tracker",
}

func Execute() {

	cobra.CheckErr(rootCommand.Execute())
}
