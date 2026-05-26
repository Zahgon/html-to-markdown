package domutils

import (
	"context"

	"golang.org/x/net/html"
)

// TODO: make this configurable via the options???
func getMarkdownStructure(name string) string { _ = "STUB: not implemented"; return "" }

// A container block can also contain other blocks.

// Note: "p" would also be part of "leaf_block"

// Leaf blocks can contain inline content
// but NOT other blocks.

// Since these are just placing newlines,
// we dont categorize them.

func headingAlternative(ctx context.Context, node *html.Node) { _ = "STUB: not implemented"; return }

func blockquoteAlternative(ctx context.Context, node *html.Node) { _ = "STUB: not implemented"; return }

func preAlternative(ctx context.Context, node *html.Node) { _ = "STUB: not implemented"; return }

func hrAlternative(ctx context.Context, node *html.Node) { _ = "STUB: not implemented"; return }

// TODO: make this configurable via the options?
var alternatives = map[string]func(ctx context.Context, node *html.Node){
	"h1":         headingAlternative,
	"h2":         headingAlternative,
	"h3":         headingAlternative,
	"h4":         headingAlternative,
	"h5":         headingAlternative,
	"h6":         headingAlternative,
	"blockquote": blockquoteAlternative,
	"pre":        preAlternative,
	"hr":         hrAlternative,
}

func LeafBlockAlternatives(ctx context.Context, doc *html.Node) { _ = "STUB: not implemented"; return }

// A block inside an inline OR a block inside a leaf-block
// is not valid markdown so cannot be rendered.
//
// For example, you cannot place a blockquote inside a heading.
//
// Instead of this weird output (## Heading > My Quote)
// we try to find alternatives (## Heading "My Quote")

// - - - - - - - - - - - - - - - - - - - - - - //
