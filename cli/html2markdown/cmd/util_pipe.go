package cmd

import (
	"io"
	"io/fs"
)

type ReadWriterWithStat interface {
	io.ReadWriter

	Stat() (fs.FileInfo, error)
}

func isPipe(f ReadWriterWithStat) (bool, error) { _ = "STUB: not implemented"; return false, nil }
