package table

import (
	"golang.org/x/net/html"
)

func selectHeaderRowNode(node *html.Node) *html.Node { _ = "STUB: not implemented"; return nil }

// YEAH we found the "tr" inside the "thead"

// YEAH we found the "th"

func selectNormalRowNodes(tableNode *html.Node, selectedHeaderRowNode *html.Node) []*html.Node {
	_ = "STUB: not implemented"
	return nil
}

// We want to make sure to not select the header row a *second* time.
