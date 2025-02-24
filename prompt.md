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

### Properties

- `title` - The title of the chart
- `width` - Width in pixels (typically 600-1000)
- `height` - Height in pixels (typically 400-600)
- `colors` - Comma-separated list of hex color codes (e.g., #3498db, #e74c3c)

### Data Section

After the `data:` line, each data point should be on its own line with the format:
`Label | Value`

The label is a text description, and the value must be a number.

### Side-by-Side Charts

You can place multiple charts side by side by putting the chart blocks right next to each other without any text between them:

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
```
```gosvgchart
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
```
```gosvgchart
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
6. For comparing data side by side, place multiple chart blocks directly adjacent to each other.

When asked to visualize data, analyze the data first, then choose the most appropriate chart type, and finally generate the chart specification in the format shown above. If multiple comparisons are needed, consider using side-by-side charts.