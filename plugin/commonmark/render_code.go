package commonmark

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

func (c *commonmark) renderInlineCode(_ converter.Context, w converter.Writer, n *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	// TODO: configure delimeter in options?
	return *new(converter.RenderStatus)
}

// TODO: debug flag?

// fmt.Println("expected an empty inline code to be already removed")
// panic("expected an empty inline code to be already removed")

// TODO: configurable function to decide if inline or block?

// fmt.Println("inline code contains newlines")
// return c.renderBlockCode(ctx, w, n, render)

// No stripping occurs if the code span contains _only_ spaces:

// Newlines in the text aren't great, since this is inline code and not a code block.
// Newlines will be stripped anyway in the browser, but it won't be recognized as code
// from the markdown parser when there is more than one newline.

// Code contains a backtick as first character

// Code contains a backtick as last character

func (c *commonmark) renderBlockCode(_ converter.Context, w converter.Writer, n *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}

// We want to keep the original content inside the code block untouched.
// Because multiple newlines would be trimmed, we temporarily replace it with another character.

func getCodeLanguage(n *html.Node) string { _ = "STUB: not implemented"; return "" }

func getCodeWithoutTags(startNode *html.Node) ([]byte, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// TODO: what if multiple elements have an info string?

// - - - //

// if strings.TrimSpace(n.Data) == "" && strings.Contains(n.Data, "\n") {
// 	buf.WriteString("\n")
// }
