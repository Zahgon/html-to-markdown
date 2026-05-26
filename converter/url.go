package converter

import (
	"net/url"
	"strings"
)

var percentEncodingReplacer = strings.NewReplacer(
	" ", "%20",
	"[", "%5B",
	"]", "%5D",
	"(", "%28",
	")", "%29",
	"<", "%3C",
	">", "%3E",
)

func parseBaseDomain(rawDomain string) *url.URL { _ = "STUB: not implemented"; return nil }

// Yes, we got valid domain (probably with a http/https scheme)

// Yes, we got a valid domain (by choosing a fallback scheme)

func defaultAssembleAbsoluteURL(tagName string, rawURL string, domain string) string {
	_ = "STUB: not implemented"
	return ""
}

// Golangs url.Parse does not seem to distinguish between
// no fragment and an empty fragment.

// Increase the chance that the url will be parsed

// We can't do anything with this url because it is invalid

// This is a data uri (for example an inline base64 image)

// The default Query().Encode() encodes the query parameters "sorted by key".
// Instead we want to keep the original order, but still encode the parameters.

// For better compatibility (especially in regards to mailto links),
// instead of encoding a space with a "+" we use ""%20" to prevent
// e.g. the email reading "Hi+Johannes" instead of "Hi Johannes"

// If a "domain" is provided, we use that to convert relative links
// to absolute links.

// - - - - //

func decodeAndEncode(original string) string { _ = "STUB: not implemented"; return "" }

func ParseAndEncodeQuery(rawQuery string) string { _ = "STUB: not implemented"; return "" }

// A: Just the key

// B: The key and the equal sign

// C: The key and the equal sign and the value
