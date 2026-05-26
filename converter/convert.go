package converter

import (
	"context"
	"errors"
	"io"

	"golang.org/x/net/html"
)

type convertOption struct {
	domain  string
	context context.Context
}
type ConvertOptionFunc func(o *convertOption)

func WithContext(ctx context.Context) ConvertOptionFunc {
	_ = "STUB: not implemented"
	return *new(ConvertOptionFunc)
}

// WithDomain provides a base `domain` to the converter and
// to the `AssembleAbsoluteURL` function.
//
// If a *relative* url is encountered (in an image or link) then the `domain` is used
// to convert it to a *absolute* url.
func WithDomain(domain string) ConvertOptionFunc {
	_ = "STUB: not implemented"
	return *new(ConvertOptionFunc)
}

func (conv *Converter) setError(err error) { _ = "STUB: not implemented"; return }

func (conv *Converter) getError() error { _ = "STUB: not implemented"; return nil }

var errNoRenderHandlers = errors.New(`no render handlers are registered. did you forget to register the "commonmark" and "base" plugins?`)
var errBasePluginMissing = errors.New(`you registered the "commonmark" plugin but the "base" plugin is also required`)

// ConvertNode converts a `*html.Node` to a markdown byte slice.
//
// If you have already parsed an HTML page using the `html.Parse()` function
// from the "golang.org/x/net/html" package then you can pass this node
// directly to the converter.
func (conv *Converter) ConvertNode(doc *html.Node, opts ...ConvertOptionFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// There can be errors while calling `Init` on the plugins (e.g. validation errors).
// Now is the first opportunity where we can return that error.

// If there are no render handlers registered this is
// usually a user error - since people want the Commonmark Plugin in 99% of cases.

// - - - - - - - - - - - - - - - - - - - //

// - - - - - - - - - - - - - - - - - - - //

// Pre-Render

// Render

// Post-Render

// ConvertReader converts the html from the reader to markdown.
//
// Under the hood `html.Parse()` is used to parse the HTML.
func (conv *Converter) ConvertReader(r io.Reader, opts ...ConvertOptionFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConvertString converts a html-string to a markdown-string.
//
// Under the hood `html.Parse()` is used to parse the HTML.
func (conv *Converter) ConvertString(htmlInput string, opts ...ConvertOptionFunc) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
