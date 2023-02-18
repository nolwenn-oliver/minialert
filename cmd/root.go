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

		clientFlag, _ := cmd.Flags().GetBool("client")
		if clientFlag {
			fmt.Printf("Running in client mode \n")
		}
		serverFlag, _ := cmd.Flags().GetBool("server")
		portFlag, _ := cmd.Flags().GetString("port")
		if serverFlag {
			fmt.Printf("Running in server mode \n")
			server.LaunchServer(portFlag)
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolP("client", "c", false, "Running minialert in client mode.")
	rootCmd.PersistentFlags().BoolP("server", "s", false, "Running minialert in client mode.")
	rootCmd.Flags().String("port", server.DefaultPort, "Server port.")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
