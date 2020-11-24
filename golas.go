package golas

import (
	"io"

	"github.com/oze4/golas/pkg/file"
	"github.com/oze4/golas/pkg/parser"
)

// Parse parses las data
func Parse(r io.Reader) file.LASFile {
	p := parser.NewParser(r)
	return p.Parse()
}
