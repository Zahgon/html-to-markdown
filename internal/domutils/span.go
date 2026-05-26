package domutils

import (
	"context"

	"golang.org/x/net/html"
)

func isFakeSpan(node *html.Node) bool { _ = "STUB: not implemented"; return false }

// RenameFakeSpans renames all "span" nodes to "div" if
// any block element is found as a child.
func RenameFakeSpans(ctx context.Context, doc *html.Node) { _ = "STUB: not implemented"; return }
