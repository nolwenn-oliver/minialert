package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:     "send",
	Aliases: []string{"s"},
	Short:   "Send metric.",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		metric := args[0]
		value := args[1]

		fmt.Printf("Received metric '%s' value %s.\n", metric, value)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
