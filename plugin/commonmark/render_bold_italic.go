package commonmark

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

func (c commonmark) getDelimiter(n *html.Node) []byte { _ = "STUB: not implemented"; return nil }

func (c commonmark) renderBoldItalic(ctx converter.Context, w converter.Writer, n *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}

// Depending on the options & whether it is bold or italic there
// is going to be a different delimiter.

// If there is a newline character between the start and end delimiter
// the delimiters won't be recognized. Either we remove all newline characters
// OR on _every_ line we put start & end delimiters.
