package errors

import (
	"fmt"
	"strings"

	"github.com/bww/go-util/v1/debug"
	"github.com/bww/go-util/v1/text"
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
	for _, f := range e.stack {
		b.WriteString(fmt.Sprintf("%v\n", f))
	}
	return fmt.Sprintf("%s:\n%s", e.err.Error(), text.Indent(b.String(), "    "))
}
