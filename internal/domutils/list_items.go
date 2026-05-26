package domutils

import (
	"context"

	"golang.org/x/net/html"
)

// MoveListItems moves non-"li" nodes into the previous "li" nodes.
func MoveListItems(ctx context.Context, n *html.Node) { _ = "STUB: not implemented"; return }

// Collect children to avoid modifying the slice while iterating.

// Skip the node, probably just formatting of code

// We expect that inside an "ol"/"ul" there are *only* "li" nodes.
// But sometimes that is not the case...

// There is a previous "li" node,
// so we move this content into the other "li" node.

// There is no previous "li" node,
// so we wrap this node with it's own "li" node.
