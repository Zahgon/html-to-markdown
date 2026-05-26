package commonmark

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

func getStartAt(node *html.Node) int { _ = "STUB: not implemented"; return 0 }

func (c commonmark) getPrefixFunc(n *html.Node, sliceLength int) func(int) string {
	_ = "STUB: not implemented"
	return nil
}

// Pad the numbers so that all prefix numbers in the list take up the same space
// `%02d.` -> "01. "

func renderMultiLineListItem(w converter.Writer, content []byte, indentCount int) {
	_ = "STUB: not implemented"
	return
}

// Add indent to code block newlines

// The first line is already indented through the prefix,
// all other lines need the correct amount of spaces.

func (c commonmark) renderListContainer(ctx converter.Context, w converter.Writer, n *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}

// An item might have different lines that each
// must be indented with the correct count of spaces.
