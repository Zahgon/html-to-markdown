package tester

import (
	"flag"
	"testing"
)

var enableRoundTrip = flag.Bool("round", false, "enable the round trip testing")

const suffixInputFile = ".in.html"
const suffixOutputFile = ".out.md"

func getInputFiles(pathOfFolder string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GoldenFiles(t *testing.T, convert ConvertFunc, roundTripConvert ConvertFunc) {
	_ = "STUB: not implemented"
	return
}

// - - - - - - - Golden File Test - - - - - - - //

// Simple, ColoredDiff, ClassicDiff
// goldie.WithDiffEngine(goldie.Simple),

// - - - - - - - Round Trip Test - - - - - - - //

// - - - //

// TODO: clear folder to override earlier runs
// TODO: enable writing using command line
// err2 := res.WriteToFiles("./.tmp/roundtrip")
// if err2 != nil {
// 	t.Error(err2)
// }
