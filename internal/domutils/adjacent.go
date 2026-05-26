package domutils

import (
	"golang.org/x/net/html"
)

func collectAdjacentNodes(node *html.Node, matchFn func(n *html.Node) bool) []*html.Node {
	_ = "STUB: not implemented"
	return nil
}

// A span has no special meaning. So we just skip it...

// Return the collected nodes

func mergeChildren(destinationNode *html.Node, nodes ...*html.Node) {
	_ = "STUB: not implemented"
	return
}

// We move all the children to the `destinationNode`.

func MergeAdjacent(doc *html.Node, matchFn func(*html.Node) bool) {
	_ = "STUB: not implemented"
	return
}

// - - - - - - - - //

func MergeAdjacentTextNodes(n *html.Node) { _ = "STUB: not implemented"; return }

// Combine adjacent text nodes
