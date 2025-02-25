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
	ChartType    string
	Title        string
	Width        int
	Height       int
	Colors       []string
	Data         []float64
	Labels       []string
	AutoHeight   bool
	Series       []SeriesDefinition
	SeriesColors []string
	Stacked      bool
}

// SeriesDefinition represents a data series in a chart
type SeriesDefinition struct {
	Name string
	Data []float64
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
	chartDef.Stacked = false

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

	// Variables for series support
	var currentSeries string

	// Variables for tabular format
	var inTabularFormat bool = false
	var seriesNames []string

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
			currentSeries = "" // Reset current series
			continue
		}

		// Check for series section with tabular format
		if strings.HasPrefix(line, "series:") && !strings.Contains(line, "|") {
			dataStarted = true
			foundDataSection = true

			// Check if the next line contains series names in tabular format
			if i+1 < len(lines) && strings.Contains(lines[i+1], "|") {
				inTabularFormat = true

				// Parse series names from the next line
				seriesLine := strings.TrimSpace(lines[i+1])
				seriesParts := strings.Split(seriesLine, "|")

				// Skip the first part (usually empty or "Series")
				for j := 1; j < len(seriesParts); j++ {
					seriesName := strings.TrimSpace(seriesParts[j])
					if seriesName != "" {
						seriesNames = append(seriesNames, seriesName)
						// Create a new series
						chartDef.Series = append(chartDef.Series, SeriesDefinition{
							Name: seriesName,
							Data: []float64{},
						})
					}
				}

				// Skip the series names line
				i++
				continue
			}
		}

		// Check for traditional series section
		if strings.HasPrefix(line, "series:") && !inTabularFormat {
			dataStarted = true
			foundDataSection = true
			// Extract series name
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				currentSeries = strings.TrimSpace(parts[1])
				// Create a new series
				chartDef.Series = append(chartDef.Series, SeriesDefinition{
					Name: currentSeries,
					Data: []float64{},
				})
			} else {
				dataErrors = append(dataErrors, fmt.Sprintf("line %d: invalid series format, expected 'series: name'", i+1))
			}
			continue
		}

		if dataStarted {
			// We're in the data section

			// Handle tabular format
			if inTabularFormat && strings.Contains(line, "|") {
				parts := strings.Split(line, "|")

				// First part is the label
				if len(parts) > 0 {
					label := strings.TrimSpace(parts[0])

					// Add label to main labels
					chartDef.Labels = append(chartDef.Labels, label)

					// Process each value for each series
					for j := 1; j < len(parts) && j-1 < len(seriesNames); j++ {
						valueStr := strings.TrimSpace(parts[j])

						// Add data to the corresponding series
						if val, err := strconv.ParseFloat(valueStr, 64); err == nil {
							chartDef.Series[j-1].Data = append(chartDef.Series[j-1].Data, val)
						} else {
							dataErrors = append(dataErrors, fmt.Sprintf("line %d: '%s' is not a valid number for series '%s'",
								i+1, valueStr, seriesNames[j-1]))
						}
					}
				}
				continue
			}

			// Handle traditional format
			parts := strings.Split(line, "|")

			// If line contains a pipe, it has a label
			if len(parts) == 2 {
				label := strings.TrimSpace(parts[0])
				valueStr := strings.TrimSpace(parts[1])

				if label == "" {
					dataErrors = append(dataErrors, fmt.Sprintf("line %d: missing label before '|'", i+1))
				}

				// If we're in a series, add to that series
				if currentSeries != "" {
					// Find the current series
					seriesIndex := -1
					for idx, s := range chartDef.Series {
						if s.Name == currentSeries {
							seriesIndex = idx
							break
						}
					}

					if seriesIndex == -1 {
						dataErrors = append(dataErrors, fmt.Sprintf("line %d: internal error - series '%s' not found", i+1, currentSeries))
						continue
					}

					// Add label to main labels if not already there
					labelExists := false
					for _, existingLabel := range chartDef.Labels {
						if existingLabel == label {
							labelExists = true
							break
						}
					}
					if !labelExists {
						chartDef.Labels = append(chartDef.Labels, label)
					}

					// Add data to series
					if val, err := strconv.ParseFloat(valueStr, 64); err == nil {
						chartDef.Series[seriesIndex].Data = append(chartDef.Series[seriesIndex].Data, val)
					} else {
						dataErrors = append(dataErrors, fmt.Sprintf("line %d: '%s' is not a valid number", i+1, valueStr))
					}
				} else {
					// Legacy single series
					chartDef.Labels = append(chartDef.Labels, label)

					if val, err := strconv.ParseFloat(valueStr, 64); err == nil {
						chartDef.Data = append(chartDef.Data, val)
					} else {
						dataErrors = append(dataErrors, fmt.Sprintf("line %d: '%s' is not a valid number", i+1, valueStr))
					}
				}
			} else if len(parts) == 1 && parts[0] != "" {
				// No label, just data
				if currentSeries != "" {
					// Find the current series
					seriesIndex := -1
					for idx, s := range chartDef.Series {
						if s.Name == currentSeries {
							seriesIndex = idx
							break
						}
					}

					if seriesIndex == -1 {
						dataErrors = append(dataErrors, fmt.Sprintf("line %d: internal error - series '%s' not found", i+1, currentSeries))
						continue
					}

					// Add data to series
					if val, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64); err == nil {
						chartDef.Series[seriesIndex].Data = append(chartDef.Series[seriesIndex].Data, val)
					} else {
						dataErrors = append(dataErrors, fmt.Sprintf("line %d: '%s' is not a valid number", i+1, parts[0]))
					}
				} else {
					// Legacy single series
					if val, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64); err == nil {
						chartDef.Data = append(chartDef.Data, val)
					} else {
						dataErrors = append(dataErrors, fmt.Sprintf("line %d: '%s' is not a valid number", i+1, parts[0]))
					}
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
			case "seriescolors":
				colorList := parseList(value)
				if len(colorList) == 0 {
					configErrors = append(configErrors, fmt.Sprintf("line %d: invalid seriescolors format - must be comma-separated list", i+1))
				} else {
					chartDef.SeriesColors = colorList
				}
			case "stacked":
				if strings.ToLower(value) == "true" || strings.ToLower(value) == "yes" || value == "1" {
					chartDef.Stacked = true
				} else if strings.ToLower(value) == "false" || strings.ToLower(value) == "no" || value == "0" {
					chartDef.Stacked = false
				} else {
					configErrors = append(configErrors, fmt.Sprintf("line %d: invalid stacked value '%s' - must be true/false, yes/no, or 1/0", i+1, value))
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
		errors = append(errors, "missing 'data:' or 'series:' section - chart must include a data section")
	}

	// Check if we have data in either the legacy format or series format
	hasData := len(chartDef.Data) > 0
	for _, series := range chartDef.Series {
		if len(series.Data) > 0 {
			hasData = true
			break
		}
	}

	if !hasData {
		errors = append(errors, "no valid data points found - chart requires at least one data point")
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
		lineChart := gosvgchart.NewLineChart()
		chart = lineChart
		// Enable legend for multiple series
		if len(chartDef.Series) > 0 {
			lineChart.ShowLegend = true
		}
	case "bar", "barchart":
		barChart := gosvgchart.NewBarChart()
		chart = barChart
		// Set stacked property if specified
		barChart.Stacked = chartDef.Stacked
		// Enable legend for multiple series
		if len(chartDef.Series) > 0 {
			barChart.ShowLegend = true
		}
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

	// Set colors
	if len(chartDef.Colors) > 0 {
		chart.SetColors(chartDef.Colors)
	}

	// Set series colors if specified
	if len(chartDef.SeriesColors) > 0 {
		chart.SetSeriesColors(chartDef.SeriesColors)
	}

	// Check if we have multiple series
	if len(chartDef.Series) > 0 {
		// Add each series to the chart
		for _, series := range chartDef.Series {
			chart.AddSeries(series.Name, series.Data)
		}
	} else {
		// Legacy single series support
		// Set data and labels
		chart.SetData(chartDef.Data)
	}

	// Set labels
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
