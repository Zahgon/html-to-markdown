package cmd

import (
	"io"
)

type CLIError struct {
	cause    error
	printers []Printer
}

func extractCLIError(err error) (CLIError, bool) {
	_ = "STUB: not implemented"
	return *new(CLIError), false
}

func NewCLIError(cause error, printers ...Printer) error { _ = "STUB: not implemented"; return nil }

func (e CLIError) Error() string { _ = "STUB: not implemented"; return "" }

func (e CLIError) PrintDetails(w io.Writer) { _ = "STUB: not implemented"; return }

// Prepend the error printer

func (cli CLI) PrintErr(err error) { _ = "STUB: not implemented"; return }

func (cli CLI) PrintWarn(err error) { _ = "STUB: not implemented"; return }
