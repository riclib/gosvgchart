# GoSVGChart Goldmark Extension

This extension for the [Goldmark](https://github.com/yuin/goldmark) markdown parser allows you to embed charts directly in your markdown documents using the familiar GoSVGChart syntax.

## Installation

```bash
go get github.com/riclib/gosvgchart
```

## Usage

Add the GoSVGChart extension to your Goldmark parser:

```go
import (
    "github.com/riclib/gosvgchart/goldmark"
    gm "github.com/yuin/goldmark"
)

// Create a new Goldmark instance with the gosvgchart extension
markdown := gm.New(
    gm.WithExtensions(
        goldmark.New(),
    ),
)

// Convert markdown to HTML
var output []byte
if err := markdown.Convert([]byte(markdownContent), &output); err != nil {
    // Handle error
}
```

## Syntax

Use a fenced code block with the language identifier `gosvgchart` and include the chart definition:

````markdown
```gosvgchart
barchart
title: Quarterly Revenue
width: 600
height: 400
colors: #2ecc71, #e74c3c, #f39c12, #9b59b6

data:
Q1 | 850
Q2 | 940
Q3 | 1100
Q4 | 1200
```
````

The extension will replace this code block with the rendered SVG chart.

## Chart Types

The extension supports all chart types provided by GoSVGChart:

1. **Line Chart**: Use `linechart` as the first line in your code block
2. **Bar Chart**: Use `barchart` as the first line in your code block
3. **Pie Chart**: Use `piechart` as the first line in your code block

## Complete Example

Here's a full example of a markdown document with various chart types:

````markdown
# Sales Report

## Quarterly Revenue

```gosvgchart
barchart
title: Quarterly Revenue
width: 600
height: 400
colors: #2ecc71, #e74c3c, #f39c12, #9b59b6

data:
Q1 | 850
Q2 | 940
Q3 | 1100
Q4 | 1200
```

## Monthly Sales Trend

```gosvgchart
linechart
title: Monthly Sales
width: 600
height: 400
colors: #3498db

data:
Jan | 120
Feb | 250
Mar | 180
Apr | 310
May | 270
Jun | 390
```

## Market Share

```gosvgchart
piechart
title: Market Share
width: 500
height: 500
colors: #3498db, #2ecc71, #e74c3c, #f39c12, #9b59b6

data:
Product A | 35
Product B | 25
Product C | 20
Product D | 15
Others | 5
```
````

## Side-by-Side Charts

When you place multiple chart code blocks directly next to each other without any text in between, the extension will automatically render them side-by-side in a responsive layout:

````markdown
## Revenue Comparison

```gosvgchart
barchart
title: 2023 Revenue
width: 500
height: 300
colors: #3498db, #2ecc71

data:
Q1 | 850
Q2 | 940
Q3 | 1100
Q4 | 1200
```
```gosvgchart
barchart
title: 2024 Revenue
width: 500
height: 300
colors: #e74c3c, #f39c12

data:
Q1 | 950
Q2 | 1040
Q3 | 1200
Q4 | 1400
```
````

You can also place three or more charts side by side, and they will automatically wrap to new lines on smaller screens:

````markdown
## Regional Market Share

```gosvgchart
piechart
title: North Region
width: 300
height: 300
colors: #3498db, #2ecc71, #e74c3c

data:
Product A | 45
Product B | 35
Product C | 20
```
```gosvgchart
piechart
title: South Region
width: 300
height: 300
colors: #3498db, #2ecc71, #e74c3c

data:
Product A | 30
Product B | 45
Product C | 25
```
```gosvgchart
piechart
title: East Region
width: 300
height: 300
colors: #3498db, #2ecc71, #e74c3c

data:
Product A | 25
Product B | 25
Product C | 50
```
````

## Error Handling

If there's an error in your chart definition, the extension will output an HTML comment with the error message instead of the SVG chart.

## License

MIT