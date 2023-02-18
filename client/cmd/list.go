package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "alerts list",
	Aliases: []string{"als"},
	Short:   "List alerts history.",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Listing alerts.....\n")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
