// Description: a simple CLI utility to completely strip HTML tags from the input text.
// Makes use of the bluemonday package for sanitization and supports the following flags:
// -i: path to the input file
// -o: path to the output file
//
// Also allows for reading from stdin and writing to stdout if no flags are provided.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/microcosm-cc/bluemonday"
)

var (
	inputFile  = flag.String("i", "", "path to the input file")
	outputFile = flag.String("o", "", "path to the output file")
)

func main() {
	flag.Parse()

	var (
		input  *os.File = os.Stdin
		output *os.File = os.Stdout
	)

	if *inputFile != "" {
		input, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer input.Close()
	}

	if *outputFile != "" {
		output, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer output.Close()
	}

	p := bluemonday.StrictPolicy()
	if _, err := io.Copy(output, p.SanitizeReader(input)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
