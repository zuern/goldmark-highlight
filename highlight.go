// Copyright 2021 Kevin Zuern. All rights reserved.

package highlight

import (
	"fmt"

	"github.com/yuin/goldmark/ast"
)

// Kind is the ast.NodeKind for *Highlight AST nodes.
var Kind = ast.NewNodeKind("Highlight")

// A Highlight ast.Node whose contents is highlighted.
type Highlight struct {
	ast.BaseInline
	Level int
}

// Kind of the node.
func (n *Highlight) Kind() ast.NodeKind {
	return Kind
}

// Dump prints the node to stdout for debugging purposes.
func (n *Highlight) Dump(src []byte, level int) {
	ast.DumpHelper(n, src, level, map[string]string{
		"Level": fmt.Sprintf("%v", n.Level),
	}, nil)
}
