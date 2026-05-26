package cmd

import (
	"bytes"
	"flag"
	"os"

	"github.com/andybalholm/cascadia"
)

var (
	projectBinary = "html2markdown"
)

// OsExiter is the function used when the app exits. If not set defaults to os.Exit.
var OsExiter = os.Exit

// - - - - - - - - - - - - - //

type Config struct {
	// args are the positional (non-flag) command-line arguments.
	args []string

	inputFilepath   string
	outputFilepath  string
	outputOverwrite bool

	// - - - - - General - - - - - //
	version bool
	domain  string

	includeSelector cascadia.SelectorGroup
	excludeSelector cascadia.SelectorGroup

	// - - - - - Options - - - - - //
	strongDelimiter string

	// - - - - - Plugins - - - - - //
	enablePluginStrikethrough bool

	enablePluginTable        bool
	tableSkipEmptyRows       bool
	tableHeaderPromotion     bool
	tableSpanCellBehavior    string
	tablePresentationTables  bool
	tableNewlineBehavior     string
	tableCellPaddingBehavior string
}

// Release holds the information (from the 3 ldflags) that goreleaser sets.
type Release struct {
	// Current Git tag (the v prefix is stripped)
	Version string

	// Current git commit SHA
	Commit string

	// Date in the RFC3339 format
	Date string
}
type CLI struct {
	Stdin  ReadWriterWithStat
	Stdout ReadWriterWithStat
	Stderr ReadWriterWithStat

	OsArgs []string

	Release Release

	isStdinPipe  bool
	isStdoutPipe bool
	isStderrPipe bool

	flags  *flag.FlagSet
	config Config

	usageText bytes.Buffer
}

func (cli *CLI) Init() error { _ = "STUB: not implemented"; return nil }

func (cli *CLI) Execute() { _ = "STUB: not implemented"; return }

// General Error

func (cli *CLI) run() ([]error, error) { _ = "STUB: not implemented"; return nil, nil }

// - - - - - - - - - - - - - - - //
