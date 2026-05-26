package domutils

import (
	"context"

	"golang.org/x/net/html"
)

var ListEndCommentData = "THE END"

func AddListEndComments(ctx context.Context, doc *html.Node) { _ = "STUB: not implemented"; return }

func nameIsList(node *html.Node) bool { _ = "STUB: not implemented"; return false }

func insertComment(listNode *html.Node) { _ = "STUB: not implemented"; return }

func nextNameIsList(startNode *html.Node) bool { _ = "STUB: not implemented"; return false }

// If there is any text between two lists
// they are automatically not connected anymore.

// - - - - //

// A divider already seperates two lists...

// TODO: RunContext.Render()
// -> get acess to keepRemoveMap

// TODO: look in the KeepRemoveMap?
// e.g. ul then script then ul
