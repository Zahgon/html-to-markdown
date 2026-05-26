package textutils

// DelimiterForEveryLine puts the delimiter not just at the start and end of the string
// but if the text is divided on multiple lines, puts the delimiters on every line with content.
//
// Otherwise the bold/italic delimiters won't be recognized if it contains new line characters.
func DelimiterForEveryLine(text []byte, delimiter []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// For empty lines, we don't need a delimiter

// To join the lines again, add a newlines character
