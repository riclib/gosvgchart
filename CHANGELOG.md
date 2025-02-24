# Changelog

All notable changes to the GoSVGChart project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2025-02-24

### Added
- Goldmark extension for integrating charts directly into markdown documents
- Side-by-side chart rendering in Goldmark extension for responsive chart layouts
- New tab in web server demo to showcase Goldmark extension and side-by-side charts
- Documentation for Goldmark integration and side-by-side chart feature

### Changed
- Updated all code examples to use `gosvgchart` language identifier for markdown code blocks
- Improved error handling for chart parsing and rendering
- Enhanced prompt.md with better examples and side-by-side chart guidance

## [0.1.0] - 2025-02-21

### Added
- Initial release of GoSVGChart
- Support for three chart types: line, bar, and pie/donut
- Simple markdown-like text format for defining charts
- Go API with chainable methods for creating charts
- Command-line tool (mdchart) for converting markdown to SVG
- Web server (mdchartserver) for generating charts dynamically
- Basic styling and customization options