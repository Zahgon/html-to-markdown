package cmd

type input struct {
	inputFullFilepath  string
	outputFullFilepath string
	data               []byte
}

// E.g. "website.html" -> "website"
func fileNameWithoutExtension(fileName string) string { _ = "STUB: not implemented"; return "" }

// If the output is a file, it would be "output.md"
var defaultBasename = "output"

func (cli *CLI) listInputs() ([]*input, error) {
	_ = "STUB: not implemented"

	// NOTE: When both stdin and --input are specified,
	// the explicit --file argument takes precedence.
	// This improves interoperability with other tools like `xargs`.
	// https://github.com/JohannesKaufmann/html-to-markdown/issues/170
	return nil, nil
}

// The inputFilepath wasn't actually a glob but was pointing to an existing folder.
// The user probably wanted to convert all files in that folder — so we recommend the glob.

func (cli *CLI) readInput(in *input) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
