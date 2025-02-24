package goldmark

import (
	"strings"
	"testing"

	gm "github.com/yuin/goldmark"
)

func TestChartExtension(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		contains []string // Strings that should be in the output
	}{
		{
			name: "Bar Chart",
			markdown: "```gosvgchart\nbarchart\ntitle: Test Bar Chart\nwidth: 400\nheight: 300\n\ndata:\nA | 10\nB | 20\n```",
			contains: []string{
				"<svg", 
				"width=\"400\"", 
				"height=\"300\"",
				"Test Bar Chart",
			},
		},
		{
			name: "Pie Chart",
			markdown: "```gosvgchart\npiechart\ntitle: Test Pie Chart\nwidth: 400\nheight: 300\n\ndata:\nA | 30\nB | 70\n```",
			contains: []string{
				"<svg", 
				"width=\"400\"", 
				"height=\"300\"",
				"Test Pie Chart",
			},
		},
		{
			name: "Line Chart",
			markdown: "```gosvgchart\nlinechart\ntitle: Test Line Chart\nwidth: 400\nheight: 300\n\ndata:\nA | 10\nB | 20\nC | 15\n```",
			contains: []string{
				"<svg", 
				"width=\"400\"", 
				"height=\"300\"",
				"Test Line Chart",
			},
		},
		{
			name: "Multiple Charts",
			markdown: "```gosvgchart\nlinechart\ntitle: First Chart\nwidth: 400\nheight: 300\n\ndata:\nA | 10\nB | 20\n\n---\n\nbarchart\ntitle: Second Chart\nwidth: 400\nheight: 300\n\ndata:\nX | 30\nY | 40\n```",
			contains: []string{
				"display: flex",
				"First Chart",
				"Second Chart",
				"<svg", // Should have multiple SVG tags
			},
		},
		{
			name: "Invalid Chart",
			markdown: "```gosvgchart\ninvalidchart\ntitle: Invalid Chart\n\ndata:\nA | 10\nB | 20\n```",
			contains: []string{
				"<!-- gosvgchart error:",
				"unknown chart type",
			},
		},
		{
			name: "Regular Markdown",
			markdown: "# Header\n\nRegular paragraph\n\n```go\nfmt.Println(\"hello\")\n```",
			contains: []string{
				"<h1>Header</h1>",
				"<pre><code class=\"language-go\">",
				"fmt.Println(&quot;hello&quot;)",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new Goldmark instance with the chart extension
			markdown := gm.New(
				gm.WithExtensions(
					New(),
				),
			)

			// Convert markdown to HTML
			var output strings.Builder
			if err := markdown.Convert([]byte(tt.markdown), &output); err != nil {
				t.Errorf("failed to convert markdown: %v", err)
				return
			}

			// Check if output contains expected strings
			html := output.String()
			for _, s := range tt.contains {
				if !strings.Contains(html, s) {
					t.Errorf("output does not contain %q, got: %s", s, html)
				}
			}
		})
	}
}