# Multiple Series Examples in Markdown Format

This document demonstrates how to use multiple series in charts using the markdown format.

## Line Chart with Multiple Series

```gosvgchart
linechart
title: Monthly Sales by Product
width: 800
height: auto
seriescolors: #4285F4, #EA4335, #FBBC05, #34A853

series: Product A
Jan | 120
Feb | 150
Mar | 180
Apr | 210
May | 240
Jun | 270

series: Product B
Jan | 200
Feb | 180
Mar | 160
Apr | 140
May | 120
Jun | 100

series: Product C
Jan | 50
Feb | 80
Mar | 110
Apr | 140
May | 170
Jun | 200
```

## Bar Chart with Multiple Series (Grouped)

```gosvgchart
barchart
title: Quarterly Revenue by Region
width: 800
height: auto
stacked: false
seriescolors: #4285F4, #EA4335, #FBBC05, #34A853

series: North
Q1 | 150
Q2 | 180
Q3 | 210
Q4 | 240

series: South
Q1 | 120
Q2 | 140
Q3 | 160
Q4 | 180

series: East
Q1 | 90
Q2 | 110
Q3 | 130
Q4 | 150

series: West
Q1 | 180
Q2 | 200
Q3 | 220
Q4 | 240
```

## Bar Chart with Multiple Series (Stacked)

```gosvgchart
barchart
title: Quarterly Revenue by Region (Stacked)
width: 800
height: auto
stacked: true
seriescolors: #4285F4, #EA4335, #FBBC05, #34A853

series: North
Q1 | 150
Q2 | 180
Q3 | 210
Q4 | 240

series: South
Q1 | 120
Q2 | 140
Q3 | 160
Q4 | 180

series: East
Q1 | 90
Q2 | 110
Q3 | 130
Q4 | 150

series: West
Q1 | 180
Q2 | 200
Q3 | 220
Q4 | 240
```

## Comparing Multiple Series Side by Side

You can also place multiple charts side by side using the `---` separator:

```gosvgchart
linechart
title: Product A Sales
width: 450
height: 300
colors: #4285F4

data:
Jan | 120
Feb | 150
Mar | 180
Apr | 210
May | 240
Jun | 270

---

linechart
title: Product B Sales
width: 450
height: 300
colors: #EA4335

data:
Jan | 200
Feb | 180
Mar | 160
Apr | 140
May | 120
Jun | 100
```

## Legacy Single Series Format (For Comparison)

The library still supports the legacy single series format:

```gosvgchart
barchart
title: Monthly Sales
width: 800
height: auto
colors: #4285F4, #EA4335, #FBBC05, #34A853

data:
Jan | 120
Feb | 150
Mar | 180
Apr | 210
May | 240
Jun | 270
```

## Key Features of Multiple Series Support

1. Use `series: Name` to start a new data series
2. Each series can have its own data points
3. Use `seriescolors: #color1, #color2, ...` to set colors for each series
4. For bar charts, use `stacked: true` or `stacked: false` to control the display style
5. Legends are automatically shown when multiple series are present 

## Tabular Format for Multiple Series (New)

GoSVGChart now supports a more intuitive tabular format for defining multiple series:

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
Apr | 210 | 140 | 140
May | 240 | 120 | 170
Jun | 270 | 100 | 200
```

The tabular format can also be used for bar charts:

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

### Benefits of the Tabular Format:
1. More intuitive and readable
2. Easier to maintain and edit
3. Better alignment of data points across series
4. Reduced repetition of labels
5. Clearer visualization of the data structure 