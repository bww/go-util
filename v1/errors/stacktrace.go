package errors

import (
	"fmt"
	"strings"

	"github.com/bww/go-util/v1/debug"
)

type stacktraceError struct {
	err   error
	stack []debug.Frame
}

// Stacktrace captures the current stack (excluding itself) and
// wraps the provided error with it.
func Stacktrace(err error) error {
	return stacktraceError{
		err:   err,
		stack: debug.Stacktrace()[1:], // trim off one frame (the Stacktrace function itself)
	}
}

func (e stacktraceError) Frames() []debug.Frame {
	return e.stack
}

func (e stacktraceError) Unwrap() error {
	return e.err
}

func (e stacktraceError) Error() string {
	b := &strings.Builder{}
	b.WriteString(e.err.Error())
	b.WriteString(":\n")
	for _, f := range e.stack {
		b.WriteString(fmt.Sprintf("\t%v\n", f))
	}
	return b.String()
}
