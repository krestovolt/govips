package vips

import (
	"bufio"
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

	reader := bufio.NewReader(fin)

	imgRef, err := NewImageFromReader(reader)
	defer imgRef.Close()

	assert.NoError(t, err)
	assert.NotEmpty(t, imgRef.Width())
	assert.NotEmpty(t, imgRef.Height())

	imgRef.Thumbnail(320, 85, InterestingNone)

	buf, mt, err := imgRef.Export(&ExportParams{
		Format:      ImageTypeJPEG, // defaults to the starting encoder
		Quality:     60,
		Compression: 7,
		Interlaced:  true,
		Lossless:    false,
		Effort:      6,
	})

	fmt.Printf(">>>> w=%d h=%d b=%d \n", imgRef.Width(), imgRef.Height(), reader.Buffered())
	fmt.Printf(">>>> %v %v %v\n", imgRef, len(buf), mt.Format.FileExt())

	imgRef.Close()

	err = ioutil.WriteFile(resources+"Snake_River.output.jpg", buf, 0644)
}

func Benchmark_VipsCustomSource__JPEG(b *testing.B) {
	startupIfNeeded()

	fin, _ := os.Open(resources + "fur-cats-siamese-cat-like-mammal-395436-pxhere.com.jpg")
	reader := bufio.NewReader(fin)

	for i := 0; i < b.N; i += 1 {
		imgRef, _ := NewImageFromReader(reader)

		imgRef.Thumbnail(320, 85, InterestingNone)
		buf, mt, err := imgRef.Export(&ExportParams{
			Format:      ImageTypeJPEG, // defaults to the starting encoder
			Quality:     60,
			Compression: 7,
			Interlaced:  true,
			Lossless:    false,
			Effort:      6,
		})
		if len(buf) == 0 || mt == nil || err != nil {
		}

		fin.Seek(0, io.SeekStart)
		reader.Reset(fin)
	}
}
