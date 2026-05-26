package converter

import (
	"golang.org/x/net/html"
)

func (conv *Converter) handleRenderNodes(ctx Context, w Writer, nodes ...*html.Node) {
	_ = "STUB: not implemented"
	return
}

func (conv *Converter) handleRenderNode(ctx Context, w Writer, node *html.Node) RenderStatus {
	_ = "STUB: not implemented"
	return *new(RenderStatus)
}

// - - A: the #text node - - //

// - - B: the render handlers - - //

// - - C: the fallback - - //
// If nothing works we fallback to this:

func (conv *Converter) handleRenderFallback(ctx Context, w Writer, node *html.Node) RenderStatus {
	_ = "STUB: not implemented"
	return *new(RenderStatus)
}

func (conv *Converter) handleRenderText(ctx Context, w Writer, node *html.Node) RenderStatus {
	_ = "STUB: not implemented"
	return *new(RenderStatus)
}
