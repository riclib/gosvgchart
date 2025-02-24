package goldmark

import (
	"bytes"
	"github.com/riclib/gosvgchart/mdparser"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// ChartCodeBlockHandler is a Goldmark extension for handling gosvgchart code blocks
type ChartCodeBlockHandler struct{}

// New returns a new ChartCodeBlockHandler extension
func New() goldmark.Extender {
	return &ChartCodeBlockHandler{}
}

// Extend adds the gosvgchart code block handling to the Goldmark parser
func (c *ChartCodeBlockHandler) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(NewChartTransformer(), 100),
		),
	)
	md.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewChartRenderer(), 100),
		),
	)
}

// ChartTransformer transforms fenced code blocks with "gosvgchart" language to ChartNode
type ChartTransformer struct{}

// NewChartTransformer returns a new chart transformer
func NewChartTransformer() parser.ASTTransformer {
	return &ChartTransformer{}
}

// Transform looks for code blocks with "gosvgchart" language and transforms them to ChartNode
func (c *ChartTransformer) Transform(doc *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		// Look for fenced code blocks
		fencedCodeBlock, ok := node.(*ast.FencedCodeBlock)
		if !ok {
			return ast.WalkContinue, nil
		}

		// Check if the language is "gosvgchart"
		language := fencedCodeBlock.Language(reader.Source())
		if !bytes.Equal(language, []byte("gosvgchart")) {
			return ast.WalkContinue, nil
		}

		// Extract the content of the code block
		var codeContent bytes.Buffer
		lines := fencedCodeBlock.Lines()
		for i := 0; i < lines.Len(); i++ {
			line := lines.At(i)
			codeContent.Write(line.Value(reader.Source()))
			if i < lines.Len()-1 {
				codeContent.WriteString("\n")
			}
		}

		// Create a chart node to replace the code block
		chart := NewChartNode()
		chart.SetMarkdown(codeContent.String())

		// Replace the code block with the chart node
		node.Parent().ReplaceChild(node.Parent(), node, chart)

		// Skip processing the rest of this node since we just replaced it
		return ast.WalkSkipChildren, nil
	})
}

// ChartNode represents a chart in the AST
type ChartNode struct {
	ast.BaseBlock
	markdown string
	svg      string
}

// NewChartNode returns a new ChartNode
func NewChartNode() *ChartNode {
	return &ChartNode{
		BaseBlock: ast.BaseBlock{},
	}
}

// Dump implements Node.Dump
func (n *ChartNode) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, map[string]string{
		"markdown": n.markdown,
	}, nil)
}

// Kind implements Node.Kind
func (n *ChartNode) Kind() ast.NodeKind {
	return ast.KindExtension
}

// SetMarkdown sets the markdown content of the chart
func (n *ChartNode) SetMarkdown(markdown string) {
	n.markdown = markdown
}

// Markdown returns the markdown content of the chart
func (n *ChartNode) Markdown() string {
	return n.markdown
}

// ChartRenderer renders ChartNode
type ChartRenderer struct{}

// NewChartRenderer returns a new ChartRenderer
func NewChartRenderer() renderer.NodeRenderer {
	return &ChartRenderer{}
}

// RegisterFuncs registers the render functions for ChartNode
func (r *ChartRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindExtension, r.renderChart)
}

// renderChart renders a ChartNode to HTML
func (r *ChartRenderer) renderChart(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	chartNode, ok := node.(*ChartNode)
	if !ok {
		return ast.WalkContinue, nil
	}

	// Parse the markdown and get the SVG
	svg, err := mdparser.ParseMarkdownChart(chartNode.Markdown())
	if err != nil {
		// If there's an error, output it as HTML comment
		w.WriteString("<!-- gosvgchart error: ")
		w.WriteString(err.Error())
		w.WriteString(" -->")
		return ast.WalkContinue, nil
	}

	// Write the SVG to the output
	w.WriteString(svg)
	return ast.WalkContinue, nil
}