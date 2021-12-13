// Copyright 2021 Kevin Zuern. All rights reserved.

package highlight

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// Parser for parsing highlight markdown into a *Highlight ast.Node.
type Parser struct{}

// Trigger returns a list of characters that triggers Parse method of
// this parser.
func (p *Parser) Trigger() []byte {
	return []byte{'='}
}

// Parse parse the given block into an inline node.
func (p *Parser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	before := block.PrecendingCharacter()
	line, seg := block.PeekLine()
	node := parser.ScanDelimiter(line, before, 2, defaultDelimiterProcessor)
	if node == nil {
		return nil
	}
	node.Segment = seg.WithStop(seg.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}

var _ parser.InlineParser = (*Parser)(nil)
