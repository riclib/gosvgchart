package mdparser

import (
	"fmt"
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
	
	// Test Auto-Height
	autoHeightMD := `barchart
title: Auto Height Chart
width: 600
height: auto
colors: #00ff00

data:
X | 100
Y | 200
Z | 150`

	autoHeightSVG, err := ParseMarkdownChart(autoHeightMD)
	if err != nil {
		t.Errorf("Error parsing auto-height chart: %v", err)
	}
	
	if !strings.Contains(autoHeightSVG, "<svg") || !strings.Contains(autoHeightSVG, "</svg>") {
		t.Error("Auto-height chart SVG doesn't contain SVG tags")
	}
	
	if !strings.Contains(autoHeightSVG, "Auto Height Chart") {
		t.Error("Auto-height chart doesn't contain title")
	}
	
	// For bar chart with auto-height, height should be width * 9 / 16 = 600 * 9 / 16 = 337.5 ~ 337
	expectedHeight := 337
	if !strings.Contains(autoHeightSVG, fmt.Sprintf("height=\"%d\"", expectedHeight)) {
		t.Errorf("Auto-height chart doesn't have correct auto-calculated height: %s", autoHeightSVG)
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
	} else if !strings.Contains(err.Error(), "unknown chart type") {
		t.Errorf("Expected error about unknown chart type, got: %v", err)
	}
	
	// Test too few lines
	tooFewLinesMD := `linechart`
	
	_, err = ParseMarkdownChart(tooFewLinesMD)
	if err == nil {
		t.Error("Expected error for too few lines, but got none")
	} else if !strings.Contains(err.Error(), "too few lines") {
		t.Errorf("Expected error about too few lines, got: %v", err)
	}
	
	// Test invalid width value
	invalidWidthMD := `linechart
title: Test Chart
width: notanumber
height: 400

data:
A | 10`

	_, err = ParseMarkdownChart(invalidWidthMD)
	if err == nil {
		t.Error("Expected error for invalid width, but got none")
	} else if !strings.Contains(err.Error(), "invalid width value") {
		t.Errorf("Expected error about invalid width, got: %v", err)
	}
	
	// Test invalid height value
	invalidHeightMD := `linechart
title: Test Chart
width: 600
height: -400

data:
A | 10`

	_, err = ParseMarkdownChart(invalidHeightMD)
	if err == nil {
		t.Error("Expected error for invalid height, but got none")
	} else if !strings.Contains(err.Error(), "invalid height value") {
		t.Errorf("Expected error about invalid height, got: %v", err)
	}
	
	// Test invalid data format
	invalidDataMD := `linechart
title: Test Chart
width: 600
height: 400

data:
A | notanumber`

	_, err = ParseMarkdownChart(invalidDataMD)
	if err == nil {
		t.Error("Expected error for invalid data, but got none")
	} else if !strings.Contains(err.Error(), "not a valid number") {
		t.Errorf("Expected error about invalid number, got: %v", err)
	}
	
	// Test missing data section
	noDataSectionMD := `linechart
title: Test Chart
width: 600
height: 400`

	_, err = ParseMarkdownChart(noDataSectionMD)
	if err == nil {
		t.Error("Expected error for missing data section, but got none")
	} else if !strings.Contains(err.Error(), "missing 'data:' section") {
		t.Errorf("Expected error about missing data section, got: %v", err)
	}
	
	// Test empty data section
	emptyDataMD := `linechart
title: Test Chart
width: 600
height: 400

data:`

	_, err = ParseMarkdownChart(emptyDataMD)
	if err == nil {
		t.Error("Expected error for empty data section, but got none")
	} else if !strings.Contains(err.Error(), "no valid data points found") {
		t.Errorf("Expected error about no data points, got: %v", err)
	}
	
	// Test mismatched labels and data
	mismatchedLabelsMD := `linechart
title: Test Chart
width: 600
height: 400

data:
A | 10
B | 20
C |
D | 40`

	_, err = ParseMarkdownChart(mismatchedLabelsMD)
	if err == nil {
		t.Error("Expected error for mismatched labels and data, but got none")
	} else if !strings.Contains(err.Error(), "mismatched labels and data points") || 
		!strings.Contains(err.Error(), "not a valid number") {
		t.Errorf("Expected error about mismatched data points, got: %v", err)
	}
	
	// Test multiple errors
	multipleErrorsMD := `linechart
title: Test Chart
width: badwidth
unknownkey: value

data:
A | 10
B | invalid
C | 30`

	_, err = ParseMarkdownChart(multipleErrorsMD)
	if err == nil {
		t.Error("Expected multiple errors, but got none")
	} else {
		errMsg := err.Error()
		if !strings.Contains(errMsg, "invalid width value") {
			t.Error("Expected error about invalid width")
		}
		if !strings.Contains(errMsg, "unknown configuration key") {
			t.Error("Expected error about unknown configuration key")
		}
		if !strings.Contains(errMsg, "not a valid number") {
			t.Error("Expected error about invalid number")
		}
		// Check bullet point format
		if !strings.Contains(errMsg, "• ") {
			t.Error("Expected bullet point format for multiple errors")
		}
	}
	
	// Test multiple charts in single code block
	multipleChartsMD := `linechart
title: First Chart
width: 400
height: 300
colors: #ff0000

data:
A | 10
B | 20

---

barchart
title: Second Chart
width: 400
height: 300
colors: #00ff00

data:
X | 30
Y | 40`

	multiChartSVG, err := ParseMarkdownChart(multipleChartsMD)
	if err != nil {
		t.Errorf("Error parsing multiple charts: %v", err)
	}
	
	if !strings.Contains(multiChartSVG, "display: flex") {
		t.Error("Multiple charts should be wrapped in a flex container")
	}
	
	if !strings.Contains(multiChartSVG, "First Chart") || !strings.Contains(multiChartSVG, "Second Chart") {
		t.Error("Multiple charts SVG doesn't contain both chart titles")
	}
	
	// Count SVG tags to ensure both charts are rendered
	svgCount := strings.Count(multiChartSVG, "<svg")
	if svgCount != 2 {
		t.Errorf("Expected 2 SVG elements, got %d", svgCount)
	}
}