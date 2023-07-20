package mem_tests

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/davidbyttow/govips/v2/vips/iox"
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
	fmt.Printf("[%s] Mem: %d, MemHigh: %d, Files: %d, Allocs: %d\n", loader, ms.Mem, ms.MemHigh, ms.Files, ms.Allocs)
}

var testFile = resources + "jpg-24bit.jpg"
var testFunc = func(img *vips.ImageRef) {
	img.Thumbnail(100, 100, vips.InterestingCentre)
}

func BenchmarkLoadFromFile(b *testing.B) {
	benchHelper(b, "vips.NewImageFromFile", testFile, vips.NewImageFromFile, testFunc)
}
func BenchmarkLoadFromSource(b *testing.B) {
	benchHelper(b, "vips.NewImageSourceFromReader", testFile, func(file string) (*vips.ImageRef, error) {
		// Test with using an image source
		imageFile, _ := iox.NewBufferedFileReader(file)
		return vips.NewImageSourceFromReader(imageFile, true)
	}, testFunc)
}
