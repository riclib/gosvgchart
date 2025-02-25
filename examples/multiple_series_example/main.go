// Example of multiple series support in GoSVGChart
package main

import (
	"fmt"
	"os"

	"github.com/riclib/gosvgchart"
)

func main() {
	// Create a line chart with multiple series
	lineChart := gosvgchart.NewLineChart()
	lineChart.SetTitle("Monthly Sales by Product").
		SetSize(800, 500).
		SetLabels([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"})

	// Add multiple series
	lineChart.AddSeries("Product A", []float64{120, 150, 180, 210, 240, 270})
	lineChart.AddSeries("Product B", []float64{200, 180, 160, 140, 120, 100})
	lineChart.AddSeries("Product C", []float64{50, 80, 110, 140, 170, 200})

	// Set colors for the series
	lineChart.SetSeriesColors([]string{"#4285F4", "#EA4335", "#FBBC05", "#34A853"})

	// Enable legend
	lineChart.ShowLegend = true

	// Render the line chart
	lineSvg := lineChart.Render()

	// Save the line chart to a file
	err := os.WriteFile("line_chart_multiple_series.svg", []byte(lineSvg), 0644)
	if err != nil {
		fmt.Printf("Error saving line chart: %v\n", err)
	} else {
		fmt.Println("Line chart with multiple series saved to line_chart_multiple_series.svg")
	}

	// Create a grouped bar chart with multiple series
	groupedBarChart := gosvgchart.NewBarChart()
	groupedBarChart.SetTitle("Quarterly Revenue by Region").
		SetSize(800, 500).
		SetLabels([]string{"Q1", "Q2", "Q3", "Q4"})

	// Add multiple series
	groupedBarChart.AddSeries("North", []float64{150, 180, 210, 240})
	groupedBarChart.AddSeries("South", []float64{120, 140, 160, 180})
	groupedBarChart.AddSeries("East", []float64{90, 110, 130, 150})
	groupedBarChart.AddSeries("West", []float64{180, 200, 220, 240})

	// Set colors for the series
	groupedBarChart.SetSeriesColors([]string{"#4285F4", "#EA4335", "#FBBC05", "#34A853"})

	// Enable legend
	groupedBarChart.ShowLegend = true

	// Render the grouped bar chart
	groupedBarSvg := groupedBarChart.Render()

	// Save the grouped bar chart to a file
	err = os.WriteFile("grouped_bar_chart.svg", []byte(groupedBarSvg), 0644)
	if err != nil {
		fmt.Printf("Error saving grouped bar chart: %v\n", err)
	} else {
		fmt.Println("Grouped bar chart saved to grouped_bar_chart.svg")
	}

	// Create a stacked bar chart with multiple series
	stackedBarChart := gosvgchart.NewBarChart()
	stackedBarChart.SetTitle("Quarterly Revenue by Region (Stacked)").
		SetSize(800, 500).
		SetLabels([]string{"Q1", "Q2", "Q3", "Q4"})

	// Set stacked property
	stackedBarChart.Stacked = true

	// Add multiple series
	stackedBarChart.AddSeries("North", []float64{150, 180, 210, 240})
	stackedBarChart.AddSeries("South", []float64{120, 140, 160, 180})
	stackedBarChart.AddSeries("East", []float64{90, 110, 130, 150})
	stackedBarChart.AddSeries("West", []float64{180, 200, 220, 240})

	// Set colors for the series
	stackedBarChart.SetSeriesColors([]string{"#4285F4", "#EA4335", "#FBBC05", "#34A853"})

	// Enable legend
	stackedBarChart.ShowLegend = true

	// Render the stacked bar chart
	stackedBarSvg := stackedBarChart.Render()

	// Save the stacked bar chart to a file
	err = os.WriteFile("stacked_bar_chart.svg", []byte(stackedBarSvg), 0644)
	if err != nil {
		fmt.Printf("Error saving stacked bar chart: %v\n", err)
	} else {
		fmt.Println("Stacked bar chart saved to stacked_bar_chart.svg")
	}
}
