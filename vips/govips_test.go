package vips

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	running = false
	Startup(&Config{})
	startupIfNeeded()
}
