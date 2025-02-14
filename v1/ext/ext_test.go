package ext

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestZeroer(t *testing.T) {
	var when time.Time
	var iface interface{} = when
	z, ok := iface.(Zeroer)
	assert.Equal(t, true, ok)
	assert.Equal(t, true, z.IsZero())
}

func TestChoose(t *testing.T) {
	assert.Equal(t, "Hello", Choose(true, "Hello", "Goodbye"))
	assert.Equal(t, "Goodbye", Choose(false, "Hello", "Goodbye"))

	assert.Equal(t, 1, Choose(true, 1, 2))
	assert.Equal(t, 2, Choose(false, 1, 2))
}

func TestCoalesce(t *testing.T) {
	i := 100
	assert.Equal(t, "hello", Coalesce("", "hello", ""))
	assert.Equal(t, &i, Coalesce(nil, &i, nil))
}

func TestNonzero(t *testing.T) {
	var (
		zero    time.Time
		nonzero = time.Now()
	)
	assert.Equal(t, nonzero, Nonzero(zero, nonzero))
	assert.Equal(t, nonzero, Nonzero(nonzero, zero))
	assert.Equal(t, nonzero, Nonzero(zero, zero, nonzero, zero))
}
