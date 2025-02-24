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
		return "", fmt.Errorf("error: chart format invalid - too few lines. Need at least chart type, configuration, and data sections")
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
		return "", fmt.Errorf("error: unknown chart type '%s'. Must be one of: linechart, barchart, piechart", chartType)
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
	var foundDataSection bool = false
	var dataErrors []string
	var configErrors []string
	
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
			foundDataSection = true
			continue
		}
		
		if dataStarted {
			// We're in the data section
			parts := strings.Split(line, "|")
			
			// If line contains a pipe, it has a label
			if len(parts) == 2 {
				label := strings.TrimSpace(parts[0])
				valueStr := strings.TrimSpace(parts[1])
				
				if label == "" {
					dataErrors = append(dataErrors, fmt.Sprintf("line %d: missing label before '|'", i+1))
				}
				
				labels = append(labels, label)
				
				if val, err := strconv.ParseFloat(valueStr, 64); err == nil {
					data = append(data, val)
				} else {
					dataErrors = append(dataErrors, fmt.Sprintf("line %d: '%s' is not a valid number", i+1, valueStr))
				}
			} else if len(parts) == 1 && parts[0] != "" {
				// No label, just data
				if val, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64); err == nil {
					data = append(data, val)
				} else {
					dataErrors = append(dataErrors, fmt.Sprintf("line %d: '%s' is not a valid number", i+1, parts[0]))
				}
			} else {
				dataErrors = append(dataErrors, fmt.Sprintf("line %d: invalid data format, expected 'label | value'", i+1))
			}
		} else {
			// We're in the configuration section
			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				configErrors = append(configErrors, fmt.Sprintf("line %d: invalid configuration format, expected 'key: value'", i+1))
				continue
			}
			
			key := strings.TrimSpace(strings.ToLower(parts[0]))
			value := strings.TrimSpace(parts[1])
			
			switch key {
			case "title":
				title = value
			case "width":
				if w, err := strconv.Atoi(value); err == nil && w > 0 {
					width = w
				} else {
					configErrors = append(configErrors, fmt.Sprintf("line %d: invalid width value '%s' - must be a positive number", i+1, value))
				}
			case "height":
				if h, err := strconv.Atoi(value); err == nil && h > 0 {
					height = h
				} else {
					configErrors = append(configErrors, fmt.Sprintf("line %d: invalid height value '%s' - must be a positive number", i+1, value))
				}
			case "colors":
				colorList := parseList(value)
				if len(colorList) == 0 {
					configErrors = append(configErrors, fmt.Sprintf("line %d: invalid colors format - must be comma-separated list", i+1))
				} else {
					colors = colorList
				}
			default:
				configErrors = append(configErrors, fmt.Sprintf("line %d: unknown configuration key '%s'", i+1, key))
			}
		}
	}
	
	// Required validation
	var errors []string
	errors = append(errors, configErrors...)
	errors = append(errors, dataErrors...)
	
	if !foundDataSection {
		errors = append(errors, "missing 'data:' section - chart must include a data section")
	}
	
	if len(data) == 0 {
		errors = append(errors, "no valid data points found - chart requires at least one data point")
	}
	
	if len(labels) > 0 && len(labels) != len(data) {
		errors = append(errors, fmt.Sprintf("mismatched labels and data points - found %d labels but %d data points", len(labels), len(data)))
	}
	
	// If we have validation errors, return them
	if len(errors) > 0 {
		return "", fmt.Errorf("chart definition errors:\n• %s", strings.Join(errors, "\n• "))
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