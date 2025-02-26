// Example showing a heatmap with positive and negative values for feedback
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/riclib/gosvgchart"
)

func main() {
	// Create a new heatmap chart for feedback data
	feedbackHeatmap := gosvgchart.NewHeatmapChart()

	// Generate sample data - 6 months of feedback ratings (-1 to +1)
	today := time.Now()
	startDate := today.AddDate(0, -6, 0)
	var labels []string
	var data []float64

	// Generate 180 days of sample feedback data
	current := startDate
	for i := 0; i < 180; i++ {
		labels = append(labels, current.Format("2006-01-02"))

		// Generate feedback values between -1 and +1
		// This pattern creates alternating positive and negative sections
		// with varying intensities for demo purposes
		var value float64
		switch {
		case i%30 < 10:
			// First third of each month: positive feedback
			value = 0.5 + float64(i%10)*0.05 // Values from 0.5 to 0.95
		case i%30 < 20:
			// Second third: negative feedback
			value = -0.3 - float64(i%10)*0.07 // Values from -0.3 to -0.93
		default:
			// Last third: mixed feedback
			if i%2 == 0 {
				value = 0.2 + float64(i%5)*0.15 // Positive values
			} else {
				value = -0.1 - float64(i%5)*0.18 // Negative values
			}
		}

		data = append(data, value)
		current = current.AddDate(0, 0, 1)
	}

	// Configure the heatmap
	feedbackHeatmap.SetTitle("6 Month Feedback Heatmap").
		SetSize(800, 250).
		SetData(data).
		SetLabels(labels).
		// Green gradient for positive feedback
		SetColors([]string{"#ebedf0", "#9be9a8", "#40c463", "#30a14e", "#216e39"})
		
	// Red gradient for negative feedback
	feedbackHeatmap.SetNegativeColors([]string{"#ebedf0", "#f9a8a8", "#f67575", "#e64545", "#c92a2a"})
	
	// Enable support for negative values (already enabled by default)
	feedbackHeatmap.EnableNegativeValues(true)

	// Render the SVG
	svg := feedbackHeatmap.Render()

	// Write to file
	err := os.WriteFile("feedback_heatmap.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Println("Error writing SVG file:", err)
		return
	}

	fmt.Println("Feedback heatmap saved to feedback_heatmap.svg")
}