package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "remoto",
	Short: "Remoto is a complete RPC solution with a very simple design.",
	Long:  `Remoto is a complete RPC solution with a very simple design.`,
	Run:   func(cmd *cobra.Command, args []string) {},
	Args:  cobra.MinimumNArgs(1),
}
