package vips

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/davidbyttow/govips/v2/vips/iox"
	"github.com/stretchr/testify/assert"
)

func Test_VipsCustomSource__JPEG(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "Snake_River.jpg")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imgRef, err := NewImageSourceFromReader(fin, true)
	defer imgRef.Close()

	assert.NoError(t, err)
	assert.NotEmpty(t, imgRef.Width())
	assert.NotEmpty(t, imgRef.Height())

	imgRef.Thumbnail(320, 85, InterestingNone)

	buf, mt, err := imgRef.ExportNative()

	fmt.Printf(">>>> w=%d h=%d\n", imgRef.Width(), imgRef.Height())
	fmt.Printf(">>>> %v %v %v\n", imgRef, len(buf), mt.Format.FileExt())

	imgRef.Close()

	err = ioutil.WriteFile(resources+"Snake_River.output.jpg", buf, 0644)
}

func Benchmark_VipsCustomSource__JPEG(b *testing.B) {
	LoggingSettings(nil, LogLevelError)
	startupIfNeeded()

	fin, _ := iox.NewBufferedFileReader(resources + "fur-cats-siamese-cat-like-mammal-395436-pxhere.com.jpg")
	for i := 0; i < b.N; i += 1 {
		imgRef, _ := NewImageSourceFromReader(fin, true)

		imgRef.Thumbnail(320, 85, InterestingNone)
		buf, mt, err := imgRef.ExportNative()
		if len(buf) == 0 || mt == nil || err != nil {
		}

		if !fin.Rewind() {
			b.Error("Unable to rewind the file input buffer")
		}
	}
}
