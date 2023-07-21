package vips

import (
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {
	if _, ok := os.LookupEnv("SKIP_TEST_INIT"); ok {
		t.SkipNow()
	}

	running = false
	Startup(&Config{CollectStats: true, CacheTrace: true})
	running = false
	startupIfNeeded()
}
