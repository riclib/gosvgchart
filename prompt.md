# Chart Generation Instructions

When I ask you to visualize data or create a chart, please generate a chart specification using the format below. This chart will be rendered into an SVG image.

## Chart Format

The chart specification should be in this format:

```gosvgchart
charttype
title: Chart Title
width: auto
height: auto
colors: #color1, #color2, #color3

data:
Label1 | Value1
Label2 | Value2
Label3 | Value3
```

### Chart Types

Choose one of these chart types based on the data:

- `linechart` - For time series or trends over a continuous range
- `barchart` - For comparing values across categories
- `piechart` - For showing proportions of a whole
- `heatmapchart` - For showing activity patterns over time (GitHub-style)

### Properties

- `title` - The title of the chart
- `width` - Can be "auto" or "100%" for responsive display (recommended), or a specific pixel value (e.g., 800)
- `height` - Can be "auto" for proportional sizing (recommended), or a specific pixel value (e.g., 500)
- `colors` - Comma-separated list of hex color codes (e.g., #3498db, #e74c3c)

### Data Section

After the `data:` line, each data point should be on its own line with the format:
`Label | Value`

The label is a text description, and the value must be a number.

### Side-by-Side Charts

You can place multiple charts side by side using the `---` separator within a single code block:

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
- Width will automatically be set to ~48% for 2 charts
- Width will automatically be set to ~31% for 3+ charts

## Examples

### Line Chart Example

```gosvgchart
linechart
title: Monthly Sales 2024
width: auto
height: auto
colors: #3498db, #2ecc71

data:
Jan | 45000
Feb | 52000
Mar | 49000
Apr | 63000
May | 58000
Jun | 72000
```

### Bar Chart Example

```gosvgchart
barchart
title: Product Sales by Category
width: auto
height: auto
colors: #3498db, #e74c3c, #f39c12, #2ecc71

data:
Electronics | 125000
Clothing | 85000
Home & Garden | 62000
Toys | 43000
```

### Pie Chart Example

```gosvgchart
piechart
title: Market Share by Region
width: auto
height: auto
colors: #3498db, #e74c3c, #f39c12, #2ecc71, #9b59b6

data:
North America | 35
Europe | 28
Asia | 22
South America | 10
Other | 5
```

### Heatmap Chart Example

```gosvgchart
heatmapchart
title: GitHub Contributions
width: auto
height: auto
colors: #ebedf0, #9be9a8, #40c463, #30a14e, #216e39

data:
2025-01-01 | 5
2025-01-05 | 12
2025-01-10 | 3
2025-01-15 | 15
2025-01-20 | 8
2025-01-25 | 4
2025-02-01 | 7
2025-02-05 | 14
2025-02-10 | 6
2025-02-15 | 11
```

### Side-by-Side Comparison Example

```gosvgchart
piechart
title: 2023 Market Share
width: auto
height: auto
colors: #3498db, #e74c3c, #f39c12

data:
Product A | 45
Product B | 35
Product C | 20

---

piechart
title: 2024 Market Share
width: auto
height: auto
colors: #3498db, #e74c3c, #f39c12

data:
Product A | 30
Product B | 40
Product C | 30
```

## Important Notes

1. Always place the chart in a code block with the language specified as `gosvgchart` (```gosvgchart).
2. Make sure the data values are numeric.
3. Choose an appropriate chart type for the data.
4. Select suitable colors that work well together.
5. Ensure the data is properly formatted with the pipe symbol (|) separating labels and values.
6. Use `width: auto` and `height: auto` for responsive charts that adapt to the container size.
7. For comparing data side by side, use the `---` separator within a single chart block.

When asked to visualize data, analyze the data first, then choose the most appropriate chart type, and finally generate the chart specification in the format shown above. If multiple comparisons are needed, consider using side-by-side charts with the separator.