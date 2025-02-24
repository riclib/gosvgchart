package mdparser

import (
	"bufio"
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
		chart = gosvgchart.New()
	case "pie", "piechart":
		chart = gosvgchart.New()
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
	
	scanner := bufio.NewScanner(strings.NewReader(markdown))
	scanner.Scan() // Skip first line (chart type)
	
	// Parse configuration lines
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		
		// Break at data section marker
		if line == "data:" {
			break
		}
		
		// Parse configuration options
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
	
	// Set basic properties
	chart.SetTitle(title)
	chart.SetSize(width, height)
	
	if len(colors) > 0 {
		chart.SetColors(colors)
	}
	
	// Parse data section
	dataStarted := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		
		if line == "data:" {
			dataStarted = true
			continue
		}
		
		if dataStarted {
			parts := strings.SplitN(line, "|", 2)
			
			// If line contains a pipe, it has a label
			if len(parts) == 2 {
				label := strings.TrimSpace(parts[0])
				labels = append(labels, label)
				
				if val, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64); err == nil {
					data = append(data, val)
				}
			} else if len(parts) == 1 {
				// No label, just data
				if val, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64); err == nil {
					data = append(data, val)
				}
			}
		}
	}
	
	// Set data and labels
	chart.SetData(data)
	if len(labels) > 0 {
		chart.SetLabels(labels)
	}
	
	// Apply chart-specific settings based on type
	switch chartType {
	case "line", "linechart":
		if lc, ok := chart.(*gosvgchart.LineChart); ok {
			// Apply line-specific settings here if needed
			lc.ShowDataPoints(true)
		}
	case "bar", "barchart":
		if bc, ok := chart.(*gosvgchart.BarChart); ok {
			// Apply bar-specific settings here if needed
		}
	case "pie", "piechart":
		if pc, ok := chart.(*gosvgchart.PieChart); ok {
			// Apply pie-specific settings here if needed
		}
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