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

var (
	// FIXME: need to find a way to handle this properly,
	// maybe the `sources` will grow indefinitely if this is used in long-lived
	// process?
	sourceCtr int
	sources   = make(map[int]*Source)
	sourceMu  = sync.RWMutex{}
)

type Source struct {
	id     int
	reader iox.PeekableReader
	seeker io.Seeker
	closer io.Closer
	args   *C.struct__GoSourceArguments
	src    *C.struct__VipsSourceCustom
	// read signal handler id
	rsigHandle C.gulong
	// seek signal handler id
	ssigHandle C.gulong
}

// NewSource creates a new image source that uses a `iox.PeekableReader` (e.g. bufio.Reader)
//
// By default it will attach a `read` signal and call the `Read` method of the reader,
// if the reader also supports the `io.Seeker` then the `seek` signal handler would also be attached.
//
// The needs of `iox.PeekableReader` interface determined by how the internal load function extract the `vips.ImageType` from the actual source's reader.
// The load function needs to get the first 12 bytes without advancing the reader position, this makes it possible to prevent reading invalid buffer that can cause a whole application to crash (or segvault-ing).
func NewSource(image iox.PeekableReader) *Source {
	src := &Source{
		reader: image,
	}

	skr, ok := image.(io.Seeker)
	if ok {
		src.seeker = skr
	}

	clr, ok := image.(io.Closer)
	if ok {
		src.closer = clr
	}

	sourceMu.Lock()
	id := sourceCtr
	sources[id] = src
	src.id = id
	sourceCtr++
	sourceMu.Unlock()

	govipsLog("govips", LogLevelDebug, fmt.Sprintf("Created image source %d", id))

	src.args = C.create_go_source_arguments(C.int(id))
	src.src = C.create_go_custom_source(src.args)
	src.rsigHandle = C.connect_go_signal_read(src.src, src.args)
	src.ssigHandle = C.connect_go_signal_seek(src.src, src.args)

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
	sourceMu.Lock()
	imageID := int(s.id)
	govipsLog("govips", LogLevelInfo, fmt.Sprintf("Closing source %d", imageID))

	C.free_go_custom_source(s.src, s.args, s.rsigHandle, s.ssigHandle)

	if s.closer != nil {
		s.closer.Close()
	}

	s.closer = nil
	s.reader = nil
	s.seeker = nil
	s.src = nil
	s.args = nil
	sources[imageID] = nil

	sourceMu.Unlock()
}

//export goSourceRead
func goSourceRead(cImageID C.int, buffer unsafe.Pointer, bufSize C.gint64) (read C.gint64) {
	imageID := int(cImageID)
	sourceMu.RLock()
	src, ok := sources[imageID]
	sourceMu.RUnlock()
	if !ok {
		govipsLog("govips", LogLevelError, fmt.Sprintf("goSourceRead[id %d]: Source not found", imageID))
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
		govipsLog("govips", LogLevelDebug, fmt.Sprintf("goSourceRead[id %d]: EOF [read %d]", imageID, n))
		return C.gint64(n)
	} else if err != nil {
		govipsLog("govips", LogLevelError, fmt.Sprintf("goSourceRead[id %d]: Error: %v [read %d]", imageID, err, n))
		return -1
	}

	govipsLog("govips", LogLevelDebug, fmt.Sprintf("goSourceRead[id %d]: OK [read %d]", imageID, n))
	return C.gint64(n)
}

//export goSourceSeek
func goSourceSeek(cImageID C.int, offset C.gint64, cWhence C.int) (newOffset C.gint64) {
	imageID := int(cImageID)
	sourceMu.RLock()
	src, ok := sources[imageID]
	sourceMu.RUnlock()
	if !ok {
		govipsLog("govips", LogLevelError, fmt.Sprintf("goSourceSeek[id %d]: Source not found", imageID))
		return -1 // Not found
	}

	if src.seeker == nil {
		govipsLog("govips", LogLevelDebug, fmt.Sprintf("goSourceRead[id %d]: Seek not supported", imageID))
		return -1 // Unsupported!
	}

	whence := int(cWhence)
	switch whence {
	case io.SeekStart, io.SeekCurrent, io.SeekEnd:
	default:
		govipsLog("govips", LogLevelError, fmt.Sprintf("goSourceSeek[id %d]: Invalid whence value [%d]", imageID, whence))
		return -1
	}

	n, err := src.seeker.Seek(int64(offset), whence)
	if err != nil {
		govipsLog("govips", LogLevelError, fmt.Sprintf("goSourceSeek[id %d]: Error: %v [offset %d | whence %d]", imageID, err, n, whence))
		return -1
	}

	govipsLog("govips", LogLevelDebug, fmt.Sprintf("goSourceSeek[id %d]: OK [seek %d | whence %d]", imageID, n, whence))

	return C.gint64(n)
}
