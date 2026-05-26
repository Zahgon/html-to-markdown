package domutils

import (
	"golang.org/x/net/html"
)

func RemoveRedundant(doc *html.Node, matchFn func(*html.Node, *html.Node) bool) {
	_ = "STUB: not implemented"
	return
}

func hasSameTypeAncestor(n *html.Node, matchFn func(*html.Node, *html.Node) bool) bool {
	_ = "STUB: not implemented"
	return false
}
