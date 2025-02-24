# Changelog

All notable changes to the GoSVGChart project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Added new heatmap chart type for GitHub-style contribution/activity charts

## [0.3.0] - 2025-02-25

### Changed
- Side-by-side charts now require a single code block with multiple chart definitions separated by `---` instead of multiple adjacent code blocks

## [0.2.1] - 2025-02-25

### Added
- Added LICENSE file with MIT license
- Enhanced CLAUDE.md with comprehensive project information and documentation
- Added explicit license information to the project

## [0.2.0] - 2025-02-24

### Added
- Goldmark extension for integrating charts directly into markdown documents (a82f44b)
- Side-by-side chart rendering in Goldmark extension for responsive chart layouts (516e92e)
- Documentation in goldmark/README.md with examples and usage instructions
- New features in the web server demo to showcase Goldmark extension capabilities

### Changed
- Updated documentation to use `gosvgchart` language identifier for markdown code blocks (e4cf737)
- Improved error handling in mdparser with comprehensive test coverage (2781d54)
- Enhanced prompt.md with examples for side-by-side chart rendering (e4cf737)
- Updated README.md with Goldmark extension documentation and examples (bda2804)

### Fixed
- Several bug fixes in the markdown parser and chart rendering logic (2781d54)
- Web server demo improvements and interface enhancements

## [0.1.0] - 2025-02-21

### Added
- Initial release of GoSVGChart
- Support for three chart types: line, bar, and pie/donut
- Simple markdown-like text format for defining charts
- Go API with chainable methods for creating charts
- Command-line tool (mdchart) for converting markdown to SVG
- Web server (mdchartserver) for generating charts dynamically
- Basic styling and customization options