# Goldmark-Highlight

[![https://pkg.go.dev/github.com/zuern/goldmark-highlight](https://pkg.go.dev/badge/github.com/zuern/goldmark-highlight.svg)](https://pkg.go.dev/github.com/zuern/goldmark-highlight)
[![Go Report Card](https://goreportcard.com/badge/github.com/zuern/goldmark-highlight)](https://goreportcard.com/report/github.com/zuern/goldmark-highlight)

An extension to the [Goldmark Markdown Parser](https://github.com/yuin/goldmark)
which adds parsing / rendering capabilities for rendering highlighted text.

Highlighted text is denoted by an opening and closing pair of equals signs
e.g. `==this text is highlighted==` source text renders into
`<p><mark>this text is highlighted</mark></p>`.

See `example_test.go` in this package for a minimal example how to integrate
this module into your Goldmark parser/renderer.
