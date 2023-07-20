package iox

import (
	"bufio"
	"io"
	"os"
	"runtime"
)

type bufferedFileReader struct {
	fileHandle *os.File
	reader     *bufio.Reader
}

type bufferedFileReaderOpt func(bf *bufferedFileReader)

func WithDefaultFileBuf() bufferedFileReaderOpt {
	return func(bf *bufferedFileReader) {
		bf.reader = bufio.NewReader(bf.fileHandle)
	}
}

func WithSizedFileBuf(bufSize int) bufferedFileReaderOpt {
	return func(bf *bufferedFileReader) {
		bf.reader = bufio.NewReaderSize(bf.fileHandle, bufSize)
	}
}

var (
	applyDefaultFileBuf = WithDefaultFileBuf()
)

func NewBufferedFileReader(path string, opts ...bufferedFileReaderOpt) (RewindReader, error) {
	fin, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	bf := &bufferedFileReader{
		fileHandle: fin,
	}

	for _, opt := range opts {
		opt(bf)
	}

	if bf.reader == nil {
		applyDefaultFileBuf(bf)
	}

	runtime.SetFinalizer(bf, finalizeBufferedFileReader)
	return bf, err
}

func finalizeBufferedFileReader(bf *bufferedFileReader) {
	bf.Close()
}

func (bf *bufferedFileReader) Close() error {
	return bf.fileHandle.Close()
}

func (bf *bufferedFileReader) Read(p []byte) (n int, err error) {
	return bf.reader.Read(p)
}

func (bf *bufferedFileReader) Peek(n int) (nbytes []byte, err error) {
	return bf.reader.Peek(n)
}

func (bf *bufferedFileReader) Rewind() bool {
	bf.fileHandle.Seek(0, io.SeekStart)
	bf.reader.Reset(bf.fileHandle)
	return true
}
