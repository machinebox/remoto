package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/machinebox/remoto/generator"
	"github.com/spf13/cobra"
)

func init() {
	var outputFile string
	var generateCmd = &cobra.Command{
		Use:   "generate definition template",
		Short: "Generate source code from a template and remoto definition.",
		Args:  cobra.ExactArgs(2),

		Run: func(cmd *cobra.Command, args []string) {
			definition := args[0]
			template := args[1]
			f, err := os.Open(definition)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			defer f.Close()
			def, err := generator.Parse(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "parse: %v\n", err)
				os.Exit(1)
			}
			var o io.Writer = os.Stdout
			if outputFile != "" {
				if err := os.MkdirAll(filepath.Dir(outputFile), 0777); err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
					os.Exit(1)
				}
				outFile, err := os.Create(outputFile)
				if err != nil {
					fmt.Fprintf(os.Stderr, "parse: %v\n", err)
					os.Exit(1)
				}
				defer outFile.Close()
				o = outFile
			}
			b, err := ioutil.ReadFile(template)
			if err != nil {
				fmt.Fprintf(os.Stderr, "template: %v\n", err)
				os.Exit(1)
			}
			if err := generator.Render(o, template, string(b), def); err != nil {
				fmt.Fprintf(os.Stderr, "render template: %v\n", err)
				os.Exit(1)
			}
		},
	}
	generateCmd.Flags().StringVarP(&outputFile, "output", "o", "", "output file (default stdout)")
	rootCmd.AddCommand(generateCmd)
}
