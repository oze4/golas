package golas

import (
	"io"
)

// Parse parses las data
func Parse(r io.Reader) LAS {
	p := NewParser(r)
	return p.Parse()
}
