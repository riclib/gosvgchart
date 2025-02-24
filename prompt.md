# Chart Generation Instructions

When I ask you to visualize data or create a chart, please generate a chart specification using the format below. This chart will be rendered into an SVG image.

## Chart Format

The chart specification should be in this format:

```
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

## Examples

### Line Chart Example

```
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

```
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

```
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

## Important Notes

1. Always place the chart in a code block (```).
2. Make sure the data values are numeric.
3. Choose an appropriate chart type for the data.
4. Select suitable colors that work well together.
5. Ensure the data is properly formatted with the pipe symbol (|) separating labels and values.

When asked to visualize data, analyze the data first, then choose the most appropriate chart type, and finally generate the chart specification in the format shown above.