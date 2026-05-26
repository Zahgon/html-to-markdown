package base

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

// RenderAsHTML will render the node as HTML using `html.Render()`
// Newlines will be inserted depending on the configured `TagType`.
//
// As an example, you could do such a combination:
//
//	"A text with <strong>bold</strong> and *italic* text"`
func RenderAsHTML(ctx converter.Context, w converter.Writer, node *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}

// TODO: what to do with error?

// RenderAsHTMLWrapper will render the node as HTML
// and render the children as markdown.
func RenderAsHTMLWrapper(ctx converter.Context, w converter.Writer, node *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}

// TODO: also render the attributes?

// RenderAsPlaintextWrapper will keep the children of this node as markdown.
func RenderAsPlaintextWrapper(ctx converter.Context, w converter.Writer, node *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}
