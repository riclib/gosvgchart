package main

import (
	"fmt"
	"os"

	"github.com/riclib/gosvgchart"
)

func main() {
	// Create a new line chart (dark mode is enabled by default)
	lineChart := gosvgchart.NewLineChart()

	// Configure the line chart
	lineChart.SetTitle("Monthly Sales (Dark Mode Compatible)").
		SetSize(800, 500).
		SetData([]float64{120, 250, 180, 310, 270, 390, 320, 410, 450, 380, 330, 420}).
		SetLabels([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}).
		SetColors([]string{"#3498db", "#e74c3c"})

	// Enable line-specific features
	lineChart.ShowDataPoints(true)

	// Note: Dark mode is already enabled by default,
	// but we can customize the theme colors if desired

	// Customize dark theme
	lineChart.SetDarkTheme(
		"#1e1e2e", // Background - dark blue/gray
		"#cdd6f4", // Text - light cream
		"#89b4fa", // Axis - light blue
		"#45475a", // Grid - medium gray
	)

	// Customize light theme
	lineChart.SetLightTheme(
		"#f8f9fa", // Background - off white
		"#11191f", // Text - very dark blue/gray
		"#3498db", // Axis - blue
		"#e2e8f0", // Grid - light gray
	)

	// Render the SVG
	svg := lineChart.Render()

	// Write to file
	err := os.WriteFile("dark_mode_chart.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Println("Error writing SVG file:", err)
		return
	}

	fmt.Println("Chart saved to dark_mode_chart.svg")
	fmt.Println("Open in a browser to see it adapt based on your system's color scheme!")
}
