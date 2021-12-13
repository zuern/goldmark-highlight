// Copyright 2021 Kevin Zuern. All rights reserved.

package highlight_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	highlight "github.com/zuern/goldmark-highlight"
)

func TestHighlightDump(t *testing.T) {
	// capture stdout in file
	file := filepath.Join(t.TempDir(), "stdout")
	out, err := os.Create(file)
	if err != nil {
		t.Fatal(err)
	}
	defer func(stdout *os.File) { os.Stdout = stdout }(os.Stdout)
	os.Stdout = out

	const src = "==highlight=="
	node := &highlight.Highlight{}
	node.AppendChild(node, ast.NewTextSegment(text.NewSegment(0, len(src))))

	node.Dump([]byte(src), 0)

	if err = out.Close(); err != nil {
		t.Fatal(err)
	}

	bytes, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}

	actual := string(bytes)
	const expect = `Highlight {
    Level: 0
    Text: "==highlight=="
}
`

	if expect != actual {
		t.Fatalf("%s != %s", expect, actual)
	}
}
