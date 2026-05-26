package table

import (
	"golang.org/x/net/html"
)

// The content should be at least 1 character wide.
// This also ensures that the table is correctly *recognized* as a markdown table.
const defaultCellWidth = 1

func calculateMaxCounts(rows [][][]byte) []int { _ = "STUB: not implemented"; return nil }

func fillUpRows(rows [][][]byte, maxColumnCount int) [][][]byte {
	_ = "STUB: not implemented"
	return nil
}

func getNumberAttributeOr(node *html.Node, key string, fallback int) int {
	_ = "STUB: not implemented"
	return 0
}

type modification struct {
	y    int
	x    int
	data []byte
}

func calculateModifications(currentRowIndex, currentColIndex, rowSpan, colSpan int, data []byte) []modification {
	_ = "STUB: not implemented"
	return nil
}

// No modification is needed

// Calculate modifications for colspan

// Add modifications for the same row

// Calculate modifications for subsequent rows

func applyGroupedModifications(contents [][][]byte, groupedMods [][]modification) [][][]byte {
	_ = "STUB: not implemented"
	// By applying the modifications in reverse we correctly
	// handle overlapping modifications.
	return nil
}

func applyModifications(contents [][][]byte, mods []modification) [][][]byte {
	_ = "STUB: not implemented"
	return nil

	// Grow on the y axis
}

// Grow on the x axis
// (Note: we only grow x-1 since `Insert` takes care of the rest)

// Now we can do our change:

// growSlice ensures the slice has enough capacity to access the given index.
func growSlice[T any](contents []T, index int, placeholderVal T) []T {
	_ = "STUB: not implemented"
	// Calculate the required growth
	return nil
}

// Grow the slice by appending values

func isEmptyRow(cells [][]byte) bool { _ = "STUB: not implemented"; return false }

func removeEmptyRows(rows [][][]byte) [][][]byte { _ = "STUB: not implemented"; return nil }

// Always keep the first row (the header row)

// If all the rows are empty (including the header row)
// then the table is completely empty...

func removeFirstRowIfEmpty(rows [][][]byte) [][][]byte { _ = "STUB: not implemented"; return nil }

// The first row (the header row) is empty. So lets remove it...
