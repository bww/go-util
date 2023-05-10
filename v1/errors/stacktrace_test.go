package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const errorMessagePrefix = "This is the error:\n    v1/errors/stacktrace_test.go:14\n        github.com/bww/go-util/v1/errors.TestStacktraceError\n"

func TestStacktraceError(t *testing.T) {
	// The test below depends on the line the test is on in this file; don't change it:
	assert.Equal(t, errorMessagePrefix, Stacktrace(errors.New("This is the error")).Error()[:len(errorMessagePrefix)])
}
