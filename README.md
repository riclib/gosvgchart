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
- Automatic dark mode support for system color scheme adaptation

## Dark Mode Support

SVGs created with GoSVGChart automatically adapt to the user's system color scheme preference (light/dark mode). This feature uses CSS and the `prefers-color-scheme` media query to switch between defined color themes without requiring JavaScript.

### Dark Mode (Enabled by Default)

All charts created with GoSVGChart have dark mode support enabled by default. This means your charts will automatically adapt to the user's system preferences without any additional configuration.

If you want to customize the color themes or disable dark mode, you can use the following methods:

```go
// Create your chart normally
chart := gosvgchart.NewLineChart().
    SetTitle("Chart with Custom Themes").
    SetSize(600, 400).
    SetData([]float64{120, 250, 180, 310})

// Dark mode is already enabled by default
// You can customize the dark theme colors if desired:
chart.SetDarkTheme(
    "#121212", // Background color
    "#ffffff", // Text color
    "#aaaaaa", // Axis color
    "#333333", // Grid color
)

// You can also customize the light theme colors:
chart.SetLightTheme(
    "#ffffff", // Background color
    "#000000", // Text color
    "#666666", // Axis color 
    "#dddddd", // Grid color
)

// If you need to disable dark mode for some reason:
chart.EnableDarkModeSupport(false)
```

When the chart is rendered, it will automatically adapt based on the system's color scheme preferences. This works in modern browsers and SVG viewers that support the CSS `prefers-color-scheme` media query.

See `examples/dark_mode_example.go` for a complete example.

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

### Using Auto-Height with the Go API

You can use the `SetAutoHeight` method to automatically calculate the height based on the width:

```go
// Create a chart with auto-height
chart := gosvgchart.New().
    SetTitle("Monthly Sales").
    SetData([]float64{120, 250, 180, 310, 270, 390}).
    SetLabels([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}).
    SetColors([]string{"#3498db"}).
    SetSize(800, 0). // Height will be ignored when auto-height is enabled
    SetAutoHeight(true)

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
chart.SetCellSize(15).          // Size of each cell (will adapt to available space)
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

#### Heatmap Adaptivity

Heatmaps will automatically:

1. Use single-letter day labels (S, M, T, W, T, F, S) for each row
2. Calculate cell sizes based on available space to ensure the entire calendar fits
3. Adjust to different container sizes while maintaining the calendar structure

```go
// Create a heatmap chart with auto-sizing cells and single-letter day labels
heatmapChart := gosvgchart.NewHeatmapChart().
    SetTitle("Activity Heatmap").
    SetSize(800, 200).  // The cell size will automatically adapt to this space
    SetData(activityData).
    SetStartDate("2023-01-01").
    SetColors([]string{"#ebedf0", "#9be9a8", "#40c463", "#30a14e", "#216e39"})

// The SetCellSize() method is optional - if omitted, cells will be sized to fill available space
// Day labels are automatically set to single letters (S, M, T, W, T, F, S)

svg := heatmapChart.Render()
```

See `cmd/examples/main.go` for a working example of adaptive heatmaps with various container sizes.

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

### Using Auto-Sizing

You can use auto-sizing features for responsive charts that adapt to their container:

```gosvgchart
linechart
title: Responsive Chart
width: auto
height: auto
colors: #3498db

data:
Jan | 120
Feb | 250
Mar | 180
Apr | 310
May | 270
Jun | 390
```

Width options:
- `width: auto` or `width: 100%` - Full width of container
- `width: 50%` - Half width of container
- `width: 33%` - One-third width of container
- `width: 800` - Fixed width in pixels

When `height: auto` is specified, the height is automatically calculated based on the width (using a 16:9 aspect ratio for standard charts, or 250px for heatmaps).

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
width: auto
height: auto
colors: #3498db, #2ecc71

data:
Q1 | 850
Q2 | 940
Q3 | 1100
Q4 | 1200

---

barchart
title: 2024 Revenue
width: auto
height: auto
colors: #e74c3c, #f39c12

data:
Q1 | 950
Q2 | 1040
Q3 | 1200
Q4 | 1400
```

When using side-by-side charts:
- For 2 charts, each chart gets approximately 48% of the container width
- For 3 or more charts, each chart gets approximately 31% of the container width

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
| `SetAutoHeight(auto bool)` | Enables automatic height calculation based on width (16:9 ratio for standard charts, 250px for heatmaps) |
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

## Running Examples

GoSVGChart includes several examples to demonstrate different features:

### Individual Examples

The `examples/` directory contains standalone examples you can run directly:

```bash
go run examples/dark_mode_example.go
go run examples/heatmap_autosizing_example.go
```

### Unified Examples Command

The `cmd/examples/` directory contains a unified command that can run multiple examples:

```bash
# Build the examples command
go build -o chart-examples ./cmd/examples

# Run an example (default is "heatmap")
./chart-examples

# Run a specific example by type
./chart-examples -type heatmap
./chart-examples -type line
./chart-examples -type bar
./chart-examples -type pie
```

Each example will generate SVG files in the current directory that you can open in a web browser or SVG viewer:

- Heatmap: `heatmap_autosizing.svg` and `heatmap_autosizing_small.svg`
- Line chart: `line_chart_example.svg`
- Bar chart: `bar_chart_example.svg`
- Pie chart: `pie_chart_example.svg`

## License

MIT