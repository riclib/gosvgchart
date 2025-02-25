package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/riclib/gosvgchart"
)

func main() {
	// Parse command-line flags
	pieOnly := flag.Bool("pie", false, "Show only the pie chart example")
	flag.Parse()

	// Full examples (all chart types) or just pie chart based on the flag
	if *pieOnly {
		generatePieChartExample()
		fmt.Println("Pie chart example generated: pie_chart_example.svg")
	} else {
		generateAllExamples()
	}
}

func generateAllExamples() {
	// Generate line chart
	lineChart := gosvgchart.NewLineChart().
		SetTitle("Monthly Sales").
		SetData([]float64{120, 250, 180, 310, 270, 390}).
		SetLabels([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}).
		SetColors([]string{"#3498db"}).
		SetSize(800, 400)

	// Save line chart
	os.WriteFile("line_chart_example.svg", []byte(lineChart.Render()), 0644)
	fmt.Println("Line chart example generated: line_chart_example.svg")

	// Generate bar chart
	barChart := gosvgchart.NewBarChart().
		SetTitle("Quarterly Revenue").
		SetData([]float64{850, 940, 1100, 1200}).
		SetLabels([]string{"Q1", "Q2", "Q3", "Q4"}).
		SetColors([]string{"#2ecc71", "#e74c3c", "#f39c12", "#9b59b6"}).
		SetSize(800, 400)

	// Save bar chart
	os.WriteFile("bar_chart_example.svg", []byte(barChart.Render()), 0644)
	fmt.Println("Bar chart example generated: bar_chart_example.svg")

	// Generate pie chart
	pieChart := gosvgchart.NewPieChart().
		SetTitle("Market Share").
		SetData([]float64{35, 25, 20, 15, 5}).
		SetLabels([]string{"Product A", "Product B", "Product C", "Product D", "Others"}).
		SetColors([]string{"#3498db", "#2ecc71", "#e74c3c", "#f39c12", "#9b59b6"}).
		SetSize(600, 500)

	// Save pie chart
	os.WriteFile("pie_chart_example.svg", []byte(pieChart.Render()), 0644)
	fmt.Println("Pie chart example generated: pie_chart_example.svg")

	// Generate example with very long labels
	generatePieChartWithLongLabels()
}

func generatePieChartExample() {
	// Generate pie chart with default options
	pieChart := gosvgchart.NewPieChart().
		SetTitle("Market Share").
		SetData([]float64{35, 25, 20, 15, 5}).
		SetLabels([]string{"Product A", "Product B", "Product C", "Product D", "Others"}).
		SetColors([]string{"#3498db", "#2ecc71", "#e74c3c", "#f39c12", "#9b59b6"}).
		SetSize(600, 500)

	// Save pie chart
	os.WriteFile("pie_chart_example.svg", []byte(pieChart.Render()), 0644)
}

// This function generates a pie chart example with very long labels
// to test our label truncation and tooltip features
func generatePieChartWithLongLabels() {
	// Generate pie chart with long labels
	pc := gosvgchart.NewPieChart()
	pieChart := pc.
		SetTitle("Market Share with Long Labels").
		SetData([]float64{35, 25, 20, 15, 3, 2}).
		SetLabels([]string{
			"Very Long Product Name A (Enterprise Edition)",
			"Extremely Long Product Name B (Professional Plus)",
			"Super Long Product Name C (Deluxe Version 2.0)",
			"Extremely Long Product Name D (Premium Version)",
			"Very Small Value Product",
			"Tiny Value",
		}).
		SetColors([]string{"#3498db", "#2ecc71", "#e74c3c", "#f39c12", "#9b59b6", "#1abc9c"}).
		SetSize(600, 500)
		
	// Type-specific settings that aren't part of the Chart interface
	pc.SetMaxLabelLength(15)  // Truncate labels to 15 characters
	pc.EnableTooltips(true)   // Show tooltips for truncated labels

	// Save pie chart
	os.WriteFile("pie_chart_long_labels.svg", []byte(pieChart.Render()), 0644)
	fmt.Println("Pie chart with long labels example generated: pie_chart_long_labels.svg")
}