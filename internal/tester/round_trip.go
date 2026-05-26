package tester

import (
	"time"

	// "github.com/darmiel/gohtml"
	"github.com/yuin/goldmark"
	goldmarkHtml "github.com/yuin/goldmark/renderer/html"
)

type ConvertFunc func(html []byte) (markdown []byte, err error)

var goldmarkConverter = goldmark.New(
	goldmark.WithRendererOptions(
		// Also render "javascript:" links
		goldmarkHtml.WithUnsafe(),
	),
)

type Result struct {
	Identifier string

	FirstDuration  time.Duration
	SecondDuration time.Duration

	OriginalHtml     []byte
	FirstMarkdown    []byte
	IntermediateHtml []byte
	SecondMarkdown   []byte
}

func (r Result) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (r Result) PrintStatus() { _ = "STUB: not implemented"; return }

func (r Result) WriteToFiles(folderpath string) error { _ = "STUB: not implemented"; return nil }

// originalHtmlPretty := gohtml.Format(string(r.OriginalHtml), true)
// err = ioutil.WriteFile(filepath.Join(folderpath, r.Identifier, "01_pretty.html"), []byte(originalHtmlPretty), 0644)
// if err != nil {
// 	return err
// }

// intermediateHtmlPretty := gohtml.Format(string(r.IntermediateHtml), true)
// err = ioutil.WriteFile(filepath.Join(folderpath, r.Identifier, "03_pretty.html"), []byte(intermediateHtmlPretty), 0644)
// if err != nil {
// 	return err
// }

func RoundTrip(identifier string, originalHtml []byte, convert ConvertFunc) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Hurray, the converter produced exactly the same result. Well done!!!
