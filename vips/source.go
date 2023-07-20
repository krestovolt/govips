package vips

// #cgo pkg-config: vips
// #include "source.h"
import "C"
import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/davidbyttow/govips/v2/vips/iox"
)

type Source struct {
	reader iox.PeekableReader
	seeker io.Seeker
	closer io.Closer
	src    *C.struct__VipsSourceCustom
	// read signal handler id
	rsigHandle C.gulong
	// seek signal handler id
	ssigHandle C.gulong
	lock       sync.Mutex
	closeOnce  sync.Once
}

// NewSource creates a new image source that uses a `iox.PeekableReader` (e.g. bufio.Reader)
//
// By default it will attach a `read` signal and call the `Read` method of the reader,
// if the reader also supports the `io.Seeker` then the `seek` signal handler would also be attached.
//
// The needs of `iox.PeekableReader` interface determined by how the internal load function extract the `vips.ImageType` from the actual source's reader.
// The load function needs to get the first 12 bytes without advancing the reader position, this makes it possible to prevent reading invalid buffer that can cause a whole application to crash (or segvault-ing).
func NewSource(image iox.PeekableReader) *Source {
	src := &Source{}
	srcPtr := (C.gpointer)(unsafe.Pointer(src))

	// Initiate this step first, since srcPtr still clear from any go pointer
	src.src = C.create_go_custom_source()
	src.rsigHandle = C.connect_go_signal_read(src.src, srcPtr)
	src.ssigHandle = C.connect_go_signal_seek(src.src, srcPtr)

	src.reader = image

	skr, ok := image.(io.Seeker)
	if ok {
		src.seeker = skr
	}

	clr, ok := image.(io.Closer)
	if ok {
		src.closer = clr
	}

	runtime.SetFinalizer(src, finalizeSource)

	return src
}

func finalizeSource(src *Source) {
	govipsLog("govips", LogLevelDebug, fmt.Sprintf("closing image %p", src))
	if src != nil {
		src.Close()
	}
}

func (s *Source) Close() {
	// s.lock.Lock()
	// defer s.lock.Unlock()
	s.closeOnce.Do(func() {
		govipsLog("govips", LogLevelInfo, fmt.Sprintf("Closing source %p", s))

		if s.src != nil {
			C.free_go_custom_source(s.src, s.rsigHandle, s.ssigHandle)
			s.src = nil
		}

		if s.closer != nil {
			s.closer.Close()
			s.closer = nil
		}
		s.reader = nil
		s.seeker = nil
	})
}

//export goSourceRead
func goSourceRead(goSource C.gpointer, buffer unsafe.Pointer, bufSize C.gint64) (read C.gint64) {
	src := (*Source)(unsafe.Pointer(goSource))
	if src.reader == nil {
		return -1
	}

	// https://stackoverflow.com/questions/51187973/how-to-create-an-array-or-a-slice-from-an-array-unsafe-pointer-in-golang
	sh := &reflect.SliceHeader{
		Data: uintptr(buffer),
		Len:  int(bufSize),
		Cap:  int(bufSize),
	}
	buf := *(*[]byte)(unsafe.Pointer(sh))

	n, err := src.reader.Read(buf)
	if errors.Is(err, io.EOF) {
		govipsLog("govips", LogLevelDebug, fmt.Sprintf("goSourceRead[id %p]: EOF [read %d]", src, n))
		return C.gint64(n)
	} else if err != nil {
		govipsLog("govips", LogLevelError, fmt.Sprintf("goSourceRead[id %p]: Error: %v [read %d]", src, err, n))
		return -1
	}

	govipsLog("govips", LogLevelDebug, fmt.Sprintf("goSourceRead[id %p]: OK [read %d]", src, n))
	return C.gint64(n)
}

//export goSourceSeek
func goSourceSeek(goSource C.gpointer, offset C.gint64, cWhence C.int) (newOffset C.gint64) {
	src := (*Source)(unsafe.Pointer(goSource))

	if src.seeker == nil {
		govipsLog("govips", LogLevelDebug, fmt.Sprintf("goSourceRead[id %p]: Seek not supported", src))
		return -1 // Unsupported!
	}

	whence := int(cWhence)
	switch whence {
	case io.SeekStart, io.SeekCurrent, io.SeekEnd:
	default:
		govipsLog("govips", LogLevelError, fmt.Sprintf("goSourceSeek[id %p]: Invalid whence value [%d]", src, whence))
		return -1
	}

	n, err := src.seeker.Seek(int64(offset), whence)
	if err != nil {
		govipsLog("govips", LogLevelError, fmt.Sprintf("goSourceSeek[id %p]: Error: %v [offset %d | whence %d]", src, err, n, whence))
		return -1
	}

	govipsLog("govips", LogLevelDebug, fmt.Sprintf("goSourceSeek[id %p]: OK [seek %d | whence %d]", src, n, whence))

	return C.gint64(n)
}
