package collapse

func byteSliceToString(b []byte) string {
	_ = "STUB: not implemented"
	/* #nosec G103 */ return ""
}

func replaceAnyWhitespaceWithSpace(source string) string { _ = "STUB: not implemented"; return "" }

// Some performance optimizations:
// - If no replacement was done, we return the original slice and dont allocate.
// - We batch appends

// Start of newlines

// Middle of newlines

// Character after the last newline character

// There was only one `isWhitespace` match & that is a space.
// So the replacement would be exactly the same...

// a) no changes need to be done

// b) Only the normal characters until the end still need to be added

// c) There is a match, but it is exactly the same as the replacement
//    If there is no new slice, we can skip the replacement.

// d) The match still needs to be replaced (and possible the previous normal characters be added)

// Huray, we did not do any allocations with make()
// and instead just return the original slice.
