package highlight_test

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
	highlight "github.com/zuern/goldmark-highlight"
)

func Example() {
	// Create a combined parser / renderer which understands highlight syntax.
	md := goldmark.New(
		goldmark.WithExtensions(
			&highlight.Extender{},
		),
	)

	markdown := []byte("==this is some highlighted text==")

	var buffer bytes.Buffer
	_ = md.Convert(markdown, &buffer)

	fmt.Println(buffer.String())
	// output: <p><mark>this is some highlighted text</mark></p>
}
