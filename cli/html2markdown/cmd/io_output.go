package cmd

type outputType string

const (
	outputTypeStdout    outputType = "stdout"
	outputTypeDirectory outputType = "directory"
	outputTypeFile      outputType = "file"
)

// The user can indicate that they mean a directory by having a slash as the suffix.
func hasFolderSuffix(outputPath string) bool {
	_ = "STUB: not implemented"
	// Note: We generally support the os.PathSeparator (e.g. "\" on windows).
	//
	//	But also "/" is always supported.
	return false
}

func determineOutputType(_inputPath string, countInputs int, outputPath string) (outputType, error) {
	_ = "STUB: not implemented"
	return *new(outputType), nil
}

// - - - - - - - - - //
// We can now assume that the output path specifies a file.
// But let's make sure...

// There are multiple inputs, so the input MUST have been a glob or directory.
// It also means that the output MUST be a directory.

// TODO: The glob can also be a folder with just one file...
//       So we should check if the path contains any glob characters.

// Check if output path exists

// With a file extension it is LIKELY to be a file.

// Default to file for single input

func calculateOutputPaths(inputFilepath string, inputs []*input) error {
	_ = "STUB: not implemented"
	return nil
}

// -> The standard filename

// We hash the relative path (based from the globBase)
// since the globBase is *the same* for all files.
// Bonus: It makes testing easier as the temporary folder does not matter.

// -> The filename for duplicates

func hashFilepath(path string) string { _ = "STUB: not implemented"; return "" }

// Ensure that regardless of operating system the path has the same format.
// Bonus: Easier testing as we always hash the same.

func ensureOutputDirectories(outputType outputType, outputFilepath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cli *CLI) writeOutput(outputType outputType, filename string, markdown []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteFile writes data to a file with override control
// If override is false and file exists, returns an error
// If override is true, truncates existing file or creates new one
func WriteFile(filename string, data []byte, override bool) error {
	_ = "STUB: not implemented"
	// As the base flags we have:
	//   O_WRONLY = write to the file, not read
	//   O_CREATE = create the file if it doesn't exist
	return nil
}

// We add this flag:
//   O_TRUNC = the existing contents are truncated to zero length

// We add this flag:
//   O_EXCL = if used with O_CREATE, causes error if file already exists
