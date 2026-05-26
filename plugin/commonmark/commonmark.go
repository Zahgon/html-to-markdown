package commonmark

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
)

type commonmark struct {
	config
}

type OptionFunc = func(config *config)

// _ or *
//
// default: *
func WithEmDelimiter(delimiter string) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// ** or __
//
// default: **
func WithStrongDelimiter(delimiter string) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// Any Thematic break
//
// default: "* * *"
func WithHorizontalRule(rule string) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// "-", "+", or "*"
//
// default: "-"
func WithBulletListMarker(marker string) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

func WithListEndComment(enabled bool) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// ``` or ~~~
//
// default: ```
func WithCodeBlockFence(fence string) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// "setext" or "atx"
//
// default: "atx"
func WithHeadingStyle(style headingStyle) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// WithLinkEmptyHrefBehavior configures how links with *empty hrefs* are rendered.
// Take for example:
//
//	<a href="">the link content</a>
//
// LinkBehaviorRenderAsLink would result in "[the link content]()"
//
// LinkBehaviorSkipLink would result in "the link content"
func WithLinkEmptyHrefBehavior(behavior linkRenderingBehavior) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// WithLinkEmptyContentBehavior configures how links *without content* are rendered.
// Take for example:
//
//	<a href="/page"></a>
//
// LinkBehaviorRenderAsLink would result in "[](/page)"
//
// LinkBehaviorSkipLink would result in an empty string.
func WithLinkEmptyContentBehavior(behavior linkRenderingBehavior) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// TODO: allow changing the link style once the render logic is implemented
//
// "inlined" or "referenced_index" or "referenced_short"
//
// default: inlined
// func WithLinkStyle(style linkStyle) OptionFunc {
// 	return func(config *config) {
// 		config.LinkStyle = style
// 	}
// }

// NewCommonmarkPlugin registers the markdown syntax of commonmark.
func NewCommonmarkPlugin(opts ...OptionFunc) converter.Plugin {
	_ = "STUB: not implemented"
	return *new(converter.Plugin)
}

func (s *commonmark) Name() string { _ = "STUB: not implemented"; return "" }

func (cm *commonmark) Init(conv *converter.Converter) error { _ = "STUB: not implemented"; return nil }

// - - - - - - - - //

// Note: Should run after "collapse" & also after "remove"

// Early return if the feature is unwanted

func (cm commonmark) handlePostRenderCodeBlockNewline(ctx converter.Context, content []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (cm commonmark) handleTextTransform(ctx converter.Context, content string) string {
	_ = "STUB: not implemented"
	return ""
}

// if isEnabled, ok := ctx.Value("is_inside_heading").(bool); ok && isEnabled {
// 	// The "#" character would be completely removed, if at the _end_
// 	// of the heading content. So always escape it inside headings.
// 	content = strings.Replace(content, string(marker.MarkerEscaping)+`#`, `\#`, -1)
// }
