package base

import (
	"strings"

	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"

	"golang.org/x/net/html"
)

type base struct{}

// NewBasePlugin registers a bunch of stuff that is not necessarily related to commonmark,
// like removing nodes, trimming whitespace, collapsing whitespace, ...
func NewBasePlugin() converter.Plugin { _ = "STUB: not implemented"; return *new(converter.Plugin) }

func (s *base) Name() string { _ = "STUB: not implemented"; return "" }

func (b *base) Init(conv *converter.Converter) error { _ = "STUB: not implemented"; return nil }

// "tr" is not in the `IsBlockNode` list,
// but we want to treat is as a block anyway.
// conv.Register.TagStrategy("tr", converter.StrategyMarkdownBlock, converter.PriorityStandard)
// conv.Register.TagType("tr", converter.BlockTagType, converter.PriorityStandard)

// Note: The priority is low, so that collapse runs _after_ all the other functions

func (b *base) preRenderRemove(ctx converter.Context, doc *html.Node) {
	_ = "STUB: not implemented"
	return
}

// Because we are sometimes removing a node, this causes problems
// with the for loop. Using `defer` is a cool trick!
// https://gist.github.com/loopthrough/17da0f416054401fec355d338727c46e

// - - - - - - - //

// After removing elements (see above) it can happen that we have
// two #text nodes right next to each other. This would cause problems
// with the collapse so we merge them together.

func (b *base) preRenderCollapse(ctx converter.Context, doc *html.Node) {
	_ = "STUB: not implemented"
	return
}

var characterEntityReplacer = strings.NewReplacer(
	// We are not using `html.EscapeString` because we
	// care about fewer characters
	"<", "&lt;",
	">", "&gt;",

	// Note: We are not escaping "&" as "&amp;" anymore.
	// In most cases the "&" is completely fine.
	// https://github.com/JohannesKaufmann/html-to-markdown/issues/178
)

func (b *base) handleTextTransform(ctx converter.Context, content string) string {
	_ = "STUB: not implemented"

	// TODO: similar to UnEscapers also only escape if nessesary.
	//       "<" only if not followed by space
	//       "&" only if character entity
	return ""
}

// TODO: reduce conversion between types

func (b *base) postRenderTrimContent(ctx converter.Context, result []byte) []byte {
	_ = "STUB: not implemented"
	// Remove whitespace from the beginning & end
	return nil
}

// Remove too many newlines

func (b *base) postRenderUnescapeContent(ctx converter.Context, result []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}
