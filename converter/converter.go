package converter

import "sync"

type Converter struct {
	m sync.RWMutex

	err error

	registeredPlugins []string

	preRenderHandlers  prioritizedSlice[HandlePreRenderFunc]
	renderHandlers     prioritizedSlice[HandleRenderFunc]
	postRenderHandlers prioritizedSlice[HandlePostRenderFunc]

	textTransformHandlers prioritizedSlice[HandleTextTransformFunc]

	markdownChars    map[rune]interface{}
	unEscapeHandlers prioritizedSlice[HandleUnEscapeFunc]

	tagTypes map[string]prioritizedSlice[tagType]

	escapeMode escapeMode

	Register register
}

type converterOption = func(c *Converter) error

func NewConverter(opts ...converterOption) *Converter { _ = "STUB: not implemented"; return nil }

type escapeMode string

const (
	EscapeModeDisabled escapeMode = "disabled"
	EscapeModeSmart    escapeMode = "smart"
)

// WithEscapeMode changes the strictness of the "escaping".
//
// Some characters have a special meaning in markdown.
// For example, the character "*" can be used for lists, emphasis and dividers.
// By placing a backlash before that character (e.g. "\*") you can "escape" it.
// Then the character will render as a raw "*" without the "markdown meaning" applied.
//
// Learn more in the documentation
//
//	"disabled" or "smart"
//
//	default: "smart"
func WithEscapeMode(mode escapeMode) converterOption {
	_ = "STUB: not implemented"
	return *new(converterOption)
}
