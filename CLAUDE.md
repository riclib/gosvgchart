# GoSVGChart Information

## Tech Stack
- Go (1.23.4+)
- Standard library for core functionality
- HTML/SVG for chart rendering
- github.com/yuin/goldmark for markdown integration
- HTTP server from standard library

## Project Structure
- `chart.go`: Core chart implementation with fluent API
- `mdparser/`: Parser for markdown-like chart format
- `cmd/`: Command-line tools (mdchart and mdchartserver)
- `goldmark/`: Extension for Goldmark markdown parser
- `examples/`: Sample chart definitions

## Build and Test
- Build binaries: `./build.sh`
- Run tests: `go test ./...`
- Generate example charts: `./bin/mdchart -input examples/sample_chart.md -output chart.svg`
- Start web server: `./bin/mdchartserver -port 8080`

## Development Practices
- Add unit tests to every file
- Maintain backward compatibility
- Follow Go code style conventions
- Update CHANGELOG.md for all significant changes
- Use semantic versioning

## License
The GoSVGChart project is licensed under the MIT License.

```
MIT License

Copyright (c) 2025 Richard Jörg Libiez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```