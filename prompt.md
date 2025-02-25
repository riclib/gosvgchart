# Chart Generation Instructions

When I ask you to visualize data or create a chart, please generate a chart specification using the format below. This chart will be rendered into an SVG image.

## Chart Format

The chart specification should be in this format:

```gosvgchart
charttype
title: Chart Title
width: 800
height: 500
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
- `width` - Width in pixels (typically 600-1000)
- `height` - Height in pixels (typically 400-600)
- `colors` - Comma-separated list of hex color codes (e.g., #3498db, #e74c3c)
- `seriescolors` - Comma-separated list of hex color codes for multiple series (e.g., #3498db, #e74c3c)
- `stacked` - For bar charts with multiple series, set to `true` for stacked bars or `false` for grouped bars

### Data Section

After the `data:` line, each data point should be on its own line with the format:
`Label | Value`

The label is a text description, and the value must be a number.

### Multiple Series Support

For charts with multiple data series (line charts and bar charts), you can use either of these formats:

#### Traditional Format

```gosvgchart
linechart
title: Monthly Sales by Product
width: 800
height: auto
seriescolors: #4285F4, #EA4335, #FBBC05

series: Product A
Jan | 120
Feb | 150
Mar | 180

series: Product B
Jan | 200
Feb | 180
Mar | 160
```

#### Tabular Format (Recommended)

The tabular format is more intuitive and easier to read:

```gosvgchart
linechart
title: Monthly Sales by Product
width: 800
height: auto
seriescolors: #4285F4, #EA4335, #FBBC05

series:
Month | Product A | Product B | Product C
Jan | 120 | 200 | 50
Feb | 150 | 180 | 80
Mar | 180 | 160 | 110
```

For bar charts, you can use `stacked: true` or `stacked: false` to control whether the bars are stacked or grouped:

```gosvgchart
barchart
title: Quarterly Revenue by Region
width: 800
height: auto
stacked: false
seriescolors: #4285F4, #EA4335, #FBBC05, #34A853

series:
Quarter | North | South | East | West
Q1 | 150 | 120 | 90 | 180
Q2 | 180 | 140 | 110 | 200
Q3 | 210 | 160 | 130 | 220
Q4 | 240 | 180 | 150 | 240
```

### Side-by-Side Charts

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

## Examples

### Line Chart Example

```gosvgchart
linechart
title: Monthly Sales 2024
width: 800
height: 500
colors: #3498db, #2ecc71

data:
Jan | 45000
Feb | 52000
Mar | 49000
Apr | 63000
May | 58000
Jun | 72000
```

### Multiple Series Line Chart Example

```gosvgchart
linechart
title: Monthly Sales by Product
width: 800
height: 500
seriescolors: #4285F4, #EA4335, #FBBC05

series:
Month | Product A | Product B | Product C
Jan | 45000 | 32000 | 18000
Feb | 52000 | 34000 | 20000
Mar | 49000 | 36000 | 22000
Apr | 63000 | 38000 | 24000
May | 58000 | 40000 | 26000
Jun | 72000 | 42000 | 28000
```

### Bar Chart Example

```gosvgchart
barchart
title: Product Sales by Category
width: 700
height: 500
colors: #3498db, #e74c3c, #f39c12, #2ecc71

data:
Electronics | 125000
Clothing | 85000
Home & Garden | 62000
Toys | 43000
```

### Multiple Series Bar Chart Example (Grouped)

```gosvgchart
barchart
title: Quarterly Revenue by Region
width: 800
height: 500
stacked: false
seriescolors: #4285F4, #EA4335, #FBBC05, #34A853

series:
Quarter | North | South | East | West
Q1 | 150 | 120 | 90 | 180
Q2 | 180 | 140 | 110 | 200
Q3 | 210 | 160 | 130 | 220
Q4 | 240 | 180 | 150 | 240
```

### Multiple Series Bar Chart Example (Stacked)

```gosvgchart
barchart
title: Quarterly Revenue by Region
width: 800
height: 500
stacked: true
seriescolors: #4285F4, #EA4335, #FBBC05, #34A853

series:
Quarter | North | South | East | West
Q1 | 150 | 120 | 90 | 180
Q2 | 180 | 140 | 110 | 200
Q3 | 210 | 160 | 130 | 220
Q4 | 240 | 180 | 150 | 240
```

### Pie Chart Example

```gosvgchart
piechart
title: Market Share by Region
width: 600
height: 600
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
width: 800
height: 200
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
width: 400
height: 400
colors: #3498db, #e74c3c, #f39c12

data:
Product A | 45
Product B | 35
Product C | 20

---

piechart
title: 2024 Market Share
width: 400
height: 400
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
6. For multiple series, use the tabular format for better readability when possible.
7. For comparing data side by side, use the `---` separator within a single code block.
8. For bar charts with multiple series, specify `stacked: true` or `stacked: false` to control the display style.

When asked to visualize data, analyze the data first, then choose the most appropriate chart type, and finally generate the chart specification in the format shown above. If multiple comparisons are needed, consider using multiple series or side-by-side charts.