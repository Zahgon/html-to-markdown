package commonmark

func contains(values []string, searchVal string) bool { _ = "STUB: not implemented"; return false }

// TODO: should this be with the commonmark package? Or more general?
// TODO: Maybe make it an interface? And also have a GetPluginName function?
type ValidateConfigError struct {
	Key   string
	Value string

	// By default is "Key:Value" but can be
	// overriden to e.g. "--key=value"
	KeyWithValue string

	patternDescription string
}

func (e *ValidateConfigError) setDefaultKeyWithValue() { _ = "STUB: not implemented"; return }

func (e *ValidateConfigError) Error() string { _ = "STUB: not implemented"; return "" }

func validateConfig(cfg *config) error { _ = "STUB: not implemented"; return nil }
