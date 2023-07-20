package mem_tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"testing"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/davidbyttow/govips/v2/vips/iox"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Startup() {
	// We need zero MaxCacheSize
	vips.Startup(&vips.Config{
		MaxCacheSize: 0,
	})
}

func TestMain(m *testing.M) {
	Startup()
	ret := m.Run()
	vips.Shutdown()
	os.Exit(ret)
}

func TestMemoryLeak__Buffer(t *testing.T) {
	vips.Startup(&vips.Config{
		ReportLeaks: true,
	})

	buf, err := ioutil.ReadFile(resources + "png-24bit.png")
	require.NoError(t, err)

	iteration := func() {
		ref, err := vips.NewImageFromBuffer(buf)
		require.NoError(t, err)
		defer runtime.KeepAlive(ref)

		_, err = ref.ToBytes()
		assert.NoError(t, err)
	}

	// First iteration for some constant allocations...
	iteration()
	runtime.GC()

	var after, before vips.MemoryStats
	vips.ReadVipsMemStats(&before)

	// More image processing iterations
	for pass := 1; pass < 5; pass++ {
		iteration()
		runtime.GC()
	}

	vips.ReadVipsMemStats(&after)
	delta := after.Mem - before.Mem
	t.Log(fmt.Sprintf("Memory usage: before %d, after %d, delta %d", before.Mem, after.Mem, delta))
	assert.True(t, delta < 10*1024*1024, "Memory usage delta too big: %d", delta)
}

func TestMemoryLeak__File(t *testing.T) {
	vips.Startup(&vips.Config{
		ReportLeaks: true,
	})

	iteration := func() {
		ref, err := vips.NewImageFromFile(resources + "png-24bit.png")
		require.NoError(t, err)
		defer runtime.KeepAlive(ref)

		_, err = ref.ToBytes()
		assert.NoError(t, err)
	}

	// First iteration for some constant allocations...
	iteration()
	runtime.GC()

	var after, before vips.MemoryStats
	vips.ReadVipsMemStats(&before)

	// More image processing iterations
	for pass := 1; pass < 5; pass++ {
		iteration()
		runtime.GC()
	}

	vips.ReadVipsMemStats(&after)
	delta := after.Mem - before.Mem
	t.Log(fmt.Sprintf("Memory usage: before %d, after %d, delta %d", before.Mem, after.Mem, delta))
	assert.True(t, delta < 10*1024*1024, "Memory usage delta too big: %d", delta)
}

func TestMemoryLeak__Source(t *testing.T) {
	vips.Startup(&vips.Config{
		ReportLeaks: true,
	})

	iteration := func() {
		buf, err := iox.NewBufferedFileReader(resources + "png-24bit.png")
		require.NoError(t, err)

		ref, err := vips.NewImageFromReader(buf, true)
		require.NoError(t, err)
		defer runtime.KeepAlive(ref)

		_, err = ref.ToBytes()
		assert.NoError(t, err)
	}

	// First iteration for some constant allocations...
	iteration()
	runtime.GC()

	var after, before vips.MemoryStats
	vips.ReadVipsMemStats(&before)

	// More image processing iterations
	for pass := 1; pass < 5; pass++ {
		iteration()
		runtime.GC()
	}

	vips.ReadVipsMemStats(&after)
	delta := after.Mem - before.Mem
	t.Log(fmt.Sprintf("Memory usage: before %d, after %d, delta %d", before.Mem, after.Mem, delta))
	assert.True(t, delta < 10*1024*1024, "Memory usage delta too big: %d", delta)
}
