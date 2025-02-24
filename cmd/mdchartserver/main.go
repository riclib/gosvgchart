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
			</style>
		</head>
		<body>
			<h1>GoSVGChart Markdown Demo</h1>
			<p>Enter chart specification in markdown format:</p>
			
			<textarea id="markdown">linechart
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
Jun | 390</textarea>
			
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
			</div>
			
			<script>
				function generateChart() {
					const markdown = document.getElementById('markdown').value;
					fetch('/chart', {
						method: 'POST',
						body: markdown
					})
					.then(response => {
						if (response.ok) {
							return response.text().then(text => {
								return { success: true, content: text };
							});
						} else {
							return response.text().then(text => {
								return { success: false, error: text };
							});
						}
					})
					.then(result => {
						if (result.success) {
							document.getElementById('chartOutput').innerHTML = result.content;
						} else {
							const errorMsg = result.error.replace(/\n• /g, '<br>• ');
							document.getElementById('chartOutput').innerHTML = 
								'<div class="error-box">' +
								'<h3 class="error-heading">Chart Error</h3>' +
								'<div class="error-content">' + errorMsg + '</div>' +
								'</div>';
						}
					})
					.catch(error => {
						console.error('Error:', error);
						document.getElementById('chartOutput').innerHTML = '<p>Error connecting to server</p>';
					});
				}
				
				// Load the examples when clicking on them
				document.querySelectorAll('.example pre').forEach(el => {
					el.addEventListener('click', () => {
						document.getElementById('markdown').value = el.textContent;
					});
				});
				
				// Generate chart on load
				window.onload = generateChart;
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