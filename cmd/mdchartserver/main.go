package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/riclib/gosvgchart/mdparser"
)

func main() {
	// Define command line flags
	port := flag.Int("port", 8080, "Port to run the server on")
	flag.Parse()

	// Handler for chart generation from raw markdown
	http.HandleFunc("/chart", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		// Parse markdown and generate SVG
		svg, err := mdparser.ParseMarkdownChart(string(body))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing chart: %v", err), http.StatusBadRequest)
			return
		}

		// Send SVG response
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write([]byte(svg))
	})

	// Handler for chart generation from markdown in URL parameter
	http.HandleFunc("/charturl", func(w http.ResponseWriter, r *http.Request) {
		markdownEncoded := r.URL.Query().Get("md")
		if markdownEncoded == "" {
			http.Error(w, "Missing 'md' parameter", http.StatusBadRequest)
			return
		}

		// Simple URL-safe encoding: Replace '_n_' with newlines and '_p_' with pipes
		markdown := strings.ReplaceAll(markdownEncoded, "_n_", "\n")
		markdown = strings.ReplaceAll(markdown, "_p_", "|")

		// Parse markdown and generate SVG
		svg, err := mdparser.ParseMarkdownChart(markdown)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing chart: %v", err), http.StatusBadRequest)
			return
		}

		// Send SVG response
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write([]byte(svg))
	})

	// HTML demo page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>GoSVGChart Markdown Demo</title>
			<style>
				body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
				textarea { width: 100%; height: 300px; font-family: monospace; padding: 8px; margin-bottom: 10px; }
				button { padding: 8px 16px; background: #3498db; color: white; border: none; cursor: pointer; }
				.examples { margin-top: 20px; }
				.example { margin-bottom: 10px; }
				.output { margin-top: 20px; }
				.error-box { border: 2px solid #e74c3c; background-color: #ffeaea; padding: 15px; border-radius: 5px; }
				.error-heading { color: #e74c3c; margin-top: 0; }
				.error-content { font-family: monospace; white-space: pre-wrap; }
				.tabs { display: flex; margin-bottom: 10px; border-bottom: 1px solid #ddd; }
				.tab { padding: 8px 15px; cursor: pointer; }
				.tab.active { background: #3498db; color: white; }
				.tab-content { display: none; }
				.tab-content.active { display: block; }
				pre { background: #f7f7f7; padding: 10px; border-radius: 5px; cursor: pointer; }
				pre:hover { background: #eee; }
			</style>
		</head>
		<body>
			<h1>GoSVGChart Markdown Demo</h1>
			
			<div class="tabs">
				<div class="tab active" onclick="showTab('direct')">Direct Chart Format</div>
				<div class="tab" onclick="showTab('goldmark')">Goldmark Integration</div>
			</div>
			
			<div id="direct" class="tab-content active">
				<p>Enter chart specification in markdown format:</p>
				
				<textarea id="markdown">linechart
title: Monthly Sales by Product
width: 800
height: auto
seriescolors: #4285F4, #EA4335, #FBBC05

series:
Month | Product A | Product B | Product C
Jan | 120 | 200 | 50
Feb | 150 | 180 | 80
Mar | 180 | 160 | 110
Apr | 210 | 140 | 140
May | 240 | 120 | 170
Jun | 270 | 100 | 200</textarea>
				
				<div>
					<button onclick="generateChart()">Generate Chart</button>
				</div>
				
				<div class="output">
					<div id="chartOutput"></div>
				</div>
				
				<div class="examples">
					<h3>Examples:</h3>
					
					<div class="example">
						<h4>Line Chart</h4>
						<pre>linechart
title: Monthly Sales
width: 600
height: 400
colors: #3498db, #e74c3c

data:
Jan | 120
Feb | 250
Mar | 180
Apr | 310
May | 270
Jun | 390</pre>
					</div>
					
					<div class="example">
						<h4>Multiple Series Line Chart</h4>
						<pre>linechart
title: Monthly Sales by Product
width: 800
height: auto
seriescolors: #4285F4, #EA4335, #FBBC05, #34A853

series:
Month | Product A | Product B | Product C
Jan | 120 | 200 | 50
Feb | 150 | 180 | 80
Mar | 180 | 160 | 110
Apr | 210 | 140 | 140
May | 240 | 120 | 170
Jun | 270 | 100 | 200
Jan | 120
Feb | 150
Mar | 180
Apr | 210
May | 240
Jun | 270

series: Product B
Jan | 200
Feb | 180
Mar | 160
Apr | 140
May | 120
Jun | 100

series: Product C
Jan | 50
Feb | 80
Mar | 110
Apr | 140
May | 170
Jun | 200</pre>
					</div>
					
					<div class="example">
						<h4>Bar Chart</h4>
						<pre>barchart
title: Quarterly Revenue
width: 600
height: 400
colors: #2ecc71, #e74c3c, #f39c12, #9b59b6

data:
Q1 | 850
Q2 | 940
Q3 | 1100
Q4 | 1200</pre>
					</div>
					
					<div class="example">
						<h4>Grouped Bar Chart (Multiple Series)</h4>
						<pre>barchart
title: Quarterly Revenue by Region
width: 800
height: auto
stacked: false
seriescolors: #4285F4, #EA4335, #FBBC05, #34A853

series:
Quarter | North | South | East | West
Q1 | 150 | 120 | 90 | 180
Q2 | 180 | 140 | 110 | 200
Q3 | 210 | 160 | 130 | 220
Q4 | 240 | 180 | 150 | 240
Q1 | 150
Q2 | 180
Q3 | 210
Q4 | 240

series: South
Q1 | 120
Q2 | 140
Q3 | 160
Q4 | 180

series: East
Q1 | 90
Q2 | 110
Q3 | 130
Q4 | 150

series: West
Q1 | 180
Q2 | 200
Q3 | 220
Q4 | 240</pre>
					</div>
					
					<div class="example">
						<h4>Stacked Bar Chart (Multiple Series)</h4>
						<pre>barchart
title: Quarterly Revenue by Region (Stacked)
width: 800
height: auto
stacked: true
seriescolors: #4285F4, #EA4335, #FBBC05, #34A853

series:
Quarter | North | South | East | West
Q1 | 150 | 120 | 90 | 180
Q2 | 180 | 140 | 110 | 200
Q3 | 210 | 160 | 130 | 220
Q4 | 240 | 180 | 150 | 240
Q1 | 150
Q2 | 180
Q3 | 210
Q4 | 240

series: South
Q1 | 120
Q2 | 140
Q3 | 160
Q4 | 180

series: East
Q1 | 90
Q2 | 110
Q3 | 130
Q4 | 150

series: West
Q1 | 180
Q2 | 200
Q3 | 220
Q4 | 240</pre>
					</div>
					
					<div class="example">
						<h4>Tabular Format (Multiple Series)</h4>
						<pre>linechart
title: Monthly Sales by Product (Tabular Format)
width: 800
height: auto
seriescolors: #4285F4, #EA4335, #FBBC05

series:
Month | Product A | Product B | Product C
Jan | 120 | 200 | 50
Feb | 150 | 180 | 80
Mar | 180 | 160 | 110
Apr | 210 | 140 | 140
May | 240 | 120 | 170
Jun | 270 | 100 | 200</pre>
					</div>
					
					<div class="example">
						<h4>Tabular Format Bar Chart</h4>
						<pre>barchart
title: Quarterly Revenue by Region (Tabular Format)
width: 800
height: auto
stacked: false
seriescolors: #4285F4, #EA4335, #FBBC05, #34A853

series:
Quarter | North | South | East | West
Q1 | 150 | 120 | 90 | 180
Q2 | 180 | 140 | 110 | 200
Q3 | 210 | 160 | 130 | 220
Q4 | 240 | 180 | 150 | 240</pre>
					</div>
					
					<div class="example">
						<h4>Pie Chart</h4>
						<pre>piechart
title: Market Share
width: 600
height: 500
colors: #3498db, #2ecc71, #e74c3c, #f39c12, #9b59b6

data:
Product A | 35
Product B | 25
Product C | 20
Product D | 15
Others | 5</pre>
					</div>
					
					<div class="example">
						<h4>Heatmap Chart</h4>
						<pre>heatmapchart
title: GitHub Contribution Activity
width: 800
height: 200
colors: #ebedf0, #9be9a8, #40c463, #30a14e, #216e39

data:
2025-01-01 | 5
2025-01-05 | 12
2025-01-10 | 3
2025-01-15 | 15
2025-01-20 | 8
2025-01-25 | 4
2025-02-01 | 7
2025-02-05 | 14
2025-02-10 | 6
2025-02-15 | 11</pre>
					</div>
				</div>
			</div>
			
			<div id="goldmark" class="tab-content">
				<p>This tab shows how to use the Goldmark extension to include charts in your markdown documents.</p>
				
				<h3>Edit markdown with chart code blocks:</h3>
				<textarea id="goldmarkText" style="height: 450px;">
# Sales Report

## Quarterly Revenue

` + "```gosvgchart" + `
barchart
title: Quarterly Revenue
width: 600
height: 400
colors: #2ecc71, #e74c3c, #f39c12, #9b59b6

data:
Q1 | 850
Q2 | 940
Q3 | 1100
Q4 | 1200
` + "```" + `

## Monthly Sales Trend

` + "```gosvgchart" + `
linechart
title: Monthly Sales
width: 600
height: 400
colors: #3498db

data:
Jan | 120
Feb | 250
Mar | 180
Apr | 310
May | 270
Jun | 390
` + "```" + `

## Regional Sales Comparison

` + "```gosvgchart" + `
barchart
title: Quarterly Revenue by Region
width: 800
height: auto
stacked: false
seriescolors: #4285F4, #EA4335, #FBBC05, #34A853

series:
Quarter | North | South | East | West
Q1 | 150 | 120 | 90 | 180
Q2 | 180 | 140 | 110 | 200
Q3 | 210 | 160 | 130 | 220
Q4 | 240 | 180 | 150 | 240
` + "```" + `

## Product Sales Comparison

` + "```gosvgchart" + `
linechart
title: Monthly Sales by Product
width: 800
height: auto
seriescolors: #4285F4, #EA4335, #FBBC05

series:
Month | Product A | Product B | Product C
Jan | 120 | 200 | 50
Feb | 150 | 180 | 80
Mar | 180 | 160 | 110
Apr | 210 | 140 | 140
May | 240 | 120 | 170
Jun | 270 | 100 | 200
` + "```" + `

## Developer Activity

` + "```gosvgchart" + `
heatmapchart
title: GitHub Contributions
width: 800
height: 200
colors: #ebedf0, #9be9a8, #40c463, #30a14e, #216e39

data:
2025-01-01 | 5
2025-01-05 | 12
2025-01-10 | 3
2025-01-15 | 15
2025-01-20 | 8
2025-01-25 | 4
2025-02-01 | 7
2025-02-05 | 14
2025-02-10 | 6
2025-02-15 | 11
` + "```" + `

## Side-by-Side Charts
Charts placed next to each other will display side-by-side:

` + "```gosvgchart" + `
barchart
title: 2023 Revenue
width: 450
height: 300
colors: #3498db, #2ecc71

data:
Q1 | 850
Q2 | 940
Q3 | 1100
Q4 | 1200
` + "```" + `
` + "```gosvgchart" + `
barchart
title: 2024 Revenue
width: 450
height: 300
colors: #e74c3c, #f39c12

data:
Q1 | 950
Q2 | 1040
Q3 | 1200
Q4 | 1400
` + "```" + `
				</textarea>
				
				<div style="margin-top: 10px; margin-bottom: 20px;">
					<button onclick="renderGoldmarkExample()">Render Markdown with Charts</button>
				</div>
				
				<div class="output">
					<h3>Preview:</h3>
					<div id="goldmarkOutput" style="border: 1px solid #ddd; padding: 20px; border-radius: 5px; background: white;"></div>
				</div>
				
				<div class="example" style="margin-top: 20px;">
					<h4>Goldmark Integration Code:</h4>
					<pre>
import (
    "github.com/riclib/gosvgchart/goldmark"
    gm "github.com/yuin/goldmark"
)

// Create a new Goldmark instance with the gosvgchart extension
markdown := gm.New(
    gm.WithExtensions(
        goldmark.New(),
    ),
)

// Convert markdown to HTML with embedded SVG charts
var output []byte
if err := markdown.Convert([]byte(markdownContent), &output); err != nil {
    // Handle error
}
					</pre>
				</div>
				
				<p>
					For more details, see the <a href="https://github.com/riclib/gosvgchart/tree/main/goldmark">Goldmark extension documentation</a>.
				</p>
			</div>
			
			<script>
				function showTab(tabId) {
					// Hide all tab contents
					document.querySelectorAll(".tab-content").forEach(function(el) {
						el.classList.remove("active");
					});
					
					// Remove active class from all tabs
					document.querySelectorAll(".tab").forEach(function(el) {
						el.classList.remove("active");
					});
					
					// Activate selected tab
					document.getElementById(tabId).classList.add("active");
					
					// Find the tab button by its onclick attribute
					document.querySelector(".tab[onclick*='" + tabId + "']").classList.add("active");
				}
				
				function generateChart() {
					var markdown = document.getElementById("markdown").value;
					fetch("/chart", {
						method: "POST",
						body: markdown
					})
					.then(function(response) {
						if (response.ok) {
							return response.text().then(function(text) {
								return { success: true, content: text };
							});
						} else {
							return response.text().then(function(text) {
								return { success: false, error: text };
							});
						}
					})
					.then(function(result) {
						if (result.success) {
							document.getElementById("chartOutput").innerHTML = result.content;
						} else {
							var errorMsg = result.error.replace(/\n• /g, "<br>• ");
							document.getElementById("chartOutput").innerHTML = 
								'<div class="error-box">' +
								'<h3 class="error-heading">Chart Error</h3>' +
								'<div class="error-content">' + errorMsg + '</div>' +
								'</div>';
						}
					})
					.catch(function(error) {
						console.error("Error:", error);
						document.getElementById("chartOutput").innerHTML = "<p>Error connecting to server</p>";
					});
				}
				
				function renderGoldmarkExample() {
					// Get markdown content
					var markdownText = document.getElementById("goldmarkText").value;
					var output = "";
					
					// Process the markdown with simple rendering for headings
					var lines = markdownText.split("\n");
					var inChartBlock = false;
					var chartContent = "";
					var chartType = "";
					var chartCount = 0;
					var consecutiveCharts = [];
					
					for (var i = 0; i < lines.length; i++) {
						var line = lines[i];
						
						// Check for chart code blocks
						if (line.trim() === "` + "```gosvgchart" + `") {
							inChartBlock = true;
							chartContent = "";
							chartCount++;
							continue;
						}
						
						if (inChartBlock) {
							if (line.trim() === "` + "```" + `") {
								inChartBlock = false;
								
								// Store this chart to check for consecutive charts
								consecutiveCharts.push({
									content: chartContent,
									type: chartType
								});
								
								// Process the next line to see if it's another chart
								if (i + 1 < lines.length && lines[i + 1].trim() === "` + "```gosvgchart" + `") {
									// This is a consecutive chart, continue to the next iteration
									continue;
								}
								
								// If we get here, we need to render the charts we've collected
								if (consecutiveCharts.length > 1) {
									// Multiple charts - render side by side
									output += '<div style="display: flex; flex-wrap: wrap; justify-content: space-around; align-items: center; gap: 20px; margin: 20px 0;">';
									
									for (var j = 0; j < consecutiveCharts.length; j++) {
										var chartId = "chart-" + chartCount + "-" + (j + 1);
										output += '<div id="' + chartId + '" style="flex: 1; min-width: 300px; max-width: 48%;" class="chart-placeholder">Loading chart...</div>';
										
										// Use closure to capture chart details
										(function(id, content) {
											fetch("/chart", {
												method: "POST",
												body: content
											})
											.then(function(response) {
												if (response.ok) {
													return response.text();
												} else {
													return "<!-- Chart error -->";
												}
											})
											.then(function(svg) {
												var placeholder = document.getElementById(id);
												if (placeholder) {
													placeholder.innerHTML = svg;
												}
											});
										})(chartId, consecutiveCharts[j].content);
									}
									
									output += '</div>';
									consecutiveCharts = [];
								} else {
									// Single chart
									var chartId = "chart-" + chartCount;
									output += '<div id="' + chartId + '" class="chart-placeholder">Loading chart...</div>';
									
									// Use closure to capture chartId and chartContent
									(function(id, content) {
										fetch("/chart", {
											method: "POST",
											body: content
										})
										.then(function(response) {
											if (response.ok) {
												return response.text();
											} else {
												return "<!-- Chart error -->";
											}
										})
										.then(function(svg) {
											var placeholder = document.getElementById(id);
											if (placeholder) {
												placeholder.innerHTML = svg;
											}
										});
									})(chartId, consecutiveCharts[0].content);
									
									consecutiveCharts = [];
								}
								
								continue;
							}
							
							// First line of chart content is the chart type
							if (chartContent === "") {
								chartType = line.trim();
							}
							
							chartContent += line + "\n";
						} else {
							// Basic markdown rendering for headings
							if (line.startsWith("# ")) {
								output += "<h1>" + line.substring(2) + "</h1>";
							} else if (line.startsWith("## ")) {
								output += "<h2>" + line.substring(3) + "</h2>";
							} else if (line.startsWith("### ")) {
								output += "<h3>" + line.substring(4) + "</h3>";
							} else if (line.trim() === "") {
								output += "<p></p>";
							} else {
								output += "<p>" + line + "</p>";
							}
						}
					}
					
					document.getElementById("goldmarkOutput").innerHTML = output;
				}
				
				// Load the examples when clicking on them
				document.querySelectorAll(".example pre").forEach(function(el) {
					el.addEventListener("click", function() {
						var activeTab = document.querySelector(".tab.active");
						if (activeTab.getAttribute("onclick").includes("direct")) {
							document.getElementById("markdown").value = this.textContent;
						}
					});
				});
				
				// Generate chart on load
				window.onload = function() {
					generateChart();
					// Wait a bit before rendering the goldmark example
					setTimeout(renderGoldmarkExample, 100);
				};
			</script>
		</body>
		</html>
		`

		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	// Start server
	serverAddr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting server on http://localhost%s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
