package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	
	"github.com/riclib/gosvgchart"
)

func main() {
	// Handle the root path to display a form
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>GoSVGChart Demo</title>
			<style>
				body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
				form { margin-bottom: 30px; }
				label { display: block; margin-top: 10px; }
				input, select { margin-bottom: 5px; }
				button { padding: 8px 16px; background: #3498db; color: white; border: none; cursor: pointer; }
				.charts { display: flex; flex-wrap: wrap; gap: 20px; }
				.chart { border: 1px solid #ddd; padding: 10px; border-radius: 5px; }
			</style>
		</head>
		<body>
			<h1>GoSVGChart Demo</h1>
			
			<form action="/chart" method="get">
				<div>
					<label for="chartType">Chart Type:</label>
					<select id="chartType" name="type">
						<option value="line">Line Chart</option>
						<option value="bar">Bar Chart</option>
						<option value="pie">Pie Chart</option>
					</select>
				</div>
				
				<div>
					<label for="title">Chart Title:</label>
					<input type="text" id="title" name="title" value="Sample Chart">
				</div>
				
				<div>
					<label for="data">Data (comma separated):</label>
					<input type="text" id="data" name="data" value="10,45,30,60,25,15,40">
				</div>
				
				<div>
					<label for="labels">Labels (comma separated):</label>
					<input type="text" id="labels" name="labels" value="Mon,Tue,Wed,Thu,Fri,Sat,Sun">
				</div>
				
				<div>
					<label for="colors">Colors (comma separated hex, e.g. #ff0000):</label>
					<input type="text" id="colors" name="colors" value="#3498db,#2ecc71,#e74c3c,#f39c12,#9b59b6">
				</div>
				
				<div>
					<label for="width">Width:</label>
					<input type="number" id="width" name="width" value="600">
				</div>
				
				<div>
					<label for="height">Height:</label>
					<input type="number" id="height" name="height" value="400">
				</div>
				
				<div>
					<button type="submit">Generate Chart</button>
				</div>
			</form>
			
			<div class="charts">
				<div class="chart">
					<h3>Line Chart Example</h3>
					<img src="/chart?type=line&title=Weekly%20Sales&data=10,45,30,60,25,15,40&labels=Mon,Tue,Wed,Thu,Fri,Sat,Sun&colors=%233498db&width=350&height=250" alt="Line Chart">
				</div>
				
				<div class="chart">
					<h3>Bar Chart Example</h3>
					<img src="/chart?type=bar&title=Monthly%20Revenue&data=350,420,520,410,390,450&labels=Jan,Feb,Mar,Apr,May,Jun&colors=%232ecc71,%23e74c3c,%23f39c12,%239b59b6&width=350&height=250" alt="Bar Chart">
				</div>
				
				<div class="chart">
					<h3>Pie Chart Example</h3>
					<img src="/chart?type=pie&title=Market%20Share&data=35,25,20,15,5&labels=A,B,C,D,Other&colors=%233498db,%232ecc71,%23e74c3c,%23f39c12,%239b59b6&width=350&height=250" alt="Pie Chart">
				</div>
			</div>
		</body>
		</html>
		`
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)
	})

	// Handle chart generation
	http.HandleFunc("/chart", func(w http.ResponseWriter, r *http.Request) {
		// Get parameters
		chartType := r.URL.Query().Get("type")
		title := r.URL.Query().Get("title")
		dataStr := r.URL.Query().Get("data")
		labelsStr := r.URL.Query().Get("labels")
		colorsStr := r.URL.Query().Get("colors")
		
		// Parse width and height
		width, err := strconv.Atoi(r.URL.Query().Get("width"))
		if err != nil || width <= 0 {
			width = 600
		}
		
		height, err := strconv.Atoi(r.URL.Query().Get("height"))
		if err != nil || height <= 0 {
			height = 400
		}
		
		// Parse data
		dataSlice := strings.Split(dataStr, ",")
		var data []float64
		for _, d := range dataSlice {
			val, err := strconv.ParseFloat(strings.TrimSpace(d), 64)
			if err == nil {
				data = append(data, val)
			}
		}
		
		// Parse labels
		var labels []string
		if labelsStr != "" {
			labels = strings.Split(labelsStr, ",")
			for i, label := range labels {
				labels[i] = strings.TrimSpace(label)
			}
		}
		
		// Parse colors
		var colors []string
		if colorsStr != "" {
			colors = strings.Split(colorsStr, ",")
			for i, color := range colors {
				colors[i] = strings.TrimSpace(color)
			}
		}
		
		// Create chart based on type
		var svg string
		
		switch chartType {
		case "line":
			chart := gosvgchart.New().
				SetTitle(title).
				SetData(data).
				SetLabels(labels).
				SetSize(width, height)
			
			if len(colors) > 0 {
				chart.SetColors(colors)
			}
			
			svg = chart.Render()
			
		case "bar":
			chart := gosvgchart.New().
				SetTitle(title).
				SetData(data).
				SetLabels(labels).
				SetSize(width, height)
			
			if len(colors) > 0 {
				chart.SetColors(colors)
			}
			
			svg = chart.Render()
			
		case "pie":
			chart := gosvgchart.New().
				SetTitle(title).
				SetData(data).
				SetLabels(labels).
				SetSize(width, height)
			
			if len(colors) > 0 {
				chart.SetColors(colors)
			}
			
			svg = chart.Render()
			
		default:
			http.Error(w, "Invalid chart type", http.StatusBadRequest)
			return
		}
		
		// Send SVG as response
		w.Header().Set("Content-Type", "image/svg+xml")
		fmt.Fprint(w, svg)
	})

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}