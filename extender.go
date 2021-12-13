// Copyright 2021 Kevin Zuern. All rights reserved.

package highlight

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// Extender implements the goldmark Extender interface and extends
// goldmark.Markdown to understand highlighting syntax.
type Extender struct{}

// Extend extends md to understand highlighting syntax.
func (e *Extender) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(&Parser{}, 999),
		),
	)
	md.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(&Renderer{}, 999),
		),
	)
}
