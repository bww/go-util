package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedact(t *testing.T) {
	err := errors.New("Something broke")
	red := Redactf(err, "This is safe to reveal: %d", 7)
	assert.ErrorIs(t, red, err)
	assert.Equal(t, "This is safe to reveal: 7", red.Error())
	assert.Equal(t, errors.New("Something broke"), red.Unredact())
	assert.Equal(t, errors.New("Something broke"), Unredact(red))
	assert.Equal(t, errors.New("Something broke"), Unredact(err))
	assert.Equal(t, errors.New("Something broke"), Unredact(err))
}
