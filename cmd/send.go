package cmd

import (
	"github.com/spf13/cobra"
	"minialert/client"
)

var sendCmd = &cobra.Command{
	Use:     "send",
	Aliases: []string{"s"},
	Short:   "Send metric (Usage <metric_name> <value>)",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]
		value := args[1]

		portFlag, _ := cmd.Flags().GetString("port")
		client.SendDataMetric(portFlag, name, value)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
