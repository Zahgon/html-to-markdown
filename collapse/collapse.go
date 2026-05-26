/*

The function to collapse whitespace was adapted from the "turndown" library by Dom Christie,
which was adapted from the "collapse-whitespace" library by Luc Thevenard.

It was ported from Javascript to Golang by Johannes Kaufmann for the use in the "html-to-markdown" library.
To increase performance the use of regex was replaced by custom code.

https://github.com/wooorm/collapse-white-space
https://github.com/mixmark-io/turndown
https://github.com/JohannesKaufmann/html-to-markdown

-----------

MIT License

Copyright (c) 2017 Dom Christie
Copyright (c) 2014 Luc Thevenard <lucthevenard@gmail.com>
Copyright (c) 2018 Johannes Kaufmann

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

// collapse can collapse whitespace in html elements.
//
// It is a port from the Javascript library "turndown" to Golang.
package collapse

import (
	"golang.org/x/net/html"
)

func nextNode(prev *html.Node, current *html.Node, domFuncs *DomFuncs) *html.Node {
	_ = "STUB: not implemented"
	return nil
}

func removeNode(node *html.Node) *html.Node { _ = "STUB: not implemented"; return nil }

type DomFuncs struct {
	IsBlockNode        func(node *html.Node) bool
	IsVoidNode         func(node *html.Node) bool
	IsPreformattedNode func(node *html.Node) bool
}

func fillDefaultDomFuncs(domFuncs *DomFuncs) *DomFuncs { _ = "STUB: not implemented"; return nil }

func Collapse(element *html.Node, domFuncs *DomFuncs) { _ = "STUB: not implemented"; return }

// - - - - - - - - - - - - - - - - - - //

/* node.nodeType == 4 */ // Node.TEXT_NODE or Node.CDATA_SECTION_NODE

// `text` might be empty at this point.

// Node.ELEMENT_NODE

// Avoid trimming space around non-block, non-BR void elements and inline PRE.

// Drop protection if set previously.

// TODO: Is this enough to keep the comments? Does this cause other problems?

// E.g. DoctypeNode
