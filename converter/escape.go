package converter

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/marker"
)

const (
	actionKeep   = iota
	actionEscape = iota
)

// IMPORTANT: Only internally we assume it is only byte
var placeholderByte byte = marker.BytesMarkerEscaping[0]

func (conv *Converter) escapeContent(chars []byte) []byte { _ = "STUB: not implemented"; return nil }

// For security reasons, the Unicode character U+0000 must be replaced with the REPLACEMENT CHARACTER (U+FFFD).

func (conv *Converter) unEscapeContent(chars []byte) []byte { _ = "STUB: not implemented"; return nil }

// What to do with this placeholder? Should we escape or not?
