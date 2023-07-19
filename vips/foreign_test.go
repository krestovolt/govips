package vips

import (
	"io/ioutil"
	"os"
	"testing"

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

	fin, err := os.Open(resources + "jpg-24bit-icc-iec.jpg")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imgRef, err := NewImageFromReader(fin)
	assert.NoError(t, err)
	assert.NotNil(t, imgRef)
	defer imgRef.Close()

	imageType := determinePartialImageType(imgRef.image)
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

	fin, err := os.Open(resources + "heic-24bit-exif.heic")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imgRef, err := NewImageFromReader(fin)
	assert.NoError(t, err)
	assert.NotNil(t, imgRef)
	defer imgRef.Close()

	imageType := determinePartialImageType(imgRef.image)
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

	fin, err := os.Open(resources + "heic-24bit.heic")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imgRef, err := NewImageFromReader(fin)
	assert.NoError(t, err)
	assert.NotNil(t, imgRef)
	defer imgRef.Close()

	imageType := determinePartialImageType(imgRef.image)
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

	fin, err := os.Open(resources + "png-24bit+alpha.png")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imgRef, err := NewImageFromReader(fin)
	assert.NoError(t, err)
	assert.NotNil(t, imgRef)
	defer imgRef.Close()

	imageType := determinePartialImageType(imgRef.image)
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

	fin, err := os.Open(resources + "tif.tif")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imgRef, err := NewImageFromReader(fin)
	assert.NoError(t, err)
	assert.NotNil(t, imgRef)
	defer imgRef.Close()

	imageType := determinePartialImageType(imgRef.image)
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

	fin, err := os.Open(resources + "webp+alpha.webp")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imgRef, err := NewImageFromReader(fin)
	assert.NoError(t, err)
	assert.NotNil(t, imgRef)
	defer imgRef.Close()

	imageType := determinePartialImageType(imgRef.image)
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

	fin, err := os.Open(resources + "svg.svg")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imgRef, err := NewImageFromReader(fin)
	assert.NoError(t, err)
	assert.NotNil(t, imgRef)
	defer imgRef.Close()

	imageType := determinePartialImageType(imgRef.image)
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

	fin, err := os.Open(resources + "svg_1.svg")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imgRef, err := NewImageFromReader(fin)
	assert.NoError(t, err)
	assert.NotNil(t, imgRef)
	defer imgRef.Close()

	imageType := determinePartialImageType(imgRef.image)
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

	fin, err := os.Open(resources + "pdf.pdf")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imgRef, err := NewImageFromReader(fin)
	assert.NoError(t, err)
	assert.NotNil(t, imgRef)
	defer imgRef.Close()

	imageType := determinePartialImageType(imgRef.image)
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

// TODO check compatibility for this type
// func Test_DeterminePartialImageType__BMP(t *testing.T) {
// 	Startup(&Config{})

// 	fin, err := os.Open(resources + "bmp.bmp")
// 	assert.NoError(t, err)
// 	assert.NotNil(t, fin)

// 	imgRef, err := NewImageFromReader(fin)
// 	assert.Error(t, err)
// 	assert.Nil(t, imgRef)
// 	defer imgRef.Close()

// 	imageType := determinePartialImageType(imgRef.image)
// 	assert.Equal(t, ImageTypePDF, imageType)
// }

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

	fin, err := os.Open(resources + "avif-8bit.avif")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imgRef, err := NewImageFromReader(fin)
	assert.NoError(t, err)
	assert.NotNil(t, imgRef)
	defer imgRef.Close()

	imageType := determinePartialImageType(imgRef.image)
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

	fin, err := os.Open(resources + "jp2k-orientation-6.jp2")
	assert.NoError(t, err)
	assert.NotNil(t, fin)

	imgRef, err := NewImageFromReader(fin)
	assert.NoError(t, err)
	assert.NotNil(t, imgRef)
	defer imgRef.Close()

	imageType := determinePartialImageType(imgRef.image)
	assert.Equal(t, ImageTypeJP2K, imageType)
}
