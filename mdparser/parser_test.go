package mdparser

import (
	"strings"
	"testing"
)

func TestParseMarkdownChart(t *testing.T) {
	// Test Line Chart
	lineChartMD := `linechart
title: Test Line Chart
width: 600
height: 400
colors: #ff0000, #00ff00

data:
A | 10
B | 20
C | 15
D | 25`

	lineSVG, err := ParseMarkdownChart(lineChartMD)
	if err != nil {
		t.Errorf("Error parsing line chart: %v", err)
	}
	
	if !strings.Contains(lineSVG, "<svg") || !strings.Contains(lineSVG, "</svg>") {
		t.Error("Line chart SVG doesn't contain SVG tags")
	}
	
	if !strings.Contains(lineSVG, "Test Line Chart") {
		t.Error("Line chart doesn't contain title")
	}
	
	if !strings.Contains(lineSVG, "width=\"600\"") || !strings.Contains(lineSVG, "height=\"400\"") {
		t.Error("Line chart doesn't have correct dimensions")
	}
	
	// Test Bar Chart
	barChartMD := `barchart
title: Test Bar Chart
width: 500
height: 300
colors: #ff0000

data:
X | 100
Y | 200
Z | 150`

	barSVG, err := ParseMarkdownChart(barChartMD)
	if err != nil {
		t.Errorf("Error parsing bar chart: %v", err)
	}
	
	if !strings.Contains(barSVG, "<svg") || !strings.Contains(barSVG, "</svg>") {
		t.Error("Bar chart SVG doesn't contain SVG tags")
	}
	
	if !strings.Contains(barSVG, "Test Bar Chart") {
		t.Error("Bar chart doesn't contain title")
	}
	
	// Test Pie Chart
	pieChartMD := `piechart
title: Test Pie Chart
width: 400
height: 400
colors: #ff0000, #00ff00, #0000ff

data:
Slice1 | 30
Slice2 | 50
Slice3 | 20`

	pieSVG, err := ParseMarkdownChart(pieChartMD)
	if err != nil {
		t.Errorf("Error parsing pie chart: %v", err)
	}
	
	if !strings.Contains(pieSVG, "<svg") || !strings.Contains(pieSVG, "</svg>") {
		t.Error("Pie chart SVG doesn't contain SVG tags")
	}
	
	if !strings.Contains(pieSVG, "Test Pie Chart") {
		t.Error("Pie chart doesn't contain title")
	}
	
	// Test invalid chart type
	invalidChartMD := `invalid
title: Invalid Chart
data:
A | 10`

	_, err = ParseMarkdownChart(invalidChartMD)
	if err == nil {
		t.Error("Expected error for invalid chart type, but got none")
	}
	
	// Test too few lines
	tooFewLinesMD := `linechart`
	
	_, err = ParseMarkdownChart(tooFewLinesMD)
	if err == nil {
		t.Error("Expected error for too few lines, but got none")
	}
}