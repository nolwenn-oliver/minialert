package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"minialert/server"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "minialert",
	Short: "Mini alerting system",
	Long:  `Mini alerting system allowing to send metric and trigger alerts based on business rules.`,
	Run: func(cmd *cobra.Command, args []string) {

		serverFlag, _ := cmd.Flags().GetBool("server")
		portFlag, _ := cmd.Flags().GetString("port")
		if serverFlag {
			server.LaunchServer(portFlag)
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolP("client", "c", true, "Running minialert in client mode")
	rootCmd.PersistentFlags().BoolP("server", "s", false, "Running minialert in client mode")
	rootCmd.PersistentFlags().StringP("port", "p", server.DefaultPort, "Server port")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
