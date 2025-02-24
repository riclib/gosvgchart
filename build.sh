#!/bin/bash
set -e

echo "Building gosvgchart..."

# Build CLI tool
echo "Building mdchart CLI tool..."
go build -o bin/mdchart ./cmd/mdchart

# Build web server
echo "Building mdchartserver..."
go build -o bin/mdchartserver ./cmd/mdchartserver

echo "Build complete. Binaries available in bin/ directory"