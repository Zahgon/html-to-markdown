package commonmark

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

// link in commonmark contains
// - the link text (the visible text)
// - a link destination (the URI that is the link destination)
// - an optional link title
type link struct {
	*html.Node

	before  []byte
	content []byte
	after   []byte

	href  string
	title string
}

func (c *commonmark) renderLinkInlined(w converter.Writer, l *link) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}

// The destination and title must be separated by a space

func (c *commonmark) renderLink(ctx converter.Context, w converter.Writer, n *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}

// There is *no href* for the link. Now we have two options:
// Continue rendering as a link OR skip to let other renderers take over.

// There is *no content* inside the link. Now we have two options:
// Continue rendering as a link OR skip to let other renderers take over.

// A link without href is valid, like e.g. [text]()
// But a title would make it invalid.

// Note: We don't want to use `TrimUnnecessaryHardLineBreaks` here,
// since `EscapeMultiLine` also takes care of newlines.
