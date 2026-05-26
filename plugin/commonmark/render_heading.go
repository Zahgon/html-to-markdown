package commonmark

import (
	"regexp"

	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

// TODO: remove regex
var multipleSpacesR = regexp.MustCompile(`  +`)

func (r *commonmark) setextUnderline(level int, width int) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (r *commonmark) atxPrefix(level int) []byte { _ = "STUB: not implemented"; return nil }

func getHeadingLevel(name string) int { _ = "STUB: not implemented"; return 0 }

func runeCount(chars []rune) (count int) { _ = "STUB: not implemented"; return 0 }

func getUnderlineWidth(content []byte, minVal int) int { _ = "STUB: not implemented"; return 0 }

// Count how wide the line should be,
// while using RuneCount to correctly count ä, ö, ...
//
// TODO: optimize function w := utf8.RuneCount(part)

// Technically the minimum value is only one character,
// but one dash could easily trigger a heading.

func escapePoundSignAtEnd(s []byte) []byte {
	_ = "STUB: not implemented"
	// -1 #
	// -2 placeholder
	// -3 maybe \
	return nil
}

// We don't have a # at the end,
// so there is no work to do...

// It is already escaped,
// so there is no work to do...

// Because we have a # at the end,
// we should manually force the escaping
// by overriding the placeholder.

func (c *commonmark) renderHeading(ctx converter.Context, w converter.Writer, n *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	// ctx = context.WithValue(ctx, "is_inside_heading", true)
	return *new(converter.RenderStatus)
}

// Note: We don't want to use `TrimUnnecessaryHardLineBreaks` here,
// since `EscapeMultiLine` also takes care of newlines.

// Replace multiple spaces by one space.

// A # sign at the end would be removed otherwise
