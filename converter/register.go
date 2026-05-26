package converter

import (
	"golang.org/x/net/html"
)

type register struct {
	conv *Converter
}

func (r *register) Plugin(plugin Plugin) { _ = "STUB: not implemented"; return }

// - - - - - - - - - - - - - Pre-Render - - - - - - - - - - - - - //

type HandlePreRenderFunc func(ctx Context, doc *html.Node)

func (r *register) PreRenderer(fn HandlePreRenderFunc, priority int) {
	_ = "STUB: not implemented"
	return
}

func (conv *Converter) getPreRenderHandlers() prioritizedSlice[HandlePreRenderFunc] {
	_ = "STUB: not implemented"
	return nil
}

// - - - - - - - - - - - - - Render - - - - - - - - - - - - - //

// Writer is an interface that only conforms to the Write* methods of bytes.Buffer
type Writer interface {
	Write(p []byte) (n int, err error)
	WriteByte(c byte) error
	WriteRune(r rune) (n int, err error)
	WriteString(s string) (n int, err error)
}

type HandleRenderFunc func(ctx Context, w Writer, n *html.Node) RenderStatus

func (r *register) Renderer(fn HandleRenderFunc, priority int) { _ = "STUB: not implemented"; return }

// RendererFor registers a renderer for a specific tag (e.g. "div").
// It is a small wrapper around `TagType()` and `Renderer()`.
func (r *register) RendererFor(tagName string, tagType tagType, renderFn HandleRenderFunc, priority int) {
	_ = "STUB: not implemented"

	// 1. we add the "tagType" to the map
	return
}

// 2. we register the render function

func (conv *Converter) getRenderHandlers() prioritizedSlice[HandleRenderFunc] {
	_ = "STUB: not implemented"
	return nil
}

// - - - - - - - - - - - - - Post Render - - - - - - - - - - - - - //

type HandlePostRenderFunc func(ctx Context, content []byte) []byte

func (r *register) PostRenderer(fn HandlePostRenderFunc, priority int) {
	_ = "STUB: not implemented"
	return
}

func (conv *Converter) getPostRenderHandlers() prioritizedSlice[HandlePostRenderFunc] {
	_ = "STUB: not implemented"
	return nil
}

// - - - - - - - - - - - - - Text - - - - - - - - - - - - - //

type HandleTextTransformFunc func(ctx Context, content string) string

func (r *register) TextTransformer(fn HandleTextTransformFunc, priority int) {
	_ = "STUB: not implemented"
	return
}

func (conv *Converter) getTextTransformHandlers() prioritizedSlice[HandleTextTransformFunc] {
	_ = "STUB: not implemented"
	return nil
}

// - - - - - - - - - - - - - Escaping - - - - - - - - - - - - - //

func (r *register) EscapedChar(chars ...rune) { _ = "STUB: not implemented"; return }

func (conv *Converter) checkIsEscapedChar(r rune) bool { _ = "STUB: not implemented"; return false }

type HandleUnEscapeFunc func(chars []byte, index int) int

func (r *register) UnEscaper(fn HandleUnEscapeFunc, priority int) {
	_ = "STUB: not implemented"
	return
}

func (conv *Converter) getUnEscapeHandlers() prioritizedSlice[HandleUnEscapeFunc] {
	_ = "STUB: not implemented"
	return nil
}

// - - - - - - - - - - - - - Tag Type - - - - - - - - - - - - - //

type tagType string

const (
	TagTypeBlock  tagType = "block"
	TagTypeInline tagType = "inline"

	// TagTypeRemove will remove that node in the _PreRender_ phase with a high priority.
	TagTypeRemove tagType = "remove"
)

func (r *register) TagType(tagName string, tagType tagType, priority int) {
	_ = "STUB: not implemented"
	return
}

func (conv *Converter) getTagType(tagName string) (tagType, bool) {
	_ = "STUB: not implemented"
	return *new(tagType), false
}
