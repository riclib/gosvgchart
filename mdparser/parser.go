package mdparser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	
	"github.com/riclib/gosvgchart"
)

// ChartDefinition represents a parsed chart definition
type ChartDefinition struct {
	ChartType string
	Title     string
	Width     int
	Height    int
	Colors    []string
	Data      []float64
	Labels    []string
	AutoHeight bool
}

// ParseMarkdownChart parses a chart specification from markdown text format
// and returns an SVG representation of the chart.
func ParseMarkdownChart(markdown string) (string, error) {
	// Look for chart separator, which is a line containing only "---" with optional whitespace
	re := regexp.MustCompile(`(?m)^\s*---\s*$`)
	
	// Check if we have multiple charts
	if re.MatchString(markdown) {
		// We have multiple charts, split and render them side by side
		chartBlocks := re.Split(markdown, -1)
		
		// Trim each chart block
		for i := range chartBlocks {
			chartBlocks[i] = strings.TrimSpace(chartBlocks[i])
		}
		
		return parseMultipleCharts(chartBlocks)
	}
	
	// Single chart
	chartDef, err := parseChartDefinition(markdown, 0)
	if err != nil {
		return "", err
	}
	
	return renderChartFromDefinition(chartDef)
}

// parseMultipleCharts parses multiple chart definitions and renders them side by side
func parseMultipleCharts(chartBlocks []string) (string, error) {
	// Parse each chart block
	chartDefs := make([]ChartDefinition, 0, len(chartBlocks))
	
	for i, block := range chartBlocks {
		if strings.TrimSpace(block) == "" {
			continue // Skip empty blocks
		}
		
		chartDef, err := parseChartDefinition(block, i)
		if err != nil {
			return "", fmt.Errorf("error in chart #%d: %w", i+1, err)
		}
		chartDefs = append(chartDefs, chartDef)
	}
	
	// Render charts side by side in a flex container
	var result strings.Builder
	result.WriteString(`<div style="display: flex; flex-wrap: wrap; justify-content: space-around; align-items: center; gap: 20px; margin: 20px 0;">`)
	
	for i, chartDef := range chartDefs {
		svg, err := renderChartFromDefinition(chartDef)
		if err != nil {
			return "", fmt.Errorf("error rendering chart #%d: %w", i+1, err)
		}
		
		result.WriteString(`<div style="flex: 1; min-width: 300px; max-width: 48%;">`)
		result.WriteString(svg)
		result.WriteString(`</div>`)
	}
	
	result.WriteString(`</div>`)
	return result.String(), nil
}

// parseChartDefinition parses a single chart definition
func parseChartDefinition(markdown string, chartIndex int) (ChartDefinition, error) {
	lines := strings.Split(markdown, "\n")
	
	var chartDef ChartDefinition
	
	// Default settings
	chartDef.Width = 800
	chartDef.Height = 500
	chartDef.Title = "Chart"
	chartDef.AutoHeight = false
	
	if len(lines) < 3 {
		return chartDef, fmt.Errorf("chart format invalid - too few lines. Need at least chart type, configuration, and data sections")
	}
	
	// Parse chart type from first line
	chartDef.ChartType = strings.TrimSpace(strings.ToLower(lines[0]))
	
	// Validate chart type
	validTypes := map[string]bool{
		"line": true, "linechart": true,
		"bar": true, "barchart": true,
		"pie": true, "piechart": true,
		"heatmap": true, "heatmapchart": true,
	}
	
	if !validTypes[chartDef.ChartType] {
		return chartDef, fmt.Errorf("unknown chart type '%s'. Must be one of: linechart, barchart, piechart, heatmapchart", chartDef.ChartType)
	}
	
	// Parse configuration and data
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
				
				chartDef.Labels = append(chartDef.Labels, label)
				
				if val, err := strconv.ParseFloat(valueStr, 64); err == nil {
					chartDef.Data = append(chartDef.Data, val)
				} else {
					dataErrors = append(dataErrors, fmt.Sprintf("line %d: '%s' is not a valid number", i+1, valueStr))
				}
			} else if len(parts) == 1 && parts[0] != "" {
				// No label, just data
				if val, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64); err == nil {
					chartDef.Data = append(chartDef.Data, val)
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
				chartDef.Title = value
			case "width":
				if w, err := strconv.Atoi(value); err == nil && w > 0 {
					chartDef.Width = w
				} else {
					configErrors = append(configErrors, fmt.Sprintf("line %d: invalid width value '%s' - must be a positive number", i+1, value))
				}
			case "height":
				if strings.ToLower(value) == "auto" {
					// Set auto-height flag
					chartDef.AutoHeight = true
					// Set a default height, will be recalculated in renderChartFromDefinition
					chartDef.Height = 0
				} else if h, err := strconv.Atoi(value); err == nil && h > 0 {
					chartDef.Height = h
				} else {
					configErrors = append(configErrors, fmt.Sprintf("line %d: invalid height value '%s' - must be a positive number or 'auto'", i+1, value))
				}
			case "colors":
				colorList := parseList(value)
				if len(colorList) == 0 {
					configErrors = append(configErrors, fmt.Sprintf("line %d: invalid colors format - must be comma-separated list", i+1))
				} else {
					chartDef.Colors = colorList
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
	
	if len(chartDef.Data) == 0 {
		errors = append(errors, "no valid data points found - chart requires at least one data point")
	}
	
	if len(chartDef.Labels) > 0 && len(chartDef.Labels) != len(chartDef.Data) {
		errors = append(errors, fmt.Sprintf("mismatched labels and data points - found %d labels but %d data points", len(chartDef.Labels), len(chartDef.Data)))
	}
	
	// If we have validation errors, return them
	if len(errors) > 0 {
		return chartDef, fmt.Errorf("chart definition errors:\n• %s", strings.Join(errors, "\n• "))
	}
	
	return chartDef, nil
}

// renderChartFromDefinition renders a chart from a ChartDefinition
func renderChartFromDefinition(chartDef ChartDefinition) (string, error) {
	// Create appropriate chart
	var chart gosvgchart.Chart
	
	switch chartDef.ChartType {
	case "line", "linechart":
		chart = gosvgchart.New()
	case "bar", "barchart":
		chart = gosvgchart.NewBarChart()
	case "pie", "piechart":
		chart = gosvgchart.NewPieChart()
	case "heatmap", "heatmapchart":
		chart = gosvgchart.NewHeatmapChart()
	}
	
	// Set basic properties
	chart.SetTitle(chartDef.Title)
	
	// Handle auto-height
	if chartDef.AutoHeight {
		chart.SetSize(chartDef.Width, 0)
		chart.SetAutoHeight(true)
	} else {
		chart.SetSize(chartDef.Width, chartDef.Height)
	}
	
	if len(chartDef.Colors) > 0 {
		chart.SetColors(chartDef.Colors)
	}
	
	// Set data and labels
	chart.SetData(chartDef.Data)
	if len(chartDef.Labels) > 0 {
		chart.SetLabels(chartDef.Labels)
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