package cmd

import (
	"github.com/spf13/cobra"
	"minialert/client"
)

var listCmd = &cobra.Command{
	Use:     "alerts-list",
	Aliases: []string{"als"},
	Short:   "List alerts history.",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		portFlag, _ := cmd.Flags().GetString("port")
		client.GetAlertHistory(portFlag)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
