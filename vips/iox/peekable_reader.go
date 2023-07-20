package iox

import (
	"io"
)

type PeekableReader interface {
	io.Reader
	Peek(n int) (nbytes []byte, err error)
}

type RewindReader interface {
	PeekableReader
	Rewind() bool
}
