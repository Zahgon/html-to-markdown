package table

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

func (p *tablePlugin) renderTable(ctx converter.Context, w converter.Writer, n *html.Node) converter.RenderStatus {
	_ = "STUB: not implemented"
	return *new(converter.RenderStatus)
}

// Sometime we just cannot render the table.
// Either because it is an empty table OR
// because there are newlines inside the content (which would break the table).

// Sometimes we pad the cells with extra spaces (e.g. "| text    |").
// For that we first need to know the maximum width of every column.

// Sometimes a row contains less cells that another row.
// We then fill it up with empty cells (e.g. "| text |     |").

// - - - - - - - - - - - - - - - - - - - - - - - - - - //

// - - - Header - - - //

// - - - Body - - - //

// - - - Caption - - - //

// - - - - - - //

func getAlignmentFor(alignments []string, index int) string { _ = "STUB: not implemented"; return "" }

func (s *tablePlugin) writeHeaderUnderline(w converter.Writer, alignments []string, counts []int) {
	_ = "STUB: not implemented"
	return
}

func (s *tablePlugin) writeRow(w converter.Writer, counts []int, cells [][]byte) {
	_ = "STUB: not implemented"
	return
}
