package gosvgchart

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

// Chart is the interface that all chart types must implement
type Chart interface {
	SetTitle(title string) Chart
	SetSize(width, height int) Chart
	SetData(data []float64) Chart
	SetLabels(labels []string) Chart
	SetColors(colors []string) Chart
	Render() string
}

// BaseChart contains common properties and methods for all chart types
type BaseChart struct {
	ChartType  string
	Title      string
	Width      int
	Height     int
	Data       []float64
	Labels     []string
	Colors     []string
	ShowTitle  bool
	ShowLegend bool
	Margin     struct {
		Top    int
		Right  int
		Bottom int
		Left   int
	}
	BackgroundColor string
}

// LineChart implements a line chart
type LineChart struct {
	BaseChart
	ShowPoints bool
	Smooth     bool
}

// BarChart implements a bar chart
type BarChart struct {
	BaseChart
	Horizontal bool
	Stacked    bool
}

// PieChart implements a pie/donut chart
type PieChart struct {
	BaseChart
	DonutHolePercentage float64
}

// HeatmapChart implements a heatmap chart similar to GitHub's activity heatmap
type HeatmapChart struct {
	BaseChart
	CellSize     int      // Size of each cell in pixels
	CellSpacing  int      // Spacing between cells in pixels
	CellRounding int      // Corner radius of cells
	DateFormat   string   // Date format string
	DayLabels    []string // Labels for days of week (Sunday-Saturday)
	MonthLabels  []string // Labels for months
	MaxValue     float64  // Maximum value for color scaling (0 for auto)
}

// New creates a new chart with default settings
func New() Chart {
	// Default to line chart
	return NewLineChart()
}

// NewLineChart creates a new line chart with default settings
func NewLineChart() *LineChart {
	chart := &LineChart{
		BaseChart: BaseChart{
			ChartType:       "line",
			Width:           800,
			Height:          500,
			ShowTitle:       true,
			ShowLegend:      true,
			BackgroundColor: "#ffffff",
		},
		ShowPoints: true,
		Smooth:     false,
	}
	
	chart.Margin.Top = 50
	chart.Margin.Right = 50
	chart.Margin.Bottom = 50
	chart.Margin.Left = 50
	
	// Default colors
	chart.Colors = []string{"#3498db", "#e74c3c", "#2ecc71", "#f39c12", "#9b59b6"}
	
	return chart
}

// NewBarChart creates a new bar chart with default settings
func NewBarChart() *BarChart {
	chart := &BarChart{
		BaseChart: BaseChart{
			ChartType:       "bar",
			Width:           800,
			Height:          500,
			ShowTitle:       true,
			ShowLegend:      true,
			BackgroundColor: "#ffffff",
		},
		Horizontal: false,
		Stacked:    false,
	}
	
	chart.Margin.Top = 50
	chart.Margin.Right = 50
	chart.Margin.Bottom = 50
	chart.Margin.Left = 50
	
	// Default colors
	chart.Colors = []string{"#3498db", "#e74c3c", "#2ecc71", "#f39c12", "#9b59b6"}
	
	return chart
}

// NewPieChart creates a new pie chart with default settings
func NewPieChart() *PieChart {
	chart := &PieChart{
		BaseChart: BaseChart{
			ChartType:       "pie",
			Width:           800,
			Height:          500,
			ShowTitle:       true,
			ShowLegend:      true,
			BackgroundColor: "#ffffff",
		},
		DonutHolePercentage: 0, // 0 means a regular pie chart
	}
	
	chart.Margin.Top = 50
	chart.Margin.Right = 50
	chart.Margin.Bottom = 50
	chart.Margin.Left = 50
	
	// Default colors
	chart.Colors = []string{"#3498db", "#e74c3c", "#2ecc71", "#f39c12", "#9b59b6"}
	
	return chart
}

// NewHeatmapChart creates a new heatmap chart with default settings
func NewHeatmapChart() *HeatmapChart {
	chart := &HeatmapChart{
		BaseChart: BaseChart{
			ChartType:       "heatmap",
			Width:           800,
			Height:          200,
			ShowTitle:       true,
			ShowLegend:      true,
			BackgroundColor: "#ffffff",
		},
		CellSize:     15,
		CellSpacing:  3,
		CellRounding: 2,
		DateFormat:   "2006-01-02",
		DayLabels:    []string{"", "Mon", "", "Wed", "", "Fri", ""},
		MonthLabels:  []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		MaxValue:     0, // 0 means auto-scale
	}
	
	chart.Margin.Top = 50
	chart.Margin.Right = 50
	chart.Margin.Bottom = 50
	chart.Margin.Left = 50
	
	// Default colors - from light to dark for intensity
	chart.Colors = []string{"#ebedf0", "#9be9a8", "#40c463", "#30a14e", "#216e39"}
	
	return chart
}

// LineChart methods to implement Chart interface

// SetTitle sets the chart title
func (c *LineChart) SetTitle(title string) Chart {
	c.Title = title
	return c
}

// SetSize sets the chart dimensions in pixels
func (c *LineChart) SetSize(width, height int) Chart {
	c.Width = width
	c.Height = height
	return c
}

// SetData sets the chart data values
func (c *LineChart) SetData(data []float64) Chart {
	c.Data = data
	return c
}

// SetLabels sets the chart labels
func (c *LineChart) SetLabels(labels []string) Chart {
	c.Labels = labels
	return c
}

// SetColors sets the color palette as hex values
func (c *LineChart) SetColors(colors []string) Chart {
	c.Colors = colors
	return c
}

// BarChart methods to implement Chart interface

// SetTitle sets the chart title
func (c *BarChart) SetTitle(title string) Chart {
	c.Title = title
	return c
}

// SetSize sets the chart dimensions in pixels
func (c *BarChart) SetSize(width, height int) Chart {
	c.Width = width
	c.Height = height
	return c
}

// SetData sets the chart data values
func (c *BarChart) SetData(data []float64) Chart {
	c.Data = data
	return c
}

// SetLabels sets the chart labels
func (c *BarChart) SetLabels(labels []string) Chart {
	c.Labels = labels
	return c
}

// SetColors sets the color palette as hex values
func (c *BarChart) SetColors(colors []string) Chart {
	c.Colors = colors
	return c
}

// PieChart methods to implement Chart interface

// SetTitle sets the chart title
func (c *PieChart) SetTitle(title string) Chart {
	c.Title = title
	return c
}

// SetSize sets the chart dimensions in pixels
func (c *PieChart) SetSize(width, height int) Chart {
	c.Width = width
	c.Height = height
	return c
}

// SetData sets the chart data values
func (c *PieChart) SetData(data []float64) Chart {
	c.Data = data
	return c
}

// SetLabels sets the chart labels
func (c *PieChart) SetLabels(labels []string) Chart {
	c.Labels = labels
	return c
}

// SetColors sets the color palette as hex values
func (c *PieChart) SetColors(colors []string) Chart {
	c.Colors = colors
	return c
}

// HeatmapChart methods to implement Chart interface

// SetTitle sets the chart title
func (c *HeatmapChart) SetTitle(title string) Chart {
	c.Title = title
	return c
}

// SetSize sets the chart dimensions in pixels
func (c *HeatmapChart) SetSize(width, height int) Chart {
	c.Width = width
	c.Height = height
	return c
}

// SetData sets the chart data values
func (c *HeatmapChart) SetData(data []float64) Chart {
	c.Data = data
	return c
}

// SetLabels sets the chart labels (should be ISO dates: YYYY-MM-DD)
func (c *HeatmapChart) SetLabels(labels []string) Chart {
	c.Labels = labels
	return c
}

// SetColors sets the color palette as hex values (from least to most intense)
func (c *HeatmapChart) SetColors(colors []string) Chart {
	c.Colors = colors
	return c
}

// SetCellSize sets the size of each cell in pixels
func (c *HeatmapChart) SetCellSize(size int) *HeatmapChart {
	c.CellSize = size
	return c
}

// SetCellSpacing sets the spacing between cells in pixels
func (c *HeatmapChart) SetCellSpacing(spacing int) *HeatmapChart {
	c.CellSpacing = spacing
	return c
}

// SetCellRounding sets the corner radius of cells
func (c *HeatmapChart) SetCellRounding(radius int) *HeatmapChart {
	c.CellRounding = radius
	return c
}

// SetMaxValue sets the maximum value for color scaling
func (c *HeatmapChart) SetMaxValue(max float64) *HeatmapChart {
	c.MaxValue = max
	return c
}

// SetDateFormat sets the date format string (Go time format)
func (c *HeatmapChart) SetDateFormat(format string) *HeatmapChart {
	c.DateFormat = format
	return c
}

// SetDayLabels sets the labels for days of the week
func (c *HeatmapChart) SetDayLabels(labels []string) *HeatmapChart {
	c.DayLabels = labels
	return c
}

// SetMonthLabels sets the labels for months
func (c *HeatmapChart) SetMonthLabels(labels []string) *HeatmapChart {
	c.MonthLabels = labels
	return c
}

// ShowDataPoints shows or hides data points for line charts
func (c *LineChart) ShowDataPoints(show bool) *LineChart {
	c.ShowPoints = show
	return c
}

// SetSmooth enables smooth curved lines
func (c *LineChart) SetSmooth(smooth bool) *LineChart {
	c.Smooth = smooth
	return c
}

// SetHorizontal displays bars horizontally
func (c *BarChart) SetHorizontal(horizontal bool) *BarChart {
	c.Horizontal = horizontal
	return c
}

// SetStacked stacks multiple data series
func (c *BarChart) SetStacked(stacked bool) *BarChart {
	c.Stacked = stacked
	return c
}

// SetDonutHole sets the inner circle size for donut charts
func (c *PieChart) SetDonutHole(percentage float64) *PieChart {
	if percentage < 0 {
		percentage = 0
	}
	if percentage > 0.9 {
		percentage = 0.9
	}
	c.DonutHolePercentage = percentage
	return c
}

// Render renders the line chart to an SVG string
func (c *LineChart) Render() string {
	var svg strings.Builder
	
	svg.WriteString(fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`, c.Width, c.Height))
	
	// Background
	svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="%s"/>`, c.Width, c.Height, c.BackgroundColor))
	
	// Title
	if c.ShowTitle && c.Title != "" {
		svg.WriteString(fmt.Sprintf(`<text x="%d" y="30" text-anchor="middle" font-family="Arial" font-size="20" font-weight="bold">%s</text>`,
			c.Width/2, c.Title))
	}
	
	// Chart area dimensions
	chartWidth := c.Width - c.Margin.Left - c.Margin.Right
	chartHeight := c.Height - c.Margin.Top - c.Margin.Bottom
	
	// Calculate scales
	var maxValue float64
	for _, v := range c.Data {
		if v > maxValue {
			maxValue = v
		}
	}
	
	// Add 10% padding to the max value
	maxValue *= 1.1
	
	// Draw axes
	svg.WriteString(fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="2"/>`,
		c.Margin.Left, c.Height-c.Margin.Bottom, c.Width-c.Margin.Right, c.Height-c.Margin.Bottom))
	svg.WriteString(fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="2"/>`,
		c.Margin.Left, c.Margin.Top, c.Margin.Left, c.Height-c.Margin.Bottom))
	
	// Draw data
	if len(c.Data) > 0 {
		// Calculate point coordinates
		points := make([][2]int, len(c.Data))
		for i, v := range c.Data {
			x := c.Margin.Left + i*chartWidth/(len(c.Data)-1)
			if len(c.Data) == 1 {
				x = c.Margin.Left + chartWidth/2
			}
			y := c.Height - c.Margin.Bottom - int(v/maxValue*float64(chartHeight))
			points[i] = [2]int{x, y}
		}
		
		// Draw line
		var path strings.Builder
		path.WriteString(fmt.Sprintf(`<path d="M%d,%d`, points[0][0], points[0][1]))
		for i := 1; i < len(points); i++ {
			if c.Smooth && i < len(points)-1 {
				// Calculate control points for smooth curve
				x1 := points[i-1][0]
				y1 := points[i-1][1]
				x2 := points[i][0]
				y2 := points[i][1]
				xc := (x1 + x2) / 2
				path.WriteString(fmt.Sprintf(" Q%d,%d %d,%d", xc, y1, xc, (y1+y2)/2))
				path.WriteString(fmt.Sprintf(" Q%d,%d %d,%d", xc, y2, x2, y2))
			} else {
				path.WriteString(fmt.Sprintf(" L%d,%d", points[i][0], points[i][1]))
			}
		}
		path.WriteString(`" fill="none" stroke="`+c.Colors[0]+`" stroke-width="3"/>`)
		svg.WriteString(path.String())
		
		// Draw points if enabled
		if c.ShowPoints {
			for _, p := range points {
				svg.WriteString(fmt.Sprintf(`<circle cx="%d" cy="%d" r="5" fill="%s"/>`, p[0], p[1], c.Colors[0]))
			}
		}
		
		// Draw labels if available
		if len(c.Labels) > 0 {
			for i, p := range points {
				if i < len(c.Labels) {
					svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="12">%s</text>`,
						p[0], c.Height-c.Margin.Bottom+20, c.Labels[i]))
				}
			}
		}
	}
	
	svg.WriteString("</svg>")
	return svg.String()
}

// Render renders the bar chart to an SVG string
func (c *BarChart) Render() string {
	var svg strings.Builder
	
	svg.WriteString(fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`, c.Width, c.Height))
	
	// Background
	svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="%s"/>`, c.Width, c.Height, c.BackgroundColor))
	
	// Title
	if c.ShowTitle && c.Title != "" {
		svg.WriteString(fmt.Sprintf(`<text x="%d" y="30" text-anchor="middle" font-family="Arial" font-size="20" font-weight="bold">%s</text>`,
			c.Width/2, c.Title))
	}
	
	// Chart area dimensions
	chartWidth := c.Width - c.Margin.Left - c.Margin.Right
	chartHeight := c.Height - c.Margin.Top - c.Margin.Bottom
	
	// Calculate scales
	var maxValue float64
	for _, v := range c.Data {
		if v > maxValue {
			maxValue = v
		}
	}
	
	// Add 10% padding to the max value
	maxValue *= 1.1
	
	// Draw axes
	svg.WriteString(fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="2"/>`,
		c.Margin.Left, c.Height-c.Margin.Bottom, c.Width-c.Margin.Right, c.Height-c.Margin.Bottom))
	svg.WriteString(fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="2"/>`,
		c.Margin.Left, c.Margin.Top, c.Margin.Left, c.Height-c.Margin.Bottom))
	
	// Draw data
	if len(c.Data) > 0 {
		barWidth := chartWidth / (len(c.Data) * 2)
		
		// Draw bars
		for i, v := range c.Data {
			barHeight := int(v / maxValue * float64(chartHeight))
			barX := c.Margin.Left + i*(chartWidth/len(c.Data)) + (chartWidth/len(c.Data)-barWidth)/2
			barY := c.Height - c.Margin.Bottom - barHeight
			
			// Determine color (cycle through available colors)
			colorIndex := i % len(c.Colors)
			color := c.Colors[colorIndex]
			
			svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" fill="%s"/>`,
				barX, barY, barWidth, barHeight, color))
				
			// Add value text on top of bar
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="12" fill="black">%.0f</text>`,
				barX+barWidth/2, barY-5, v))
		}
		
		// Draw labels if available
		if len(c.Labels) > 0 {
			for i := 0; i < len(c.Data); i++ {
				if i < len(c.Labels) {
					barX := c.Margin.Left + i*(chartWidth/len(c.Data)) + (chartWidth/len(c.Data))/2
					svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="12">%s</text>`,
						barX, c.Height-c.Margin.Bottom+20, c.Labels[i]))
				}
			}
		}
	}
	
	svg.WriteString("</svg>")
	return svg.String()
}

// Render renders the pie chart to an SVG string
func (c *PieChart) Render() string {
	var svg strings.Builder
	
	svg.WriteString(fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`, c.Width, c.Height))
	
	// Background
	svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="%s"/>`, c.Width, c.Height, c.BackgroundColor))
	
	// Title
	if c.ShowTitle && c.Title != "" {
		svg.WriteString(fmt.Sprintf(`<text x="%d" y="30" text-anchor="middle" font-family="Arial" font-size="20" font-weight="bold">%s</text>`,
			c.Width/2, c.Title))
	}
	
	// Calculate total
	var total float64
	for _, v := range c.Data {
		total += v
	}
	
	// Center and radius
	centerX := c.Width / 2
	centerY := c.Height / 2
	radius := int(math.Min(float64(c.Width-c.Margin.Left-c.Margin.Right), 
					   float64(c.Height-c.Margin.Top-c.Margin.Bottom))) / 2
	
	innerRadius := int(float64(radius) * c.DonutHolePercentage)
	
	// Draw pie slices
	if len(c.Data) > 0 && total > 0 {
		var startAngle float64
		
		for i, v := range c.Data {
			// Calculate angles
			sliceAngle := v / total * 2 * math.Pi
			endAngle := startAngle + sliceAngle
			
			// Calculate points
			x1 := centerX + int(math.Cos(startAngle)*float64(radius))
			y1 := centerY + int(math.Sin(startAngle)*float64(radius))
			x2 := centerX + int(math.Cos(endAngle)*float64(radius))
			y2 := centerY + int(math.Sin(endAngle)*float64(radius))
			
			// Determine large arc flag
			largeArcFlag := 0
			if sliceAngle > math.Pi {
				largeArcFlag = 1
			}
			
			// Determine color (cycle through available colors)
			colorIndex := i % len(c.Colors)
			color := c.Colors[colorIndex]
			
			// Draw path
			if c.DonutHolePercentage > 0 {
				// For donut chart, draw more complex path
				x1Inner := centerX + int(math.Cos(startAngle)*float64(innerRadius))
				y1Inner := centerY + int(math.Sin(startAngle)*float64(innerRadius))
				x2Inner := centerX + int(math.Cos(endAngle)*float64(innerRadius))
				y2Inner := centerY + int(math.Sin(endAngle)*float64(innerRadius))
				
				svg.WriteString(fmt.Sprintf(`<path d="M%d,%d L%d,%d A%d,%d 0 %d,1 %d,%d L%d,%d A%d,%d 0 %d,0 %d,%d Z" fill="%s"/>`,
					x1Inner, y1Inner, x1, y1, radius, radius, largeArcFlag, x2, y2, x2Inner, y2Inner, innerRadius, innerRadius, largeArcFlag, x1Inner, y1Inner, color))
			} else {
				// For regular pie chart, draw simple wedge
				svg.WriteString(fmt.Sprintf(`<path d="M%d,%d L%d,%d A%d,%d 0 %d,1 %d,%d L%d,%d Z" fill="%s"/>`,
					centerX, centerY, x1, y1, radius, radius, largeArcFlag, x2, y2, centerX, centerY, color))
			}
			
			// Label position (middle of slice)
			labelAngle := startAngle + sliceAngle/2
			labelDistance := float64(radius) * 0.7
			labelX := centerX + int(math.Cos(labelAngle)*labelDistance)
			labelY := centerY + int(math.Sin(labelAngle)*labelDistance)
			
			// Draw percentage
			percentage := v / total * 100
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="12" fill="white">%.1f%%</text>`,
				labelX, labelY, percentage))
			
			startAngle = endAngle
		}
		
		// Draw legend
		if c.ShowLegend && len(c.Labels) > 0 {
			legendX := c.Width - c.Margin.Right + 20
			legendY := c.Margin.Top
			
			for i, label := range c.Labels {
				if i < len(c.Data) {
					colorIndex := i % len(c.Colors)
					color := c.Colors[colorIndex]
					
					// Draw color box
					svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="15" height="15" fill="%s"/>`,
						legendX, legendY, color))
					
					// Draw label
					svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="12">%s</text>`,
						legendX + 20, legendY + 12, label))
					
					legendY += 25
				}
			}
		}
	}
	
	svg.WriteString("</svg>")
	return svg.String()
}

// Render renders the heatmap chart to an SVG string
func (c *HeatmapChart) Render() string {
	var svg strings.Builder

	svg.WriteString(fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`, c.Width, c.Height))
	
	// Background
	svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="%s"/>`, c.Width, c.Height, c.BackgroundColor))
	
	// Title
	if c.ShowTitle && c.Title != "" {
		svg.WriteString(fmt.Sprintf(`<text x="%d" y="30" text-anchor="middle" font-family="Arial" font-size="20" font-weight="bold">%s</text>`,
			c.Width/2, c.Title))
	}

	// Initialize dates and values map
	dateMap := make(map[string]float64)
	var dates []time.Time
	
	// Parse dates from labels
	for i, label := range c.Labels {
		if i < len(c.Data) {
			date, err := time.Parse(c.DateFormat, label)
			if err == nil {
				dateMap[label] = c.Data[i]
				dates = append(dates, date)
			}
		}
	}
	
	// Sort dates
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})
	
	// If no data, return empty SVG
	if len(dates) == 0 {
		svg.WriteString("</svg>")
		return svg.String()
	}
	
	// Find the first Sunday before the start date
	startDate := dates[0]
	daysToSubtract := int(startDate.Weekday())
	firstSunday := startDate.AddDate(0, 0, -daysToSubtract)
	
	// Find the last Saturday after the end date
	endDate := dates[len(dates)-1]
	daysToAdd := 6 - int(endDate.Weekday())
	lastSaturday := endDate.AddDate(0, 0, daysToAdd)
	
	// Calculate total weeks
	totalDays := int(lastSaturday.Sub(firstSunday).Hours() / 24) + 1
	totalWeeks := totalDays / 7
	
	// Find max value for color scaling
	maxVal := c.MaxValue
	if maxVal <= 0 {
		// Auto-scale
		for _, v := range c.Data {
			if v > maxVal {
				maxVal = v
			}
		}
	}
	
	// Starting position for the grid
	startX := c.Margin.Left + 30 // Space for day labels
	startY := c.Margin.Top + 50  // Space for title and month labels
	
	// Draw day labels (Mon, Tue, etc.)
	for i, label := range c.DayLabels {
		if label != "" {
			labelY := startY + i*(c.CellSize+c.CellSpacing) + c.CellSize/2 + 5
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="end">%s</text>`,
				startX-5, labelY, label))
		}
	}
	
	// Draw the heatmap grid
	currentDate := firstSunday
	for week := 0; week < totalWeeks; week++ {
		// Check if we need to draw month label
		if currentDate.Day() <= 7 {
			// This is the first week of the month
			monthLabel := c.MonthLabels[currentDate.Month()-1]
			labelX := startX + week*(c.CellSize+c.CellSpacing) + c.CellSize/2
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="middle">%s</text>`,
				labelX, startY-5, monthLabel))
		}
		
		for day := 0; day < 7; day++ {
			dateStr := currentDate.Format(c.DateFormat)
			value := 0.0
			
			// Check if we have data for this date
			if val, ok := dateMap[dateStr]; ok {
				value = val
			}
			
			// Calculate color based on value
			colorIndex := 0
			if maxVal > 0 {
				// Scale value from 0 to len(colors)-1
				colorIndex = int(math.Min(float64(len(c.Colors)-1), math.Floor(value/maxVal*float64(len(c.Colors)))))
			}
			
			color := c.Colors[0] // Default to lowest color
			if colorIndex < len(c.Colors) && colorIndex >= 0 {
				color = c.Colors[colorIndex]
			}
			
			// Calculate cell position
			cellX := startX + week*(c.CellSize+c.CellSpacing)
			cellY := startY + day*(c.CellSize+c.CellSpacing)
			
			// Draw cell
			svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" rx="%d" ry="%d" fill="%s">`,
				cellX, cellY, c.CellSize, c.CellSize, c.CellRounding, c.CellRounding, color))
				
			// Add tooltip
			svg.WriteString(fmt.Sprintf(`<title>%s: %v</title></rect>`, dateStr, value))
			
			// Move to next day
			currentDate = currentDate.AddDate(0, 0, 1)
		}
	}
	
	// Add legend
	if c.ShowLegend {
		legendX := c.Margin.Left
		legendY := startY + 7*(c.CellSize+c.CellSpacing) + 30
		legendLabelY := legendY + c.CellSize/2 + 5
		
		svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="start">Less</text>`,
			legendX, legendLabelY))
		
		for i, color := range c.Colors {
			cellX := legendX + 40 + i*(c.CellSize+c.CellSpacing)
			svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" rx="%d" ry="%d" fill="%s"/>`,
				cellX, legendY, c.CellSize, c.CellSize, c.CellRounding, c.CellRounding, color))
		}
		
		svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="start">More</text>`,
			legendX + 40 + len(c.Colors)*(c.CellSize+c.CellSpacing) + 5, legendLabelY))
	}
	
	svg.WriteString("</svg>")
	return svg.String()
}