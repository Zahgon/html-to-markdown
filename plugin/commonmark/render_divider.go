package commonmark

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

func (c *commonmark) renderDivider(_ converter.Context, w converter.Writer, _ *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}
