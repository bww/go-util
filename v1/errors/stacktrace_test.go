package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const errorMessagePrefix = "This is the error:\n\tv1/errors/stacktrace_test.go:14 github.com/bww/go-util/v1/errors.TestStacktraceError\n"

func TestStacktraceError(t *testing.T) {
	// The test below depends on the line the test is on in this file:
	assert.Equal(t, errorMessagePrefix, Stacktrace(errors.New("This is the error")).Error()[:len(errorMessagePrefix)])
}
