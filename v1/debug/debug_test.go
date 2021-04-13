package debug

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeSourcePath(t *testing.T) {
	// The test below depends on the line the test is on in this file:
	assert.Equal(t, "v1/debug/debug_test.go:11 github.com/bww/go-util/v1/debug.TestRelativeSourcePath", CurrentContext())
}
