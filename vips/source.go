package vips

// #cgo pkg-config: vips
// #include "source.h"
import "C"
import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"sync"
	"unsafe"
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
	reader io.Reader
	seeker io.Seeker
	args   *C.struct__GoSourceArguments
	src    *C.struct__VipsSourceCustom
	// read signal handler id
	rsigHandle C.gulong
	// seek signal handler id
	ssigHandle C.gulong
}

// NewSource creates a new image source that uses a regular io.Reader
func NewSource(image io.Reader) *Source {
	src := &Source{
		reader: image,
	}

	skr, ok := image.(io.ReadSeeker)
	if ok {
		src.seeker = skr
	}

	sourceMu.Lock()
	sourceCtr++
	id := sourceCtr
	sources[id] = src
	sourceMu.Unlock()

	govipsLog("govips", LogLevelDebug, fmt.Sprintf("Created image source %d", id))

	src.args = C.create_go_source_arguments(C.int(id))
	src.src = C.create_go_custom_source(src.args)
	src.rsigHandle = C.connect_go_signal_read(src.src, src.args)
	src.ssigHandle = C.connect_go_signal_seek(src.src, src.args)

	return src
}

func (s *Source) Close() {
	imageID := int(s.args.image_id)
	govipsLog("govips", LogLevelDebug, fmt.Sprintf("Closing source %d", imageID))

	sourceMu.Lock()
	sources[imageID] = nil
	s.free()
	sourceMu.Unlock()
}

func (s *Source) free() {
	C.free_go_custom_source(s.src, s.args, s.rsigHandle, s.ssigHandle)
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
