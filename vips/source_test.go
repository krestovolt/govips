package vips

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_VipsCustomSource__JPEG(t *testing.T) {
	Startup(&Config{})

	fin, err := os.Open(resources + "Snake_River.jpg")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imgRef, err := NewImageFromReader(fin)
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
	startupIfNeeded()

	fin, _ := os.Open(resources + "fur-cats-siamese-cat-like-mammal-395436-pxhere.com.jpg")
	for i := 0; i < b.N; i += 1 {
		imgRef, _ := NewImageFromReader(fin)

		imgRef.Thumbnail(320, 85, InterestingNone)
		buf, mt, err := imgRef.ExportNative()
		if len(buf) == 0 || mt == nil || err != nil {
		}

		fin.Seek(0, io.SeekStart)
	}
}
