package strikethrough

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

type option func(p *strikethroughPlugin)

func WithDelimiter(delimiter string) option { _ = "STUB: not implemented"; return *new(option) }

type strikethroughPlugin struct {
	delimiter string
}

// Strikethrough converts `<strike>`, `<s>`, and `<del>` elements
func NewStrikethroughPlugin(opts ...option) converter.Plugin {
	_ = "STUB: not implemented"
	return *new(converter.Plugin)
}

func (s *strikethroughPlugin) Name() string { _ = "STUB: not implemented"; return "" }

func (s *strikethroughPlugin) Init(conv *converter.Converter) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *strikethroughPlugin) handlePreRender(ctx converter.Context, doc *html.Node) {
	_ = "STUB: not implemented"
	return
}

func (s *strikethroughPlugin) handleUnEscapers(chars []byte, index int) int {
	_ = "STUB: not implemented"
	return 0
}

// "not followed by Unicode whitespace"

func nameIsStrikethough(node *html.Node) bool { _ = "STUB: not implemented"; return false }

func nameIsBothStrikethough(a *html.Node, b *html.Node) bool {
	_ = "STUB: not implemented"
	return false
}

func (s strikethroughPlugin) handleRender(ctx converter.Context, w converter.Writer, n *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}

func (s strikethroughPlugin) renderStrikethrough(ctx converter.Context, w converter.Writer, n *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}

// If there is a newline character between the start and end delimiter
// the delimiters won't be recognized. Either we remove all newline characters
// OR on _every_ line we put start & end delimiters.
