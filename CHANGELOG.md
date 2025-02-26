# Changelog

All notable changes to the GoSVGChart project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Added palette feature for automatic color assignment
  - New "auto" palette assigns distinct colors to data points or series
  - New "gradient" palette creates color gradients by varying lightness
  - Added SetPalette() method to the Chart interface
  - Works with all chart types including multiple series charts

## [0.9.1] - 2025-02-26

### Changed
- Simplified chart generation instructions in prompt.md
- Streamlined multiple series examples to use only tabular format
- Improved documentation for better model usability

## [0.9.0] - 2025-02-26

### Added
- Added LegendWidth property to reserve space for legends
  - New SetLegendWidth() method for controlling legend position
  - Default 20% width reservation for multiple series legends
  - Adjusted chart area dimensions to prevent legend overlap
  - Improved layout for pie charts with legends

## [0.8.0] - 2025-04-01

### Added
- Multiple series support for line and bar charts
  - New `AddSeries()` and `SetSeriesColors()` methods for all chart types
  - Support for stacked bar charts with multiple series
  - Automatic legend generation for multiple series
  - Updated markdown parser to support multiple series declarations
  - New examples demonstrating multiple series usage
- Tabular format for multiple series in markdown
  - More intuitive and readable format for defining multiple series
  - Series names defined in a header row followed by data rows
  - Reduced repetition of labels and better alignment of data points
  - Backward compatible with the original series format
  - Examples added to documentation and demo page

## [0.7.0] - 2025-03-15

### Added
- Responsive SVG output for all chart types
  - Charts now use `width="100%" height="auto"` with viewBox for better responsiveness
  - Dimensions specified now define aspect ratio rather than fixed pixel sizes
  - Charts automatically adapt to container width while maintaining proportions
  - Improved embedding experience in responsive layouts and mobile devices

## [0.6.0] - 2025-03-15

### Added
- Responsive SVG output for all chart types
  - Charts now use `width="100%" height="auto"` with viewBox for better responsiveness
  - Dimensions specified now define aspect ratio rather than fixed pixel sizes
  - Charts automatically adapt to container width while maintaining proportions
  - Improved embedding experience in responsive layouts and mobile devices

## [0.5.1] - 2025-02-26

### Added
- Improved pie chart label handling with truncation and tooltips
- Label length limiting to avoid text overflow with `SetMaxLabelLength()` method
- Tooltip support for truncated labels and small pie slices
- Better positioning algorithm for labels in small pie slices

### Fixed
- Fixed issue with pie chart legend labels getting cut off when labels are long
- Optimized legend placement to prevent overlap with chart content
- Better handling of very small pie slices with tooltips for readability

## [0.5.0] - 2025-03-01

### Added
- Dark mode support for all chart types
  - Automatic adaptation to user's system color scheme preference
  - Customizable dark and light themes via API
  - Dark mode enabled by default with sensible defaults
- Improved heatmap chart with adaptivity features:
  - Auto-sizing cells based on available space
  - Single-letter day labels (S, M, T, W, T, F, S)
  - Better calendar visualization with responsive layout
- Comprehensive examples command in `cmd/examples` with:
  - Heatmap examples showing adaptivity at different sizes
  - Dark mode compatible line, bar, and pie chart examples
  - Command-line interface to run different example types

### Changed
- Updated heatmap chart to calculate optimal cell size based on container dimensions
- Improved documentation with details on dark mode usage and heatmap adaptivity
- Deprecated explicit day labels for heatmap in favor of automatic single-letter labels

## [0.4.0] - 2025-02-25

### Added
- Added new heatmap chart type for GitHub-style contribution/activity charts
- Added auto-height option for all chart types
  - Line, Bar, and Pie charts now use a standard 16:9 aspect ratio when auto-height is enabled
  - Heatmap charts use a fixed 250px height when auto-height is enabled
  - Can be specified in markdown format with `height: auto`

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