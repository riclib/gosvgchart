package mdparser

import (
	"fmt"
	"strconv"
	"strings"
	
	"github.com/riclib/gosvgchart"
)

// ParseMarkdownChart parses a chart specification from markdown text format
// and returns an SVG representation of the chart.
func ParseMarkdownChart(markdown string) (string, error) {
	lines := strings.Split(markdown, "\n")
	
	if len(lines) < 3 {
		return "", fmt.Errorf("invalid chart format: too few lines")
	}
	
	// Parse chart type from first line
	chartType := strings.TrimSpace(strings.ToLower(lines[0]))
	
	// Create appropriate chart
	var chart gosvgchart.Chart
	
	switch chartType {
	case "line", "linechart":
		chart = gosvgchart.New()
	case "bar", "barchart":
		chart = gosvgchart.NewBarChart()
	case "pie", "piechart":
		chart = gosvgchart.NewPieChart()
	default:
		return "", fmt.Errorf("unknown chart type: %s", chartType)
	}
	
	// Default settings
	width := 800
	height := 500
	title := "Chart"
	
	// Parse configuration and data
	var data []float64
	var labels []string
	var colors []string
	var dataStarted bool = false
	
	// Process line by line
	for i := 1; i < len(lines); i++ { // Skip first line (chart type)
		line := strings.TrimSpace(lines[i])
		
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		
		// Check for data section
		if line == "data:" {
			dataStarted = true
			continue
		}
		
		if dataStarted {
			// We're in the data section
			parts := strings.Split(line, "|")
			
			// If line contains a pipe, it has a label
			if len(parts) == 2 {
				label := strings.TrimSpace(parts[0])
				valueStr := strings.TrimSpace(parts[1])
				labels = append(labels, label)
				
				if val, err := strconv.ParseFloat(valueStr, 64); err == nil {
					data = append(data, val)
				}
			} else if len(parts) == 1 && parts[0] != "" {
				// No label, just data
				if val, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64); err == nil {
					data = append(data, val)
				}
			}
		} else {
			// We're in the configuration section
			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				continue // Invalid line format
			}
			
			key := strings.TrimSpace(strings.ToLower(parts[0]))
			value := strings.TrimSpace(parts[1])
			
			switch key {
			case "title":
				title = value
			case "width":
				if w, err := strconv.Atoi(value); err == nil && w > 0 {
					width = w
				}
			case "height":
				if h, err := strconv.Atoi(value); err == nil && h > 0 {
					height = h
				}
			case "colors":
				colors = parseList(value)
			}
		}
	}
	
	// Set basic properties
	chart.SetTitle(title)
	chart.SetSize(width, height)
	
	if len(colors) > 0 {
		chart.SetColors(colors)
	}
	
	// Set data and labels
	chart.SetData(data)
	if len(labels) > 0 {
		chart.SetLabels(labels)
	}
	
	// Apply chart-specific settings based on type
	switch chartType {
	case "line", "linechart":
		// No specific settings needed
	case "bar", "barchart":
		// No specific settings needed
	case "pie", "piechart":
		// No specific settings needed
	}
	
	// Render the chart
	return chart.Render(), nil
}

// parseList splits a comma-separated list and trims each element
func parseList(input string) []string {
	parts := strings.Split(input, ",")
	result := make([]string, 0, len(parts))
	
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	
	return result
}