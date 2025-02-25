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
	SetAutoHeight(auto bool) Chart
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
	AutoHeight bool
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
	DarkModeSupport bool
	DarkTheme       struct {
		BackgroundColor string
		TextColor       string
		AxisColor       string
		GridColor       string
	}
	LightTheme struct {
		BackgroundColor string
		TextColor       string
		AxisColor       string
		GridColor       string
	}
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
	MaxLabelLength      int  // Maximum label length before truncation
	ShowTooltips        bool // Show tooltips on hover for truncated labels
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
			AutoHeight:      false,
			ShowTitle:       true,
			ShowLegend:      true,
			BackgroundColor: "#ffffff",
			DarkModeSupport: true, // Enable dark mode by default
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

	// Set up default themes
	chart.EnableDarkModeSupport(true)

	return chart
}

// NewBarChart creates a new bar chart with default settings
func NewBarChart() *BarChart {
	chart := &BarChart{
		BaseChart: BaseChart{
			ChartType:       "bar",
			Width:           800,
			Height:          500,
			AutoHeight:      false,
			ShowTitle:       true,
			ShowLegend:      true,
			BackgroundColor: "#ffffff",
			DarkModeSupport: true, // Enable dark mode by default
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

	// Set up default themes
	chart.EnableDarkModeSupport(true)

	return chart
}

// NewPieChart creates a new pie chart with default settings
func NewPieChart() *PieChart {
	chart := &PieChart{
		BaseChart: BaseChart{
			ChartType:       "pie",
			Width:           800,
			Height:          500,
			AutoHeight:      false,
			ShowTitle:       true,
			ShowLegend:      true,
			BackgroundColor: "#ffffff",
			DarkModeSupport: true, // Enable dark mode by default
		},
		DonutHolePercentage: 0,    // 0 means a regular pie chart
		MaxLabelLength:      10,   // Default to 10 characters for legend labels
		ShowTooltips:        true, // Enable tooltips by default
	}

	chart.Margin.Top = 50
	chart.Margin.Right = 120 // Increased right margin to accommodate longer labels
	chart.Margin.Bottom = 50
	chart.Margin.Left = 50

	// Default colors
	chart.Colors = []string{"#3498db", "#e74c3c", "#2ecc71", "#f39c12", "#9b59b6"}

	// Set up default themes
	chart.EnableDarkModeSupport(true)

	return chart
}

// NewHeatmapChart creates a new heatmap chart with default settings
func NewHeatmapChart() *HeatmapChart {
	chart := &HeatmapChart{
		BaseChart: BaseChart{
			ChartType:       "heatmap",
			Width:           800,
			Height:          200,
			AutoHeight:      false,
			ShowTitle:       true,
			ShowLegend:      true,
			BackgroundColor: "#ffffff",
			DarkModeSupport: true, // Enable dark mode by default
		},
		CellSize:     15,
		CellSpacing:  3,
		CellRounding: 2,
		DateFormat:   "2006-01-02",
		// We'll use the single-letter day labels defined in the Render method
		// so we don't need to set them here anymore
		DayLabels:   []string{},
		MonthLabels: []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		MaxValue:    0, // 0 means auto-scale
	}

	chart.Margin.Top = 50
	chart.Margin.Right = 50
	chart.Margin.Bottom = 50
	chart.Margin.Left = 50

	// Default colors - from light to dark for intensity
	chart.Colors = []string{"#ebedf0", "#9be9a8", "#40c463", "#30a14e", "#216e39"}

	// Set up default themes
	chart.EnableDarkModeSupport(true)

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
	c.AutoHeight = false
	return c
}

// SetAutoHeight enables automatic height calculation based on width
func (c *LineChart) SetAutoHeight(auto bool) Chart {
	c.AutoHeight = auto
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
	c.AutoHeight = false
	return c
}

// SetAutoHeight enables automatic height calculation based on width
func (c *BarChart) SetAutoHeight(auto bool) Chart {
	c.AutoHeight = auto
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
	c.AutoHeight = false
	return c
}

// SetAutoHeight enables automatic height calculation based on width
func (c *PieChart) SetAutoHeight(auto bool) Chart {
	c.AutoHeight = auto
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
	c.AutoHeight = false
	return c
}

// SetAutoHeight enables automatic height calculation based on width
func (c *HeatmapChart) SetAutoHeight(auto bool) Chart {
	c.AutoHeight = auto
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
// Note: This method is deprecated as single-letter day labels are now used by default
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

// SetMaxLabelLength sets the maximum length for labels before truncation
func (c *PieChart) SetMaxLabelLength(length int) *PieChart {
	c.MaxLabelLength = length
	return c
}

// EnableTooltips enables or disables tooltips for labels
func (c *PieChart) EnableTooltips(enable bool) *PieChart {
	c.ShowTooltips = enable
	return c
}

// Render renders the line chart to an SVG string
func (c *LineChart) Render() string {
	var svg strings.Builder

	// Apply auto-height if enabled
	if c.AutoHeight {
		// For standard charts, use a 16:9 aspect ratio (common screen format)
		c.Height = c.Width * 9 / 16
	}

	// Start SVG with namespace
	svg.WriteString(fmt.Sprintf(`<svg width="100%%" height="auto" viewBox="0 0 %d %d" preserveAspectRatio="xMidYMid meet" xmlns="http://www.w3.org/2000/svg">`, c.Width, c.Height))

	// Add dark/light mode support if enabled
	if c.DarkModeSupport {
		svg.WriteString(fmt.Sprintf(`
			<style>
				:root {
					--chart-bg: %s;
					--chart-text: %s;
					--chart-axis: %s;
					--chart-grid: %s;
				}
				@media (prefers-color-scheme: dark) {
					:root {
						--chart-bg: %s;
						--chart-text: %s;
						--chart-axis: %s;
						--chart-grid: %s;
					}
				}
			</style>
		`,
			c.LightTheme.BackgroundColor,
			c.LightTheme.TextColor,
			c.LightTheme.AxisColor,
			c.LightTheme.GridColor,
			c.DarkTheme.BackgroundColor,
			c.DarkTheme.TextColor,
			c.DarkTheme.AxisColor,
			c.DarkTheme.GridColor))

		// Background with CSS variables
		svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="var(--chart-bg)"/>`, c.Width, c.Height))
	} else {
		// Background with static color
		svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="%s"/>`, c.Width, c.Height, c.BackgroundColor))
	}

	// Title
	if c.ShowTitle && c.Title != "" {
		if c.DarkModeSupport {
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="30" text-anchor="middle" font-family="Arial" font-size="20" font-weight="bold" fill="var(--chart-text)">%s</text>`,
				c.Width/2, c.Title))
		} else {
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="30" text-anchor="middle" font-family="Arial" font-size="20" font-weight="bold">%s</text>`,
				c.Width/2, c.Title))
		}
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

	// Draw axes with theme support
	if c.DarkModeSupport {
		svg.WriteString(fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="var(--chart-axis)" stroke-width="2"/>`,
			c.Margin.Left, c.Height-c.Margin.Bottom, c.Width-c.Margin.Right, c.Height-c.Margin.Bottom))
		svg.WriteString(fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="var(--chart-axis)" stroke-width="2"/>`,
			c.Margin.Left, c.Margin.Top, c.Margin.Left, c.Height-c.Margin.Bottom))
	} else {
		svg.WriteString(fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="2"/>`,
			c.Margin.Left, c.Height-c.Margin.Bottom, c.Width-c.Margin.Right, c.Height-c.Margin.Bottom))
		svg.WriteString(fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="2"/>`,
			c.Margin.Left, c.Margin.Top, c.Margin.Left, c.Height-c.Margin.Bottom))
	}

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
		path.WriteString(`" fill="none" stroke="` + c.Colors[0] + `" stroke-width="3"/>`)
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
					if c.DarkModeSupport {
						svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="12" fill="var(--chart-text)">%s</text>`,
							p[0], c.Height-c.Margin.Bottom+20, c.Labels[i]))
					} else {
						svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="12">%s</text>`,
							p[0], c.Height-c.Margin.Bottom+20, c.Labels[i]))
					}
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

	// Apply auto-height if enabled
	if c.AutoHeight {
		// For standard charts, use a 16:9 aspect ratio (common screen format)
		c.Height = c.Width * 9 / 16
	}

	// Start SVG with namespace
	svg.WriteString(fmt.Sprintf(`<svg width="100%%" height="auto" viewBox="0 0 %d %d" preserveAspectRatio="xMidYMid meet" xmlns="http://www.w3.org/2000/svg">`, c.Width, c.Height))

	// Add dark/light mode support if enabled
	if c.DarkModeSupport {
		svg.WriteString(fmt.Sprintf(`
			<style>
				:root {
					--chart-bg: %s;
					--chart-text: %s;
					--chart-axis: %s;
					--chart-grid: %s;
				}
				@media (prefers-color-scheme: dark) {
					:root {
						--chart-bg: %s;
						--chart-text: %s;
						--chart-axis: %s;
						--chart-grid: %s;
					}
				}
			</style>
		`,
			c.LightTheme.BackgroundColor,
			c.LightTheme.TextColor,
			c.LightTheme.AxisColor,
			c.LightTheme.GridColor,
			c.DarkTheme.BackgroundColor,
			c.DarkTheme.TextColor,
			c.DarkTheme.AxisColor,
			c.DarkTheme.GridColor))

		// Background with CSS variables
		svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="var(--chart-bg)"/>`, c.Width, c.Height))
	} else {
		// Background with static color
		svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="%s"/>`, c.Width, c.Height, c.BackgroundColor))
	}

	// Title
	if c.ShowTitle && c.Title != "" {
		if c.DarkModeSupport {
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="30" text-anchor="middle" font-family="Arial" font-size="20" font-weight="bold" fill="var(--chart-text)">%s</text>`,
				c.Width/2, c.Title))
		} else {
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="30" text-anchor="middle" font-family="Arial" font-size="20" font-weight="bold">%s</text>`,
				c.Width/2, c.Title))
		}
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

	// Draw axes with theme support
	if c.DarkModeSupport {
		svg.WriteString(fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="var(--chart-axis)" stroke-width="2"/>`,
			c.Margin.Left, c.Height-c.Margin.Bottom, c.Width-c.Margin.Right, c.Height-c.Margin.Bottom))
		svg.WriteString(fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="var(--chart-axis)" stroke-width="2"/>`,
			c.Margin.Left, c.Margin.Top, c.Margin.Left, c.Height-c.Margin.Bottom))
	} else {
		svg.WriteString(fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="2"/>`,
			c.Margin.Left, c.Height-c.Margin.Bottom, c.Width-c.Margin.Right, c.Height-c.Margin.Bottom))
		svg.WriteString(fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="2"/>`,
			c.Margin.Left, c.Margin.Top, c.Margin.Left, c.Height-c.Margin.Bottom))
	}

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

			// Add value text on top of bar with dark mode support
			if c.DarkModeSupport {
				svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="12" fill="var(--chart-text)">%.0f</text>`,
					barX+barWidth/2, barY-5, v))
			} else {
				svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="12" fill="black">%.0f</text>`,
					barX+barWidth/2, barY-5, v))
			}
		}

		// Draw labels if available
		if len(c.Labels) > 0 {
			for i := 0; i < len(c.Data); i++ {
				if i < len(c.Labels) {
					barX := c.Margin.Left + i*(chartWidth/len(c.Data)) + (chartWidth/len(c.Data))/2
					if c.DarkModeSupport {
						svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="12" fill="var(--chart-text)">%s</text>`,
							barX, c.Height-c.Margin.Bottom+20, c.Labels[i]))
					} else {
						svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="12">%s</text>`,
							barX, c.Height-c.Margin.Bottom+20, c.Labels[i]))
					}
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

	// Apply auto-height if enabled
	if c.AutoHeight {
		c.Height = c.Width // For pie charts, use a square aspect ratio
	}

	// Start SVG with namespace
	svg.WriteString(fmt.Sprintf(`<svg width="100%%" height="auto" viewBox="0 0 %d %d" preserveAspectRatio="xMidYMid meet" xmlns="http://www.w3.org/2000/svg">`, c.Width, c.Height))

	// Add dark/light mode support if enabled
	if c.DarkModeSupport {
		svg.WriteString(fmt.Sprintf(`
			<style>
				:root {
					--chart-bg: %s;
					--chart-text: %s;
					--chart-axis: %s;
					--chart-grid: %s;
				}
				@media (prefers-color-scheme: dark) {
					:root {
						--chart-bg: %s;
						--chart-text: %s;
						--chart-axis: %s;
						--chart-grid: %s;
					}
				}
			</style>
		`,
			c.LightTheme.BackgroundColor,
			c.LightTheme.TextColor,
			c.LightTheme.AxisColor,
			c.LightTheme.GridColor,
			c.DarkTheme.BackgroundColor,
			c.DarkTheme.TextColor,
			c.DarkTheme.AxisColor,
			c.DarkTheme.GridColor))

		// Background with CSS variables
		svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="var(--chart-bg)"/>`, c.Width, c.Height))
	} else {
		// Background with static color
		svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="%s"/>`, c.Width, c.Height, c.BackgroundColor))
	}

	// Title
	if c.ShowTitle && c.Title != "" {
		if c.DarkModeSupport {
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="30" text-anchor="middle" font-family="Arial" font-size="20" font-weight="bold" fill="var(--chart-text)">%s</text>`,
				c.Width/2, c.Title))
		} else {
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="30" text-anchor="middle" font-family="Arial" font-size="20" font-weight="bold">%s</text>`,
				c.Width/2, c.Title))
		}
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

			// Adjust label distance based on slice size
			// For smaller slices, move labels slightly outward
			var labelDistance float64
			if sliceAngle < math.Pi/6 { // Less than 30 degrees
				// Use a slightly larger distance for small slices
				labelDistance = float64(radius) * 0.6
			} else {
				// Standard distance for normal sized slices
				labelDistance = float64(radius) * 0.7
			}

			labelX := centerX + int(math.Cos(labelAngle)*labelDistance)
			labelY := centerY + int(math.Sin(labelAngle)*labelDistance)

			// Draw percentage
			percentage := v / total * 100

			// For very small slices, show tooltip but simpler label
			if sliceAngle < math.Pi/15 { // Less than 12 degrees
				if percentage < 5 {
					// For very small percentages, just show a simple dot with tooltip
					svg.WriteString(fmt.Sprintf(`<circle cx="%d" cy="%d" r="4" fill="white"><title>%.1f%%</title></circle>`,
						labelX, labelY, percentage))
				} else {
					// Show percentage with tooltip
					svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="10" fill="white">%.0f%%<title>%.1f%%</title></text>`,
						labelX, labelY, percentage, percentage))
				}
			} else {
				// Normal percentage text
				svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="12" fill="white">%.1f%%</text>`,
					labelX, labelY, percentage))
			}

			startAngle = endAngle
		}

		// Draw legend
		if c.ShowLegend && len(c.Labels) > 0 {
			// Position legend based on available space
			legendX := c.Width - c.Margin.Right + 20
			legendY := c.Margin.Top

			// Calculate total height needed for the legend
			legendHeight := len(c.Labels) * 25

			// Adjust legend position if it would go outside chart area
			if legendY+legendHeight > c.Height-c.Margin.Bottom {
				// Reduce spacing or move legend to better position if needed
				legendY = int(math.Max(float64(c.Margin.Top), float64(c.Height-c.Margin.Bottom-legendHeight)))
			}

			for i, label := range c.Labels {
				if i < len(c.Data) {
					colorIndex := i % len(c.Colors)
					color := c.Colors[colorIndex]

					// Truncate label if needed
					displayLabel := label
					if c.MaxLabelLength > 0 && len(label) > c.MaxLabelLength {
						displayLabel = label[:c.MaxLabelLength] + "…"
					}

					// Draw color box
					svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="15" height="15" fill="%s"/>`,
						legendX, legendY, color))

					// Draw label with tooltip if needed
					if c.ShowTooltips && displayLabel != label {
						// Add label with tooltip
						if c.DarkModeSupport {
							svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="12" fill="var(--chart-text)">%s<title>%s</title></text>`,
								legendX+20, legendY+12, displayLabel, label))
						} else {
							svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="12">%s<title>%s</title></text>`,
								legendX+20, legendY+12, displayLabel, label))
						}
					} else {
						// Regular label without tooltip
						if c.DarkModeSupport {
							svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="12" fill="var(--chart-text)">%s</text>`,
								legendX+20, legendY+12, displayLabel))
						} else {
							svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="12">%s</text>`,
								legendX+20, legendY+12, displayLabel))
						}
					}

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

	// Start SVG with namespace
	svg.WriteString(fmt.Sprintf(`<svg width="100%%" height="auto" viewBox="0 0 %d %d" preserveAspectRatio="xMidYMid meet" xmlns="http://www.w3.org/2000/svg">`, c.Width, c.Height))

	// Add dark/light mode support if enabled
	if c.DarkModeSupport {
		svg.WriteString(fmt.Sprintf(`
			<style>
				:root {
					--chart-bg: %s;
					--chart-text: %s;
					--chart-axis: %s;
					--chart-grid: %s;
				}
				@media (prefers-color-scheme: dark) {
					:root {
						--chart-bg: %s;
						--chart-text: %s;
						--chart-axis: %s;
						--chart-grid: %s;
					}
				}
			</style>
		`,
			c.LightTheme.BackgroundColor,
			c.LightTheme.TextColor,
			c.LightTheme.AxisColor,
			c.LightTheme.GridColor,
			c.DarkTheme.BackgroundColor,
			c.DarkTheme.TextColor,
			c.DarkTheme.AxisColor,
			c.DarkTheme.GridColor))

		// Background with CSS variables
		svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="var(--chart-bg)"/>`, c.Width, c.Height))
	} else {
		// Background with static color
		svg.WriteString(fmt.Sprintf(`<rect width="%d" height="%d" fill="%s"/>`, c.Width, c.Height, c.BackgroundColor))
	}

	// Title
	if c.ShowTitle && c.Title != "" {
		if c.DarkModeSupport {
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="30" text-anchor="middle" font-family="Arial" font-size="20" font-weight="bold" fill="var(--chart-text)">%s</text>`,
				c.Width/2, c.Title))
		} else {
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="30" text-anchor="middle" font-family="Arial" font-size="20" font-weight="bold">%s</text>`,
				c.Width/2, c.Title))
		}
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
	totalDays := int(lastSaturday.Sub(firstSunday).Hours()/24) + 1
	totalWeeks := totalDays / 7

	// Calculate cell size based on available space
	dayLabelWidth := 15 // Width for day labels
	availableWidth := c.Width - c.Margin.Left - c.Margin.Right - dayLabelWidth
	availableHeight := c.Height - c.Margin.Top - c.Margin.Bottom - 50 // 50px for title and month labels

	// Calculate cell size to fit within the available space
	// We need to fit totalWeeks columns and 7 rows
	maxCellWidth := (availableWidth - (totalWeeks-1)*c.CellSpacing) / totalWeeks
	maxCellHeight := (availableHeight - 6*c.CellSpacing) / 7 // 6 spaces between 7 rows

	// Use the smaller of the two to maintain square cells
	calculatedCellSize := math.Min(float64(maxCellWidth), float64(maxCellHeight))

	// Ensure minimum cell size is 3px
	cellSize := int(math.Max(3, calculatedCellSize))

	// If user specified a cell size explicitly, respect it unless it's too large to fit
	if c.CellSize > 0 {
		cellSize = int(math.Min(float64(c.CellSize), calculatedCellSize))
	}

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
	startX := c.Margin.Left + dayLabelWidth // Space for day labels
	startY := c.Margin.Top + 50             // Space for title and month labels

	// Draw day labels (using single-letter labels or user-defined labels)
	if len(c.DayLabels) == 7 {
		// Use user-defined labels if they've specified exactly 7 labels (one for each day)
		for i, label := range c.DayLabels {
			labelY := startY + i*(cellSize+c.CellSpacing) + cellSize/2 + 5
			if c.DarkModeSupport {
				svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="end" fill="var(--chart-text)">%s</text>`,
					startX-5, labelY, label))
			} else {
				svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="end">%s</text>`,
					startX-5, labelY, label))
			}
		}
	} else {
		// Default to single-letter day labels
		dayLetters := []string{"S", "M", "T", "W", "T", "F", "S"} // Sunday to Saturday
		for i, letter := range dayLetters {
			labelY := startY + i*(cellSize+c.CellSpacing) + cellSize/2 + 5
			if c.DarkModeSupport {
				svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="end" fill="var(--chart-text)">%s</text>`,
					startX-5, labelY, letter))
			} else {
				svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="end">%s</text>`,
					startX-5, labelY, letter))
			}
		}
	}

	// Draw the heatmap grid
	currentDate := firstSunday
	for week := 0; week < totalWeeks; week++ {
		// Check if we need to draw month label
		if currentDate.Day() <= 7 {
			// This is the first week of the month
			monthLabel := c.MonthLabels[currentDate.Month()-1]
			labelX := startX + week*(cellSize+c.CellSpacing) + cellSize/2
			if c.DarkModeSupport {
				svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="middle" fill="var(--chart-text)">%s</text>`,
					labelX, startY-5, monthLabel))
			} else {
				svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="middle">%s</text>`,
					labelX, startY-5, monthLabel))
			}
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
			cellX := startX + week*(cellSize+c.CellSpacing)
			cellY := startY + day*(cellSize+c.CellSpacing)

			// Draw cell
			svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" rx="%d" ry="%d" fill="%s">`,
				cellX, cellY, cellSize, cellSize, c.CellRounding, c.CellRounding, color))

			// Add tooltip
			svg.WriteString(fmt.Sprintf(`<title>%s: %v</title></rect>`, dateStr, value))

			// Move to next day
			currentDate = currentDate.AddDate(0, 0, 1)
		}
	}

	// Add legend
	if c.ShowLegend {
		legendX := c.Margin.Left
		legendY := startY + 7*(cellSize+c.CellSpacing) + 30
		legendLabelY := legendY + cellSize/2 + 5

		if c.DarkModeSupport {
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="start" fill="var(--chart-text)">Less</text>`,
				legendX, legendLabelY))
		} else {
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="start">Less</text>`,
				legendX, legendLabelY))
		}

		for i, color := range c.Colors {
			cellX := legendX + 40 + i*(cellSize+c.CellSpacing)
			svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" rx="%d" ry="%d" fill="%s"/>`,
				cellX, legendY, cellSize, cellSize, c.CellRounding, c.CellRounding, color))
		}

		if c.DarkModeSupport {
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="start" fill="var(--chart-text)">More</text>`,
				legendX+40+len(c.Colors)*(cellSize+c.CellSpacing)+5, legendLabelY))
		} else {
			svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" font-family="Arial" font-size="10" text-anchor="start">More</text>`,
				legendX+40+len(c.Colors)*(cellSize+c.CellSpacing)+5, legendLabelY))
		}
	}

	svg.WriteString("</svg>")
	return svg.String()
}

// EnableDarkModeSupport enables automatic adaptation between light and dark mode
// based on the user's system preferences
func (chart *BaseChart) EnableDarkModeSupport(enable bool) *BaseChart {
	chart.DarkModeSupport = enable

	// Set default values if they're not already set
	if chart.DarkTheme.BackgroundColor == "" {
		chart.DarkTheme.BackgroundColor = "#121212" // Dark background
		chart.DarkTheme.TextColor = "#ffffff"       // White text
		chart.DarkTheme.AxisColor = "#aaaaaa"       // Light gray axes
		chart.DarkTheme.GridColor = "#333333"       // Dark gray grid
	}

	if chart.LightTheme.BackgroundColor == "" {
		chart.LightTheme.BackgroundColor = chart.BackgroundColor // Use existing background
		if chart.LightTheme.BackgroundColor == "" {
			chart.LightTheme.BackgroundColor = "#ffffff" // White background
		}
		chart.LightTheme.TextColor = "#000000" // Black text
		chart.LightTheme.AxisColor = "#666666" // Dark gray axes
		chart.LightTheme.GridColor = "#dddddd" // Light gray grid
	}

	return chart
}

// SetDarkTheme sets the color scheme for dark mode
func (chart *BaseChart) SetDarkTheme(backgroundColor, textColor, axisColor, gridColor string) *BaseChart {
	chart.DarkTheme.BackgroundColor = backgroundColor
	chart.DarkTheme.TextColor = textColor
	chart.DarkTheme.AxisColor = axisColor
	chart.DarkTheme.GridColor = gridColor
	return chart
}

// SetLightTheme sets the color scheme for light mode
func (chart *BaseChart) SetLightTheme(backgroundColor, textColor, axisColor, gridColor string) *BaseChart {
	chart.LightTheme.BackgroundColor = backgroundColor
	chart.LightTheme.TextColor = textColor
	chart.LightTheme.AxisColor = axisColor
	chart.LightTheme.GridColor = gridColor
	return chart
}
