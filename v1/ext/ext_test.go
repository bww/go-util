package ext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChoose(t *testing.T) {
	assert.Equal(t, "Hello", Choose(true, "Hello", "Goodbye"))
	assert.Equal(t, "Goodbye", Choose(false, "Hello", "Goodbye"))

	assert.Equal(t, 1, Choose(true, 1, 2))
	assert.Equal(t, 2, Choose(false, 1, 2))
}
