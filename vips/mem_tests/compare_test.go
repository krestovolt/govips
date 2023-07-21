package mem_tests

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/davidbyttow/govips/v2/vips"
)

func init() {
	vips.LoggingSettings(func(messageDomain string, messageLevel vips.LogLevel, message string) {
		//nil
	}, vips.LogLevelWarning)
}

func benchHelper(b *testing.B, loader string, file string, f func(file string) (*vips.ImageRef, error), exec func(img *vips.ImageRef)) {
	vips.Startup(nil)

	for i := 0; i < b.N; i++ {
		// Test with loading the buffer immediately into memory
		img, _ := f(file)
		exec(img)
		buf, _, _ := img.Export(nil)
		//no use
		ioutil.Discard.Write(buf)
	}
	var ms vips.MemoryStats
	vips.ReadVipsMemStats(&ms)
	b.Logf("[%s] Mem: %d, MemHigh: %d, Files: %d, Allocs: %d\n", loader, ms.Mem, ms.MemHigh, ms.Files, ms.Allocs)
}

var testFile = resources + "jpg-24bit.jpg"
var testFunc = func(img *vips.ImageRef) {
	img.Thumbnail(100, 100, vips.InterestingCentre)
}

func BenchmarkLoadFromFile(b *testing.B) {
	benchHelper(b, "vips.NewImageFromFile", testFile, vips.NewImageFromFile, testFunc)
}

func BenchmarkLoadFromSource__Sequential(b *testing.B) {
	benchHelper(b, "vips.NewImageSourceFromReader", testFile, func(file string) (*vips.ImageRef, error) {
		// Test with using an image source
		imageFile, _ := os.Open(testFile)

		importParams := vips.NewImportParams()
		importParams.AccessMode.Set(vips.VipsAccessSequential.Int())
		return vips.LoadImageFromReader(imageFile, importParams)
	}, testFunc)
}

func BenchmarkLoadFromSource__Random(b *testing.B) {
	benchHelper(b, "vips.NewImageSourceFromReader", testFile, func(file string) (*vips.ImageRef, error) {
		// Test with using an image source
		imageFile, _ := os.Open(testFile)
		return vips.LoadImageFromReader(imageFile, nil)
	}, testFunc)
}
