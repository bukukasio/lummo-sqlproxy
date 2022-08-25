package cmd

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "lummo-sqlproxy [sub]",
	Version: version,
	Short:   "CloudSQL Proxy CLI",
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to cloudsql instance",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		_, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err != nil {
			fmt.Printf("Port already in use\n")
			os.Exit(1)
		}
		connectInstance(port)
	},
}

var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "disconnect cloudsql instance proxy",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		disconnectInstance()
	},
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "for troubleshooting",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		doctor()
	},
}

func Execute() {
	err := connectCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(disconnectCmd, connectCmd, doctorCmd)
	connectCmd.PersistentFlags().IntP("port", "p", 5432, "port")
}
