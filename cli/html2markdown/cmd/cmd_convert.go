package cmd

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/commonmark"
	"golang.org/x/net/html"
)

func overrideValidationError(e *commonmark.ValidateConfigError) error {
	_ = "STUB: not implemented"

	// TODO: Maybe OptionFunc should already validate and return an error?
	//       Then it would be easier to override the Key since we have once
	//       place to assemble the []OptionFunc and directly treat the errors...
	//
	// We would basically invoke it ourselves:
	//    err := commonmark.WithStrongDelimiter(cli.config.strongDelimiter)(conv)
	return nil
}

func (cli *CLI) includeNodesFromDoc(doc *html.Node) (*html.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *CLI) excludeNodesFromDoc(doc *html.Node) error { _ = "STUB: not implemented"; return nil }

// Because we are sometimes removing a node, this causes problems
// with the for loop. Using `defer` is a cool trick!
// https://gist.github.com/loopthrough/17da0f416054401fec355d338727c46e

func (cli *CLI) parseInputWithSelectors(input []byte) (*html.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *CLI) convert(input []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
