package table

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

type tableContent struct {
	Alignments []string
	Rows       [][][]byte
	Caption    []byte
}

func containsNewline(b []byte) bool { _ = "STUB: not implemented"; return false }

func hasProblematicChildNode(node *html.Node) bool { _ = "STUB: not implemented"; return false }

// This will be caught with the newline check anyway.
// But we can safe some effort by aborting early...

func hasProblematicParentNode(node *html.Node) bool { _ = "STUB: not implemented"; return false }

func (p *tablePlugin) collectTableContent(ctx converter.Context, node *html.Node) *tableContent {
	_ = "STUB: not implemented"
	return nil
}

// In HTML-Emails many tables are used. Oftentimes these tables are nested
// which is not possible with markdown. But these tables are mostly used
// for *layout purposes* rather than displaying actual tabular data.

// So lets skip those with role="presentation" and focus on real tables...

// There are certain nodes (e.g. <hr />) that cannot be in a table.
// If we found one, we unfortunately cannot convert the table.
//
// Note: It is okay for a block node (e.g. <div>) to be in a table.
//       However once it causes multiple lines, it does not work anymore.
//       For that we have the `containsNewline` check below.

// There are certain parent nodes (e.g. <a>) that cannot contain a table.
// We would break the rendering of the link, so we unfortunately cannot convert the table.

// Replace newlines with <br /> tags

// We're configured to skip tables with newlines, return nil

// Sometimes a cell wants to *span* over multiple columns or/and rows.
// What should be displayed in those other cells?
// Render exactly the same content OR an empty string?
func (p *tablePlugin) getContentForMergedCell(originalContent []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func getFirstNode(node *html.Node, nodes ...*html.Node) *html.Node {
	_ = "STUB: not implemented"
	return nil
}

func collectAlignments(headerRowNode *html.Node, rowNodes []*html.Node) []string {
	_ = "STUB: not implemented"
	return nil
}

func (p *tablePlugin) collectCellsInRow(ctx converter.Context, rowIndex int, rowNode *html.Node) ([][]byte, []modification) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The character "|" inside the content would mistakenly be recognized as part of the table. So we have to escape it.

// - - col / row span - - //

func (p *tablePlugin) collectRows(ctx converter.Context, headerRowNode *html.Node, rowNodes []*html.Node) [][][]byte {
	_ = "STUB: not implemented"
	return nil
}

// - - 1. the header row - - //

// There needs to be *header* row so that the table is recognized.
// So it is better to have an empty header row...

// - - 2. the normal rows - - //

// Sometimes a cell wants to *span* over multiple columns or/and rows.
// We collected these modifications and are now applying it,
// by shifting the cells around.

func collectCaption(ctx converter.Context, node *html.Node) []byte {
	_ = "STUB: not implemented"
	return nil
}
