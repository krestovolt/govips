package vips

import (
	"io/ioutil"
	"testing"

	"github.com/davidbyttow/govips/v2/vips/iox"
	"github.com/stretchr/testify/assert"
)

func Test_DetermineImageType__JPEG(t *testing.T) {
	Startup(&Config{})

	buf, err := ioutil.ReadFile(resources + "jpg-24bit-icc-iec.jpg")
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	imageType := DetermineImageType(buf)
	assert.Equal(t, ImageTypeJPEG, imageType)
}

func Test_DeterminePartialImageType__JPEG(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "jpg-24bit-icc-iec.jpg")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imageType, err := DetermineImageReaderType(fin)
	assert.Equal(t, ImageTypeJPEG, imageType)
}

func Test_DetermineImageType__HEIF_HEIC(t *testing.T) {
	Startup(&Config{})

	buf, err := ioutil.ReadFile(resources + "heic-24bit-exif.heic")
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	imageType := DetermineImageType(buf)
	assert.Equal(t, ImageTypeHEIF, imageType)
}

func Test_DeterminePartialImageType__HEIF_HEIC(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "heic-24bit-exif.heic")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imageType, err := DetermineImageReaderType(fin)
	assert.NoError(t, err)
	assert.Equal(t, ImageTypeHEIF, imageType)
}

func Test_DetermineImageType__HEIF_MIF1(t *testing.T) {
	Startup(&Config{})

	buf, err := ioutil.ReadFile(resources + "heic-24bit.heic")
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	imageType := DetermineImageType(buf)
	assert.Equal(t, ImageTypeHEIF, imageType)
}

func Test_DeterminePartialImageType__HEIF_MIF1(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "heic-24bit.heic")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imageType, err := DetermineImageReaderType(fin)
	assert.Equal(t, ImageTypeHEIF, imageType)
}

func Test_DetermineImageType__PNG(t *testing.T) {
	Startup(&Config{})

	buf, err := ioutil.ReadFile(resources + "png-24bit+alpha.png")
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	imageType := DetermineImageType(buf)
	assert.Equal(t, ImageTypePNG, imageType)
}

func Test_DeterminePartialImageType__PNG(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "png-24bit+alpha.png")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imageType, err := DetermineImageReaderType(fin)
	assert.Equal(t, ImageTypePNG, imageType)
}

func Test_DetermineImageType__TIFF(t *testing.T) {
	Startup(&Config{})

	buf, err := ioutil.ReadFile(resources + "tif.tif")
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	imageType := DetermineImageType(buf)
	assert.Equal(t, ImageTypeTIFF, imageType)
}

func Test_DeterminePartialImageType__TIFF(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "tif.tif")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imageType, err := DetermineImageReaderType(fin)
	assert.Equal(t, ImageTypeTIFF, imageType)
}

func Test_DetermineImageType__WEBP(t *testing.T) {
	Startup(&Config{})

	buf, err := ioutil.ReadFile(resources + "webp+alpha.webp")
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	imageType := DetermineImageType(buf)
	assert.Equal(t, ImageTypeWEBP, imageType)
}

func Test_DeterminePartialImageType__WEBP(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "webp+alpha.webp")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imageType, err := DetermineImageReaderType(fin)
	assert.Equal(t, ImageTypeWEBP, imageType)
}

func Test_DetermineImageType__SVG(t *testing.T) {
	Startup(&Config{})

	buf, err := ioutil.ReadFile(resources + "svg.svg")
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	imageType := DetermineImageType(buf)
	assert.Equal(t, ImageTypeSVG, imageType)
}

func Test_DeterminePartialImageType__SVG(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "svg.svg")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	src := NewSource(fin)
	assert.NotNil(t, src)

	img, imageType, err := vipsLoadSource(src, false)
	if img != nil {
		clearImage(img)
		img = nil
	}

	assert.NoError(t, err)
	assert.Equal(t, ImageTypeSVG, imageType)
}

func Test_DetermineImageType__SVG_1(t *testing.T) {
	Startup(&Config{})

	buf, err := ioutil.ReadFile(resources + "svg_1.svg")
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	imageType := DetermineImageType(buf)
	assert.Equal(t, ImageTypeSVG, imageType)
}

func Test_DeterminePartialImageType__SVG_1(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "svg_1.svg")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	src := NewSource(fin)
	assert.NotNil(t, src)

	img, imageType, err := vipsLoadSource(src, false)
	if img != nil {
		clearImage(img)
		img = nil
	}

	assert.NoError(t, err)
	assert.Equal(t, ImageTypeSVG, imageType)
}

func Test_DeterminePartialImageType__SVG_2(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "svg_2.svg")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	src := NewSource(fin)
	assert.NotNil(t, src)

	img, imageType, err := vipsLoadSource(src, false)
	if img != nil {
		clearImage(img)
		img = nil
	}

	assert.NoError(t, err)
	assert.Equal(t, ImageTypeSVG, imageType)
}

func Test_DetermineImageType__PDF(t *testing.T) {
	Startup(&Config{})

	buf, err := ioutil.ReadFile(resources + "pdf.pdf")
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	imageType := DetermineImageType(buf)
	assert.Equal(t, ImageTypePDF, imageType)
}

func Test_DeterminePartialImageType__PDF(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "pdf.pdf")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imageType, err := DetermineImageReaderType(fin)
	assert.Equal(t, ImageTypePDF, imageType)
}

func Test_DetermineImageType__BMP(t *testing.T) {
	Startup(&Config{})

	buf, err := ioutil.ReadFile(resources + "bmp.bmp")
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	imageType := DetermineImageType(buf)
	assert.Equal(t, ImageTypeBMP, imageType)
}

func Test_DeterminePartialImageType__BMP(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "bmp.bmp")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imageType, err := DetermineImageReaderType(fin)
	assert.Equal(t, ImageTypeBMP, imageType)
}

func Test_DetermineImageType__AVIF(t *testing.T) {
	Startup(&Config{})

	buf, err := ioutil.ReadFile(resources + "avif-8bit.avif")
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	imageType := DetermineImageType(buf)
	assert.Equal(t, ImageTypeAVIF, imageType)
}

func Test_DeterminePartialImageType__AVIF(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "avif-8bit.avif")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imageType, err := DetermineImageReaderType(fin)
	assert.Equal(t, ImageTypeAVIF, imageType)
}

func Test_DetermineImageType__JP2K(t *testing.T) {
	Startup(&Config{})

	buf, err := ioutil.ReadFile(resources + "jp2k-orientation-6.jp2")
	assert.NoError(t, err)
	assert.NotNil(t, buf)

	imageType := DetermineImageType(buf)
	assert.Equal(t, ImageTypeJP2K, imageType)
}

func Test_DeterminePartialImageType__JP2K(t *testing.T) {
	Startup(&Config{})

	fin, err := iox.NewBufferedFileReader(resources + "jp2k-orientation-6.jp2")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imageType, err := DetermineImageReaderType(fin)
	assert.Equal(t, ImageTypeJP2K, imageType)
}
