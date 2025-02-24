package main

import (
	"fmt"
	"github.com/riclib/gosvgchart/goldmark"
	gm "github.com/yuin/goldmark"
	"os"
)

func main() {
	// Create a new Goldmark instance with the gosvgchart extension
	markdown := gm.New(
		gm.WithExtensions(
			goldmark.New(),
		),
	)

	// Example markdown content with a gosvgchart code block
	markdownContent := `
# Sample Chart

This is a sample chart rendered with gosvgchart:

` + "```gosvgchart" + `
barchart
title: Sample Bar Chart
width: 600
height: 400
colors: #3498db, #e74c3c, #2ecc71, #f39c12

data:
Q1 | 120
Q2 | 250
Q3 | 180
Q4 | 310
` + "```" + `

And here's a pie chart:

` + "```gosvgchart" + `
piechart
title: Sample Pie Chart
width: 500
height: 400
colors: #3498db, #e74c3c, #2ecc71, #f39c12

data:
Product A | 35
Product B | 25
Product C | 20
Product D | 15
` + "```" + `
`

	// Convert markdown to HTML
	var output []byte
	if err := markdown.Convert([]byte(markdownContent), &output); err != nil {
		fmt.Println("Error converting markdown:", err)
		os.Exit(1)
	}

	// Save the output to an HTML file
	if err := os.WriteFile("output.html", output, 0644); err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully converted markdown to HTML. Check output.html")
}