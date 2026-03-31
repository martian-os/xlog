package mathjax

import (
	"testing"

	"github.com/emad-elsaid/xlog/markdown"
	"github.com/emad-elsaid/xlog/markdown/parser"
	"github.com/emad-elsaid/xlog/markdown/renderer"
	"github.com/emad-elsaid/xlog/markdown/testutil"
	"github.com/emad-elsaid/xlog/markdown/util"
)

func TestMathJax(t *testing.T) {
	md := markdown.New(
		markdown.WithParserOptions(
			parser.WithBlockParsers(
				util.Prioritized(&mathJaxBlockParser{}, 999),
			),
			parser.WithInlineParsers(
				util.Prioritized(&inlineMathParser{}, 999),
			),
		),
		markdown.WithRendererOptions(
			renderer.WithNodeRenderers(
				util.Prioritized(&InlineMathRenderer{startDelim: `\(`, endDelim: `\)`}, 0),
				util.Prioritized(&MathBlockRenderer{startDelim: `\[`, endDelim: `\]`}, 0),
			),
		),
	)

	testutil.DoTestCaseFile(md, "mathjax.txt", t)
}
