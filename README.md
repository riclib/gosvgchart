# GoSVGChart

A simple, declarative SVG chart library for Go. This library allows you to create various chart types with a clean, chainable API or a simple markdown-like text format.

## Features

- Multiple ways to create charts:
  - Purely declarative Go API with method chaining
  - Simple markdown-like text format (great for LLM-generated charts)
  - Web server for dynamic chart generation
  - Goldmark extension for embedding charts in markdown documents
- No external dependencies for core functionality (uses only the Go standard library)
- Generates SVG output that can be used in web applications or saved to files
- Supports multiple chart types:
  - Line charts
  - Bar charts
  - Pie/Donut charts
  - Heatmap charts (GitHub-style activity heatmap)
- Customizable styling and options

## Installation

```bash
go get github.com/riclib/gosvgchart
```

## Usage (Go API)

### Line Chart

```go
import (
    "github.com/riclib/gosvgchart"
    "os"
)

func main() {
    // Create a line chart
    chart := gosvgchart.New().
        SetTitle("Monthly Sales").
        SetData([]float64{120, 250, 180, 310, 270, 390}).
        SetLabels([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}).
        SetColors([]string{"#3498db"}).
        SetSize(600, 400)

    // Render to SVG string
    svg := chart.Render()

    // Save to file
    os.WriteFile("line_chart.svg", []byte(svg), 0644)
}
```

### Bar Chart

```go
// Create a bar chart
chart := gosvgchart.New().
    SetTitle("Quarterly Revenue").
    SetData([]float64{850, 940, 1100, 1200}).
    SetLabels([]string{"Q1", "Q2", "Q3", "Q4"}).
    SetColors([]string{"#2ecc71", "#e74c3c", "#f39c12", "#9b59b6"}).
    SetSize(600, 400)

// Render to SVG string
svg := chart.Render()
```

### Pie Chart

```go
// Create a pie chart
chart := gosvgchart.New().
    SetTitle("Market Share").
    SetData([]float64{35, 25, 20, 15, 5}).
    SetLabels([]string{"Product A", "Product B", "Product C", "Product D", "Others"}).
    SetColors([]string{"#3498db", "#2ecc71", "#e74c3c", "#f39c12", "#9b59b6"}).
    SetSize(600, 500)

// For a donut chart, add:
chart.SetDonutHole(0.6) // 0-0.9, where 0 is a pie chart and 0.9 is a thin donut

// Render to SVG string
svg := chart.Render()
```

### Heatmap Chart (GitHub-style)

```go
import (
    "github.com/riclib/gosvgchart"
    "time"
)

// Create a heatmap chart
chart := gosvgchart.NewHeatmapChart().
    SetTitle("GitHub Contributions").
    SetSize(800, 200)

// Set activity data (dates must be in YYYY-MM-DD format)
chart.SetLabels([]string{
    "2025-01-01", "2025-01-05", "2025-01-10", 
    "2025-01-15", "2025-01-20", "2025-01-25",
    "2025-02-01", "2025-02-05", "2025-02-10",
    "2025-02-15",
}).
SetData([]float64{5, 12, 3, 15, 8, 4, 7, 14, 6, 11})

// Optional customization
chart.SetCellSize(15).          // Size of each cell
      SetCellSpacing(3).        // Space between cells
      SetCellRounding(2).       // Corner radius
      SetMaxValue(15).          // Maximum value for color scaling (0 for auto)
      SetDateFormat("2006-01-02") // Go time format string

// GitHub-style green gradient colors (light to dark)
chart.SetColors([]string{
    "#ebedf0", "#9be9a8", "#40c463", "#30a14e", "#216e39",
})

// Render to SVG string
svg := chart.Render()
```

## Markdown Chart Format

For an even simpler way to create charts, especially when working with LLMs or in text environments, you can use the markdown-like format:

### Line Chart Example

```gosvgchart
linechart
title: Monthly Sales
width: 600
height: 400
colors: #3498db, #e74c3c

data:
Jan | 120
Feb | 250
Mar | 180
Apr | 310
May | 270
Jun | 390
```

### Bar Chart Example

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

### Pie Chart Example

```gosvgchart
piechart
title: Market Share
width: 600
height: 500
colors: #3498db, #2ecc71, #e74c3c, #f39c12, #9b59b6

data:
Product A | 35
Product B | 25
Product C | 20
Product D | 15
Others | 5
```

### Heatmap Chart Example

```gosvgchart
heatmapchart
title: GitHub Contribution Activity
width: 800
height: 200
colors: #ebedf0, #9be9a8, #40c463, #30a14e, #216e39

data:
2025-01-01 | 5
2025-01-03 | 8
2025-01-10 | 3
2025-01-15 | 15
2025-01-24 | 10
2025-02-01 | 7
2025-02-05 | 14
2025-02-15 | 11
2025-02-24 | 10
```

### Side-by-Side Charts Example

You can place multiple charts side by side by using the `---` separator within a single code block:

```gosvgchart
barchart
title: 2023 Revenue
width: 450
height: 300
colors: #3498db, #2ecc71

data:
Q1 | 850
Q2 | 940
Q3 | 1100
Q4 | 1200

---

barchart
title: 2024 Revenue
width: 450
height: 300
colors: #e74c3c, #f39c12

data:
Q1 | 950
Q2 | 1040
Q3 | 1200
Q4 | 1400
```

## Command Line Tools

### Convert Markdown to SVG

```bash
go run cmd/mdchart/main.go -input chart.md -output chart.svg
```

### Run Web Server

```bash
go run cmd/mdchartserver/main.go -port 8080
```

Then visit:
- `http://localhost:8080` for the web UI
- `POST` to `http://localhost:8080/chart` with markdown content to get SVG
- `GET` to `http://localhost:8080/charturl?md=linechart_n_title:Test_n_data:_n_A_p_10_n_B_p_20` to get SVG directly via URL

## Goldmark Extension

GoSVGChart provides a [Goldmark](https://github.com/yuin/goldmark) extension that allows you to embed charts directly in your markdown documents:

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

// Convert markdown to HTML with embedded SVG charts
var output []byte
if err := markdown.Convert([]byte(markdownContent), &output); err != nil {
    // Handle error
}
```

Use the extension in your markdown with fenced code blocks:

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

For more details on the Goldmark extension, see the [goldmark directory](./goldmark/README.md).

## Common Methods (Go API)

All chart types share these common methods:

| Method | Description |
|--------|-------------|
| `SetTitle(title string)` | Sets the chart title |
| `SetSize(width, height int)` | Sets the chart dimensions in pixels |
| `SetData(data []float64)` | Sets the chart data values |
| `SetLabels(labels []string)` | Sets the chart labels |
| `SetColors(colors []string)` | Sets the color palette as hex values (e.g., "#ff0000") |
| `Render()` | Renders the chart to an SVG string |

## Chart-Specific Methods (Go API)

### Line Chart

| Method | Description |
|--------|-------------|
| `ShowDataPoints(show bool)` | Shows or hides data points |
| `SetSmooth(smooth bool)` | Enables smooth curved lines (when true) |

### Bar Chart

| Method | Description |
|--------|-------------|
| `SetHorizontal(horizontal bool)` | Displays bars horizontally (when true) |
| `SetStacked(stacked bool)` | Stacks multiple data series (when true) |

### Pie Chart

| Method | Description |
|--------|-------------|
| `SetDonutHole(percentage float64)` | Sets the inner circle size (0-0.9) |

### Heatmap Chart

| Method | Description |
|--------|-------------|
| `SetCellSize(size int)` | Sets the size of each cell in pixels |
| `SetCellSpacing(spacing int)` | Sets the spacing between cells |
| `SetCellRounding(radius int)` | Sets the corner radius of cells |
| `SetDateFormat(format string)` | Sets the date format (Go time format) |
| `SetMaxValue(max float64)` | Sets the maximum value for color scaling |
| `SetDayLabels(labels []string)` | Sets the labels for days of the week |
| `SetMonthLabels(labels []string)` | Sets the labels for months |

## Design Philosophy

GoSVGChart was designed with these principles in mind:

1. **Simplicity**: The API should be intuitive and require minimal code
2. **Zero Dependencies**: Only use the Go standard library
3. **Declarative Syntax**: Use method chaining for a clean, readable configuration
4. **LLM-Friendly**: Easy to explain to and be used by Large Language Models
5. **Multiple Interfaces**: Support both code and text-based chart definitions

## License

MIT