// Copyright 2021 Kevin Zuern. All rights reserved.

package highlight

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
)

type delimiterProcessor struct{}

var defaultDelimiterProcessor = &delimiterProcessor{}

// IsDelimiter returns true if given character is a delimiter, otherwise false.
func (d *delimiterProcessor) IsDelimiter(b byte) bool {
	return b == '='
}

// CanOpenCloser returns true if given opener can close given closer, otherwise false.
func (d *delimiterProcessor) CanOpenCloser(opener *parser.Delimiter, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

// OnMatch will be called when new matched delimiter found.
// OnMatch should return a new Node correspond to the matched delimiter.
func (d *delimiterProcessor) OnMatch(consumes int) ast.Node {
	return &Highlight{}
}
