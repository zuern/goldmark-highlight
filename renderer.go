// Copyright 2021 Kevin Zuern. All rights reserved.

package highlight

import (
	"fmt"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// Renderer renders a *Highlight ast.Node into HTML.
type Renderer struct{}

// RendererFuncs registers NodeRendererFuncs to given NodeRendererFuncRegisterer.
func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(Kind, r.Render)
}

// Render the AST node to the given writer.
func (r *Renderer) Render(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	_, ok := n.(*Highlight)
	if !ok {
		return ast.WalkStop, fmt.Errorf("expected *Highlight, got %T", n)
	}
	status := ast.WalkStop
	var err error
	if entering {
		_, err = w.WriteString("<mark>")
	} else {
		_, err = w.WriteString("</mark>")
	}
	if err == nil {
		status = ast.WalkContinue
	}
	return status, err
}
