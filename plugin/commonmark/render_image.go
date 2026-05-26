package commonmark

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

func escapeAlt(altString string) string { _ = "STUB: not implemented"; return "" }

func (c *commonmark) renderImage(ctx converter.Context, w converter.Writer, n *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}

// The alt description will be placed between two square brackets `[alt]`
// so make sure that those characters are escaped.

// The destination and title must be seperated by a space
