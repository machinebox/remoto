package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// this tool copies the remototypes/remototypes.go file into a local variable
// so that it can still be used in situations where the remototypes package
// is not available.

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() error {
	out, err := os.Create("copy_of_remototypes.go")
	if err != nil {
		return err
	}
	defer out.Close()
	b, err := ioutil.ReadFile("../remototypes/remototypes.go")
	if err != nil {
		return err
	}
	io.WriteString(out, "package generator\n\n")
	io.WriteString(out, "// remotoTypesSrc is a copy of remototypes/remototypes.go\n")
	fmt.Fprintf(out, "var remotoTypesSrc = %q\n", string(b))
	return nil
}
