package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/riclib/gosvgchart/goldmark"
	gm "github.com/yuin/goldmark"
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
# Sample Charts

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

## Side-by-Side Charts

When two or more charts are placed one after another without any text in between, they will be displayed side by side:

` + "```gosvgchart" + `
linechart
title: Monthly Sales
width: 500
height: 350
colors: #3498db, #e74c3c

data:
Jan | 120
Feb | 250
Mar | 180
Apr | 310
May | 270
Jun | 390
` + "```" + `
` + "```gosvgchart" + `
barchart
title: Quarterly Revenue
width: 500
height: 350
colors: #2ecc71, #e74c3c, #f39c12, #9b59b6

data:
Q1 | 850
Q2 | 940
Q3 | 1100
Q4 | 1200
` + "```" + `

## Three Charts in a Row

You can also place three or more charts side by side:

` + "```gosvgchart" + `
piechart
title: Market Share A
width: 300
height: 300
colors: #3498db, #e74c3c, #2ecc71

data:
Product X | 45
Product Y | 35
Product Z | 20
` + "```" + `
` + "```gosvgchart" + `
piechart
title: Market Share B
width: 300
height: 300
colors: #f39c12, #9b59b6, #e67e22

data:
Product X | 25
Product Y | 55
Product Z | 20
` + "```" + `
` + "```gosvgchart" + `
piechart
title: Market Share C
width: 300
height: 300
colors: #1abc9c, #3498db, #e74c3c

data:
Product X | 30
Product Y | 20
Product Z | 50
` + "```" + `
`

	// Convert markdown to HTML
	var buf bytes.Buffer
	if err := markdown.Convert([]byte(markdownContent), &buf); err != nil {
		fmt.Println("Error converting markdown:", err)
		os.Exit(1)
	}

	// Save the output to an HTML file
	if err := os.WriteFile("output.html", buf.Bytes(), 0644); err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully converted markdown to HTML. Check output.html")
}
