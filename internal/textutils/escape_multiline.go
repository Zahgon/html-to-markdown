package textutils

var (
	doubleSpace = []byte{' ', ' '}

	newlineBreak              = []byte{'\n'}
	hardLineBreak             = []byte{' ', ' ', '\n'}
	escapedNoContentLineBreak = []byte{'\\', '\n'}
)

// EscapeMultiLine deals with multiline content inside a link or a heading.
func EscapeMultiLine(content []byte) []byte { _ = "STUB: not implemented"; return nil }

// A blank line would interrupt the link.
// So we need to escape the line

// For the last line we don't need to add any "\n" anymore

// Now decide what ending we want:

// We already have "  " so adding a "\n" is enough

// We *prefer* having a hard-line-break "  \n"
