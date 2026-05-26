package htmltomarkdown

import (
	"io"

	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

// ConvertString converts a html-string to a markdown-string.
//
// Under the hood `html.Parse()` is used to parse the HTML.
func ConvertString(htmlInput string, opts ...converter.ConvertOptionFunc) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ConvertReader converts the html from the reader to markdown.
//
// Under the hood `html.Parse()` is used to parse the HTML.
func ConvertReader(r io.Reader, opts ...converter.ConvertOptionFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConvertNode converts a `*html.Node` to a markdown byte slice.
//
// If you have already parsed an HTML page using the `html.Parse()` function
// from the "golang.org/x/net/html" package then you can pass this node
// directly to the converter.
func ConvertNode(doc *html.Node, opts ...converter.ConvertOptionFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
