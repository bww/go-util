package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecoverable(t *testing.T) {
	err := errors.New("Something broke")
	assert.Equal(t, false, Recoverable(err))

	rec := NewRecoverable(err, false)
	assert.Equal(t, false, Recoverable(rec))

	rec = NewRecoverable(err, true)
	assert.Equal(t, true, Recoverable(rec))
	assert.ErrorIs(t, rec, err)
}
