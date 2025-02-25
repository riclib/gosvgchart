package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/riclib/gosvgchart"
)

func main() {
	// Parse command line flags
	exampleType := flag.String("type", "heatmap", "Example type: heatmap, line, bar, pie")
	flag.Parse()

	switch *exampleType {
	case "heatmap":
		runHeatmapExample()
	case "line":
		runLineExample()
	case "bar":
		runBarExample()
	case "pie":
		runPieExample()
	default:
		fmt.Printf("Unknown example type: %s\n", *exampleType)
		fmt.Println("Available examples: heatmap, line, bar, pie")
	}
}

// runHeatmapExample demonstrates the autosizing heatmap with single-letter day labels
func runHeatmapExample() {
	// Create a new heatmap chart
	heatmap := gosvgchart.NewHeatmapChart()

	// Generate sample data - 6 months of activity
	today := time.Now()
	startDate := today.AddDate(0, -6, 0)
	var labels []string
	var data []float64

	// Generate data for 180 days
	for i := 0; i < 180; i++ {
		date := startDate.AddDate(0, 0, i)
		labels = append(labels, date.Format("2006-01-02"))

		// Generate random-like data with higher values on weekends
		val := float64(i % 10)
		if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
			val += 10 // Higher activity on weekends
		}
		data = append(data, val)
	}

	// Configure the heatmap
	heatmap.SetTitle("Activity Heatmap (Autosizing)")
	heatmap.SetSize(800, 300)
	heatmap.SetData(data)
	heatmap.SetLabels(labels)
	heatmap.SetColors([]string{"#ebedf0", "#9be9a8", "#40c463", "#30a14e", "#216e39"})

	// Render and save the SVG
	svg := heatmap.Render()
	err := os.WriteFile("heatmap_autosizing.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	fmt.Println("Heatmap saved to heatmap_autosizing.svg")

	// Create a smaller version to demonstrate adaptivity
	heatmap.SetSize(400, 200)
	svg = heatmap.Render()
	err = os.WriteFile("heatmap_autosizing_small.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	fmt.Println("Small heatmap saved to heatmap_autosizing_small.svg")
}

// runLineExample demonstrates a line chart with dark mode support
func runLineExample() {
	// Create a new line chart with specific type
	lineChart := gosvgchart.NewLineChart()

	// Configure the line chart
	lineChart.SetTitle("Monthly Sales (Dark Mode Compatible)")
	lineChart.SetSize(800, 500)
	lineChart.SetData([]float64{120, 250, 180, 310, 270, 390, 210, 380, 330, 400, 280, 290})
	lineChart.SetLabels([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"})
	lineChart.SetColors([]string{"#3498db"})

	// Line chart specific settings
	lineChart.SetSmooth(true)
	lineChart.ShowDataPoints(true)

	// Render and save the SVG
	svg := lineChart.Render()
	err := os.WriteFile("line_chart_example.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	fmt.Println("Line chart saved to line_chart_example.svg")
}

// runBarExample demonstrates a bar chart with dark mode support
func runBarExample() {
	// Create a new bar chart
	barChart := gosvgchart.NewBarChart()

	// Configure the bar chart
	barChart.SetTitle("Quarterly Revenue (Dark Mode Compatible)")
	barChart.SetSize(800, 500)
	barChart.SetData([]float64{850, 940, 1100, 1200})
	barChart.SetLabels([]string{"Q1", "Q2", "Q3", "Q4"})
	barChart.SetColors([]string{"#2ecc71", "#e74c3c", "#f39c12", "#9b59b6"})

	// Render and save the SVG
	svg := barChart.Render()
	err := os.WriteFile("bar_chart_example.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	fmt.Println("Bar chart saved to bar_chart_example.svg")
}

// runPieExample demonstrates a pie chart with dark mode support
func runPieExample() {
	// Create a new pie chart
	pieChart := gosvgchart.NewPieChart()

	// Configure the pie chart
	pieChart.SetTitle("Market Share (Dark Mode Compatible)")
	pieChart.SetSize(800, 500)
	pieChart.SetData([]float64{35, 25, 20, 15, 5})
	pieChart.SetLabels([]string{"Product A", "Product B", "Product C", "Product D", "Others"})
	pieChart.SetColors([]string{"#3498db", "#2ecc71", "#e74c3c", "#f39c12", "#9b59b6"})

	// Pie chart specific settings
	pieChart.SetDonutHole(0.4) // Create a donut chart

	// Render and save the SVG
	svg := pieChart.Render()
	err := os.WriteFile("pie_chart_example.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	fmt.Println("Pie chart saved to pie_chart_example.svg")
}
