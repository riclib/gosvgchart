package main

import (
	"fmt"
	"os"
	"github.com/riclib/gosvgchart"
)

func main() {
	// Create a line chart
	lineChart := gosvgchart.New().
		SetTitle("Monthly Sales").
		SetData([]float64{120, 250, 180, 310, 270, 390}).
		SetLabels([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}).
		SetColors([]string{"#3498db"}).
		SetSize(600, 400)

	// Save the line chart to a file
	lineChartSVG := lineChart.Render()
	err := os.WriteFile("line_chart.svg", []byte(lineChartSVG), 0644)
	if err != nil {
		fmt.Printf("Error writing line chart: %v\n", err)
	}

	// Create a bar chart
	barChart := gosvgchart.New().
		SetTitle("Quarterly Revenue").
		SetData([]float64{850, 940, 1100, 1200}).
		SetLabels([]string{"Q1", "Q2", "Q3", "Q4"}).
		SetColors([]string{"#2ecc71", "#e74c3c", "#f39c12", "#9b59b6"}).
		SetSize(600, 400)

	// Save the bar chart to a file
	barChartSVG := barChart.Render()
	err = os.WriteFile("bar_chart.svg", []byte(barChartSVG), 0644)
	if err != nil {
		fmt.Printf("Error writing bar chart: %v\n", err)
	}

	// Create a pie chart
	pieChart := gosvgchart.New().
		SetTitle("Market Share").
		SetData([]float64{35, 25, 20, 15, 5}).
		SetLabels([]string{"Product A", "Product B", "Product C", "Product D", "Others"}).
		SetColors([]string{"#3498db", "#2ecc71", "#e74c3c", "#f39c12", "#9b59b6"}).
		SetSize(600, 500)

	// Save the pie chart to a file
	pieChartSVG := pieChart.Render()
	err = os.WriteFile("pie_chart.svg", []byte(pieChartSVG), 0644)
	if err != nil {
		fmt.Printf("Error writing pie chart: %v\n", err)
	}

	fmt.Println("Charts generated successfully!")
}