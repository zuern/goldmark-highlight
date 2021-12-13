// Copyright 2021 Kevin Zuern. All rights reserved.

package highlight_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yuin/goldmark"
	highlight "github.com/zuern/goldmark-highlight"
)

func TestHighlight(t *testing.T) {
	type tc struct {
		input  string
		expect string
	}
	for i, test := range []tc{
		{"=this is not highlighted=", "<p>=this is not highlighted=</p>"},
		{"this is not highlighted", "<p>this is not highlighted</p>"},
		{"==this is highlighted==", "<p><mark>this is highlighted</mark></p>"},
		{"==this\nis highlighted==", "<p><mark>this\nis highlighted</mark></p>"},
		{"==**bold**==", "<p><mark><strong>bold</strong></mark></p>"},
	} {
		md := goldmark.New(goldmark.WithExtensions(
			&highlight.Extender{},
		))
		var b bytes.Buffer
		err := md.Convert([]byte(test.input), &b)
		if err != nil {
			t.Fatal(err)
		}
		out := strings.TrimSpace(b.String())
		if out != test.expect {
			t.Fatalf("test %d:\n%q does not equal\n%q", i, out, test.expect)
		}
	}
}
