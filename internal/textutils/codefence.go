package textutils

func CalculateCodeFenceOccurrences(fenceChar rune, content string) int {
	_ = "STUB: not implemented"
	return 0
}

// We encountered a fence character, now count how many
// are directly afterwards

// If the last element in the content was a fenceChar

// CalculateCodeFence can be passed the content of a code block and it returns
// how many fence characters (` or ~) should be used.
//
// This is useful if the html content includes the same fence characters
// for example ```
// -> https://stackoverflow.com/a/49268657
func CalculateCodeFence(fenceChar rune, content string) string {
	_ = "STUB: not implemented"
	return ""
}

// The outer fence block always has to have
// at least one character more than any content inside

// You have to have at least three fence characters
// to be recognized as a code block

func findMax(a []int) (max int) { _ = "STUB: not implemented"; return 0 }
