package commonmark

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

func nameIsBold(node *html.Node) bool { _ = "STUB: not implemented"; return false }

func nameIsItalic(node *html.Node) bool { _ = "STUB: not implemented"; return false }

func nameIsBoldOrItalic(node *html.Node) bool { _ = "STUB: not implemented"; return false }

func nameIsBothBoldOrItalic(a, b *html.Node) bool { _ = "STUB: not implemented"; return false }

func nameIsPre(node *html.Node) bool { _ = "STUB: not implemented"; return false }

func nameIsInlineCode(node *html.Node) bool { _ = "STUB: not implemented"; return false }

func nameIsLink(node *html.Node) bool { _ = "STUB: not implemented"; return false }

func nameIsBothLink(a, b *html.Node) bool { _ = "STUB: not implemented"; return false }

func nameIsHeading(node *html.Node) bool { _ = "STUB: not implemented"; return false }

// func nameIsBlockquote(node *html.Node) bool {
// 	return dom.NodeName(node) == "blockquote"
// }

func (c *commonmark) handlePreRender(ctx converter.Context, doc *html.Node) {
	_ = "STUB: not implemented"
	return
}

// domutils.SplitUp(ctx, doc, nameIsBoldOrItalic, nameIsLink, atom.Span)

// domutils.SplitUp(ctx, doc, nameIsLink, nameIsHeading, atom.Div)
// domutils.SplitUp(ctx, doc, nameIsLink, nameIsBlockquote, atom.Div)

// - - - Bold / Italic - - - //

// domutils.MovePunctuation(ctx, doc, nameIsBoldOrItalic)

// - - - Code - - - //

// - - - Link - - - //

// - - - Heading - - - //

// - - - List - - - //
