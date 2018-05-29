package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	faceboxCmd.AddCommand(faceboxCheckCmd)
	faceboxCmd.AddCommand(faceboxTeachCmd)
	rootCmd.AddCommand(faceboxCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "example service method [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
	Args: cobra.MinimumNArgs(2),
}

// faceboxCmd is the service command for
// information.
var faceboxCmd = &cobra.Command{
	Use:   "example Facebox method [data]",
	Short: "Facebox provides facial detection and recognition capabilities.",
	Long:  `Facebox provides facial detection and recognition capabilities.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// faceboxCheckCmd is the method.
var faceboxCheckCmd = &cobra.Command{
	Use:   "example Facebox Check [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// faceboxTeachCmd is the method.
var faceboxTeachCmd = &cobra.Command{
	Use:   "example Facebox Teach [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
