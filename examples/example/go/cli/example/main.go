package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	greetFormatterCmd.AddCommand(greetFormatterGreetCmd)
	rootCmd.AddCommand(greetFormatterCmd)

	greeterCmd.AddCommand(greeterGreetCmd)
	greeterCmd.AddCommand(greeterGreetPhotoCmd)
	rootCmd.AddCommand(greeterCmd)

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
}

// greetFormatterCmd is the service command for
// information.
var greetFormatterCmd = &cobra.Command{
	Use:   "example GreetFormatter method [data]",
	Short: "GreetFormatter provides formattable greeting services.",
	Long:  `GreetFormatter provides formattable greeting services.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// greetFormatterGreetCmd is the method.
var greetFormatterGreetCmd = &cobra.Command{
	Use:   "example GreetFormatter Greet [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// greeterCmd is the service command for
// information.
var greeterCmd = &cobra.Command{
	Use:   "example Greeter method [data]",
	Short: "Greeter provides greeting services.",
	Long:  `Greeter provides greeting services.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// greeterGreetCmd is the method.
var greeterGreetCmd = &cobra.Command{
	Use:   "example Greeter Greet [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// greeterGreetPhotoCmd is the method.
var greeterGreetPhotoCmd = &cobra.Command{
	Use:   "example Greeter GreetPhoto [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
