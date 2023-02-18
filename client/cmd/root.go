package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "minialert",
	Short: "Mini alerting system",
	Long:  `Mini alerting system allowing to send metric and trigger alerts based on business rules.`,
}

func Execute() {
	rootCmd.PersistentFlags().BoolP("client", "c", false, "Running minialert in client mode.")
	rootCmd.PersistentFlags().BoolP("server", "s", false, "Running minialert in client mode.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
