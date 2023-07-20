package vips

// #include <vips/vips.h>
import "C"

import (
	"errors"
	"fmt"
	dbg "runtime/debug"
	"unsafe"
)

type LoadImageError struct {
	Format ImageType
	Err    error
}

func (e *LoadImageError) Error() string {
	msg := "Unknown error"
	if e.Err != nil {
		msg = e.Err.Error()
	}
	return fmt.Sprintf("error loading image with type '%s' from source: %s", e.Format.FileExt(), msg)
}

var (
	// ErrUnsupportedImageFormat when image type is unsupported
	ErrUnsupportedImageFormat = errors.New("unsupported image format")
)

func handleImageError(out *C.VipsImage) error {
	if out != nil {
		clearImage(out)
	}

	return handleVipsError()
}

func handleSaveBufferError(out unsafe.Pointer) error {
	if out != nil {
		gFreePointer(out)
	}

	return handleVipsError()
}

func handleVipsError() error {
	s := C.GoString(C.vips_error_buffer())
	C.vips_error_clear()

	return fmt.Errorf("%v\nStack:\n%s", s, dbg.Stack())
}

func newLoadImageError(err error, format ImageType) *LoadImageError {
	return &LoadImageError{
		Format: format,
		Err:    err,
	}
}
