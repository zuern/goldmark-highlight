// Copyright 2021 Kevin Zuern. All rights reserved.

package highlight_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/yuin/goldmark/ast"
	highlight "github.com/zuern/goldmark-highlight"
)

func TestRenderBadNode(t *testing.T) {
	r := &highlight.Renderer{}
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	var node ast.Node
	status, err := r.Render(w, []byte("Highlight"), node, true)
	if status != ast.WalkStop {
		t.Fatal()
	}
	if err == nil {
		t.Fatal()
	}
}
