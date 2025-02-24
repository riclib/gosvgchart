package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/riclib/gosvgchart/mdparser"
)

func main() {
	// Define command line flags
	inputFile := flag.String("input", "", "Input markdown file containing chart specification")
	outputFile := flag.String("output", "", "Output SVG file (default: derived from input filename)")
	flag.Parse()

	// Check if input file is provided
	if *inputFile == "" {
		fmt.Println("Error: Input file is required")
		flag.Usage()
		os.Exit(1)
	}

	// Read input file
	markdownBytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// Parse markdown and generate SVG
	svg, err := mdparser.ParseMarkdownChart(string(markdownBytes))
	if err != nil {
		fmt.Printf("Error parsing chart: %v\n", err)
		os.Exit(1)
	}

	// Determine output file name if not provided
	if *outputFile == "" {
		baseName := strings.TrimSuffix(filepath.Base(*inputFile), filepath.Ext(*inputFile))
		*outputFile = baseName + ".svg"
	}

	// Write SVG to output file
	err = ioutil.WriteFile(*outputFile, []byte(svg), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Chart successfully generated: %s\n", *outputFile)
}