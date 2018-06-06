package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	suggestionboxCmd.AddCommand(suggestionboxCreateModelCmd)
	suggestionboxCmd.AddCommand(suggestionboxDeleteModelCmd)
	suggestionboxCmd.AddCommand(suggestionboxGetStateCmd)
	suggestionboxCmd.AddCommand(suggestionboxListModelsCmd)
	suggestionboxCmd.AddCommand(suggestionboxPredictCmd)
	suggestionboxCmd.AddCommand(suggestionboxPutStateCmd)
	suggestionboxCmd.AddCommand(suggestionboxRewardCmd)
	rootCmd.AddCommand(suggestionboxCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "suggestionbox service method [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
	Args: cobra.MinimumNArgs(2),
}

// suggestionboxCmd is the service command for
// information.
var suggestionboxCmd = &cobra.Command{
	Use:   "suggestionbox Suggestionbox method [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// suggestionboxCreateModelCmd is the method.
var suggestionboxCreateModelCmd = &cobra.Command{
	Use:   "suggestionbox Suggestionbox CreateModel [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// suggestionboxDeleteModelCmd is the method.
var suggestionboxDeleteModelCmd = &cobra.Command{
	Use:   "suggestionbox Suggestionbox DeleteModel [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// suggestionboxGetStateCmd is the method.
var suggestionboxGetStateCmd = &cobra.Command{
	Use:   "suggestionbox Suggestionbox GetState [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// suggestionboxListModelsCmd is the method.
var suggestionboxListModelsCmd = &cobra.Command{
	Use:   "suggestionbox Suggestionbox ListModels [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// suggestionboxPredictCmd is the method.
var suggestionboxPredictCmd = &cobra.Command{
	Use:   "suggestionbox Suggestionbox Predict [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// suggestionboxPutStateCmd is the method.
var suggestionboxPutStateCmd = &cobra.Command{
	Use:   "suggestionbox Suggestionbox PutState [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// suggestionboxRewardCmd is the method.
var suggestionboxRewardCmd = &cobra.Command{
	Use:   "suggestionbox Suggestionbox Reward [data]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
