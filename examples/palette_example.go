package main

import (
	"os"

	"github.com/riclib/gosvgchart"
)

func main() {
	generateAutoPaletteExamples()
	generateGradientPaletteExamples()
}

func generateAutoPaletteExamples() {
	// Create a bar chart with auto palette
	barData := []float64{120, 250, 180, 310, 270, 390}
	barLabels := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}

	barChart := gosvgchart.NewBarChart().
		SetTitle("Monthly Sales (Auto Palette)").
		SetSize(800, 400).
		SetData(barData).
		SetLabels(barLabels).
		SetPalette("auto")

	// Render to SVG and save to file
	svg := barChart.Render()
	os.WriteFile("auto_palette_bar_chart.svg", []byte(svg), 0644)

	// Create a pie chart with auto palette
	pieData := []float64{35, 25, 20, 15, 5}
	pieLabels := []string{"Product A", "Product B", "Product C", "Product D", "Others"}

	pieChart := gosvgchart.NewPieChart().
		SetTitle("Market Share (Auto Palette)").
		SetSize(500, 500).
		SetData(pieData).
		SetLabels(pieLabels).
		SetPalette("auto")

	// Render to SVG and save to file
	svg = pieChart.Render()
	os.WriteFile("auto_palette_pie_chart.svg", []byte(svg), 0644)

	// Create a multiple series line chart with auto palette
	lineChart := gosvgchart.NewLineChart().
		SetTitle("Monthly Sales by Product (Auto Palette)").
		SetSize(800, 400).
		SetLabels([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}).
		SetPalette("auto")

	// Add multiple series
	lineChart.AddSeries("Product A", []float64{120, 150, 180, 210, 240, 270})
	lineChart.AddSeries("Product B", []float64{200, 180, 160, 140, 120, 100})
	lineChart.AddSeries("Product C", []float64{50, 80, 110, 140, 170, 200})

	// Enable legend
	lineChart.ShowLegend = true
	lineChart.SetLegendWidth(0.2)

	// Render to SVG and save to file
	svg = lineChart.Render()
	os.WriteFile("auto_palette_line_chart.svg", []byte(svg), 0644)
}

func generateGradientPaletteExamples() {
	// Create a bar chart with gradient palette
	barData := []float64{120, 250, 180, 310, 270, 390}
	barLabels := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}

	barChart := gosvgchart.NewBarChart().
		SetTitle("Monthly Sales (Gradient Palette)").
		SetSize(800, 400).
		SetData(barData).
		SetLabels(barLabels).
		SetPalette("gradient")

	// Render to SVG and save to file
	svg := barChart.Render()
	os.WriteFile("gradient_palette_bar_chart.svg", []byte(svg), 0644)

	// Create a pie chart with gradient palette
	pieData := []float64{35, 25, 20, 15, 5}
	pieLabels := []string{"Product A", "Product B", "Product C", "Product D", "Others"}

	pieChart := gosvgchart.NewPieChart().
		SetTitle("Market Share (Gradient Palette)").
		SetSize(500, 500).
		SetData(pieData).
		SetLabels(pieLabels).
		SetPalette("gradient")

	// Render to SVG and save to file
	svg = pieChart.Render()
	os.WriteFile("gradient_palette_pie_chart.svg", []byte(svg), 0644)

	// Create a multiple series line chart with gradient palette
	lineChart := gosvgchart.NewLineChart().
		SetTitle("Monthly Sales by Product (Gradient Palette)").
		SetSize(800, 400).
		SetLabels([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}).
		SetPalette("gradient")

	// Add multiple series
	lineChart.AddSeries("Product A", []float64{120, 150, 180, 210, 240, 270})
	lineChart.AddSeries("Product B", []float64{200, 180, 160, 140, 120, 100})
	lineChart.AddSeries("Product C", []float64{50, 80, 110, 140, 170, 200})

	// Enable legend
	lineChart.ShowLegend = true
	lineChart.SetLegendWidth(0.2)

	// Render to SVG and save to file
	svg = lineChart.Render()
	os.WriteFile("gradient_palette_line_chart.svg", []byte(svg), 0644)
}