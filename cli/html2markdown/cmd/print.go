package cmd

import (
	"io"
)

type Printer interface {
	Print(w io.Writer)
}

// - - - - - - - //

type coloredBox struct {
	prefix string
	text   string
}

func ColoredBox(prefix string, text string) Printer {
	_ = "STUB: not implemented"
	return *new(Printer)
}

func (p coloredBox) Print(w io.Writer) { _ = "STUB: not implemented"; return }

// - - - - - - - //

type paragraph struct {
	text string
}

func Paragraph(text string) Printer { _ = "STUB: not implemented"; return *new(Printer) }

func (p paragraph) Print(w io.Writer) { _ = "STUB: not implemented"; return }

// - - - - - - - //

type codeBlock struct {
	code string
}

func CodeBlock(code string) Printer { _ = "STUB: not implemented"; return *new(Printer) }

func (cb codeBlock) Print(w io.Writer) {
	_ = "STUB: not implemented"
	// TODO: what about indenting multiline?
	return
}
