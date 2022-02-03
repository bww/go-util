package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDump(t *testing.T) {
	err1 := fmt.Errorf("Error 1")
	err2 := fmt.Errorf("Error 2: %w", err1)
	err3 := fmt.Errorf("Error 3: %w", err2)

	assert.Equal(t, "", Sdump(nil))
	assert.Equal(t, "Error 1", Sdump(err1))
	assert.Equal(t, "Error 2: Error 1\n\tbecause: Error 1", Sdump(err2))
	assert.Equal(t, "Error 3: Error 2: Error 1\n\tbecause: Error 2: Error 1\n\tbecause: Error 1", Sdump(err3))
}
