/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dbcon [sub]",
	Short: "CloudSQL Proxy CLI",
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to cloudsql instance",
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		port, _ := cmd.Flags().GetInt("port")
		fmt.Printf("Environment: %v\n", env)
		_, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err != nil {
			// Log or report the error here
			fmt.Printf("Port already in use\n")
			os.Exit(1)
		}
		connectInstance(env, port)
	},
}

var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "disconnect cloudsql instance proxy",
	// Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		disconnectInstance()
	},
}

func Execute() {
	err := connectCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(disconnectCmd, connectCmd)
	connectCmd.PersistentFlags().String("env", "dev", "environment\nSupported environments: dev, staging, prod\n")
	connectCmd.PersistentFlags().Int("port", 5432, "port")
}
