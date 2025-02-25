# Multiple Series Examples

## Line Chart with Multiple Series

```chart
linechart
title: Monthly Sales by Product
width: 800
height: auto
colors: #4285F4, #EA4335, #FBBC05
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

```chart
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

```chart
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

## Legacy Single Series Bar Chart

```chart
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