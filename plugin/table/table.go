package table

import (
	"sync"

	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

type option func(p *tablePlugin) error

type SpanCellBehavior string

const (
	// SpanBehaviorEmpty renders an empty cell.
	SpanBehaviorEmpty SpanCellBehavior = "empty"
	// SpanBehaviorMirror renders the same content as the original cell.
	SpanBehaviorMirror SpanCellBehavior = "mirror"
)

// WithSpanCellBehavior configures how cells affected by colspan/rowspan attributes
// should be rendered. When a cell spans multiple columns or rows, the affected cells
// can either be empty or contain the same content as the original cell.
func WithSpanCellBehavior(behavior SpanCellBehavior) option {
	_ = "STUB: not implemented"
	return *new(option)
}

// TODO: should we allow empty string?

type NewlineBehavior string

const (
	// NewlineBehaviorSkip skips tables with newlines in cells (default).
	NewlineBehaviorSkip NewlineBehavior = "skip"
	// NewlineBehaviorPreserve preserves newlines in cells.
	NewlineBehaviorPreserve NewlineBehavior = "preserve"
)

// WithNewlineBehavior configures how to handle newlines in table cells.
// When set to NewlineBehaviorSkip (default), tables with newlines in cells are skipped.
// When set to NewlineBehaviorPreserve, newlines are preserved in cells.
//
// Markdown tables don't support multiline content by default, so this provides a workaround to still convert tables with newlines.
func WithNewlineBehavior(behavior NewlineBehavior) option {
	_ = "STUB: not implemented"
	return *new(option)
}

// Allow empty string to default to Skip

type CellPaddingBehavior string

const (
	// CellPaddingBehaviorAligned adds visual padding to cells to make each column equal width (default).
	CellPaddingBehaviorAligned CellPaddingBehavior = "aligned"
	// CellPaddingBehaviorMinimal keeps a very small amount of padding to balance table readability while also reducing character count.
	CellPaddingBehaviorMinimal CellPaddingBehavior = "minimal"
	// CellPaddingBehaviorNone refrains from adding the padding to the cells.
	CellPaddingBehaviorNone CellPaddingBehavior = "none"
)

// WithCellPaddingBehavior configures how to handle padding in table cells.
// When set to "aligned" (default), every cell's text is padded to the width of the largest cell in its column.
// When set to "minimal", every cell gets a space at the beginning and end of the cell for some minimal padding.
// When set to "none", no extra padding is applied to cells.
func WithCellPaddingBehavior(behavior CellPaddingBehavior) option {
	_ = "STUB: not implemented"
	return *new(option)
}

// Allow empty string to default to "aligned"

// WithSkipEmptyRows configures the table plugin to omit empty rows from the output.
// An empty row is defined as a row where all cells contain no content or only whitespace.
// When set to true, empty rows will be omitted from the output. When false (default),
// all rows are preserved.
func WithSkipEmptyRows(skip bool) option { _ = "STUB: not implemented"; return *new(option) }

// WithHeaderPromotion configures whether the first row should be treated as a header
// when the table has no explicit header row (e.g. <th> elements). When set to true, the
// first row will be converted to a header row with separator dashes. When false (default),
// all rows are treated as regular content.
func WithHeaderPromotion(promote bool) option { _ = "STUB: not implemented"; return *new(option) }

// WithPresentationTables configures whether tables marked with role="presentation"
// should be converted to markdown. When set to true, presentation tables will be
// converted like regular tables. When false (default), these tables are skipped
// since they typically represent layout rather than semantic content.
func WithPresentationTables(convert bool) option { _ = "STUB: not implemented"; return *new(option) }

type tablePlugin struct {
	m   sync.RWMutex
	err error

	spanCellBehavior          SpanCellBehavior
	newlineBehavior           NewlineBehavior
	skipEmptyRows             bool
	promoteFirstRowToHeader   bool
	convertPresentationTables bool
	cellPaddingBehavior       CellPaddingBehavior
}

func (p *tablePlugin) setError(err error) { _ = "STUB: not implemented"; return }

func (p *tablePlugin) getError() error { _ = "STUB: not implemented"; return nil }

func NewTablePlugin(opts ...option) converter.Plugin {
	_ = "STUB: not implemented"
	return *new(converter.Plugin)
}

func (s *tablePlugin) Name() string { _ = "STUB: not implemented"; return "" }

func (s *tablePlugin) Init(conv *converter.Converter) error { _ = "STUB: not implemented"; return nil }

// Any error raised from the option func

func (s *tablePlugin) handleRender(ctx converter.Context, w converter.Writer, n *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}

// Normally, when the "table" gets rendered we do NOT go into this case.
// But as a fallback we separate the rows through newlines.

func (s *tablePlugin) renderFallbackRow(ctx converter.Context, w converter.Writer, n *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}
