// This file can be run directly with:
// go run examples/heatmap_autosizing_example.go
package heatmap_example

import (
	"fmt"
	"os"
	"time"

	"github.com/riclib/gosvgchart"
)

// RunHeatmapExample demonstrates the autosizing heatmap with single-letter day labels
func RunHeatmapExample() {
	// Create a new heatmap chart
	heatmap := gosvgchart.NewHeatmapChart()

	// Generate sample data - 6 months of activity
	today := time.Now()
	startDate := today.AddDate(0, -6, 0)
	var labels []string
	var data []float64

	// Generate 180 days of sample data
	current := startDate
	for i := 0; i < 180; i++ {
		labels = append(labels, current.Format("2006-01-02"))

		// Generate some sample values - higher on weekends
		value := 1.0
		if current.Weekday() == time.Saturday || current.Weekday() == time.Sunday {
			value = 5.0
		}

		// Add some random patterns
		if i%7 == 0 {
			value = 10.0
		}
		if i%14 == 0 {
			value = 15.0
		}

		data = append(data, value)
		current = current.AddDate(0, 0, 1)
	}

	// Configure the heatmap
	heatmap.SetTitle("6 Month Activity Heatmap").
		SetSize(600, 250).
		SetData(data).
		SetLabels(labels).
		SetColors([]string{"#ebedf0", "#9be9a8", "#40c463", "#30a14e", "#216e39"})

	// The heatmap will automatically:
	// - Calculate cell size based on available space
	// - Display single-letter day labels (S, M, T, W, T, F, S)

	// Render the SVG
	svg := heatmap.Render()

	// Write to file
	err := os.WriteFile("heatmap_autosizing.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Println("Error writing SVG file:", err)
		return
	}

	fmt.Println("Heatmap saved to heatmap_autosizing.svg")

	// Create a smaller version to show adaptive sizing
	heatmapSmall := gosvgchart.NewHeatmapChart()
	heatmapSmall.SetTitle("Same Data, Smaller Size").
		SetSize(400, 200).
		SetData(data).
		SetLabels(labels).
		SetColors([]string{"#ebedf0", "#9be9a8", "#40c463", "#30a14e", "#216e39"})

	// Render the smaller SVG
	svgSmall := heatmapSmall.Render()

	// Write to file
	err = os.WriteFile("heatmap_autosizing_small.svg", []byte(svgSmall), 0644)
	if err != nil {
		fmt.Println("Error writing SVG file:", err)
		return
	}

	fmt.Println("Smaller heatmap saved to heatmap_autosizing_small.svg")
}
