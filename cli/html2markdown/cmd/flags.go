package cmd

import (
	"github.com/andybalholm/cascadia"
)

// selectorFlag sets up a flag that parses a CSS selector string into a cascadia.Selector.
func (cli *CLI) selectorFlag(target *cascadia.SelectorGroup, name string, usage string) {
	_ = "STUB: not implemented"
	return
}

// Compile the provided CSS selector string

func (cli *CLI) singleStringFlag(target *string, name string, usage string) {
	_ = "STUB: not implemented"
	return
}

func (cli *CLI) initFlags(progname string) { _ = "STUB: not implemented"; return }

// - - - - - General - - - - - //

// TODO: --tag-type-block=script,style (and check that it is not a selector)
// TODO: --tag-type-inline=script,style (and check that it is not a selector)

// - - - - - Options - - - - - //

// - - - - - Plugins - - - - - //
// TODO: --opt-strikethrough-delimiter for the strikethrough plugin

func (cli *CLI) parseFlags(args []string) error { _ = "STUB: not implemented"; return nil }

// Validate flag dependencies

// TODO: use constant for flag name & use formatFlag
//       var keyStrongDelimiter = "opt-strong-delimiter"
