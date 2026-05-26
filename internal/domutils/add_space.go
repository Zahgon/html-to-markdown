package domutils

import (
	"context"

	"golang.org/x/net/html"
)

func getFirstChildNode(startNode *html.Node, matchFn func(n *html.Node) bool) *html.Node {
	_ = "STUB: not implemented"
	return nil
}

// A span has no special meaning. So we just skip it...

func getLastChildNode(startNode *html.Node, matchFn func(n *html.Node) bool) *html.Node {
	_ = "STUB: not implemented"
	return nil
}

// A span has no special meaning. So we just skip it...

func AddSpace(ctx context.Context, doc *html.Node, isOuterNode, isInnerNode func(*html.Node) bool) {
	_ = "STUB: not implemented"
	return
}
