package errors

import (
	"errors"
	"fmt"
)

// Redacted is implemented by errors which can contain sensitive information
// of some sort and can represent both an internal and external form. The normal
// error interface returns the redacted form of the error. A caller that knows
// what it's doing can unwrap the unerlying error to report only under specific
// conditions.
type Redacted interface {
	error

	// Unredact produces a representation of this error which may contain
	// sensitive information or is otherwise unsuitable for external consumption.
	// This is the original, unredacted error. The default representation of the
	// error must not contain any sensitive information.
	Unredact() error
}

// Unredact inspects the provided error to determine if it implements the
// interface Redacted. If so, the result of err.Unredact() is returned; and if
// not, the input error itself is returned.
func Unredact(err error) error {
	var r Redacted
	if errors.As(err, &r) {
		return r.Unredact()
	} else {
		return err
	}
}

// redactedError wraps another error and provides internal details
type redactedError struct {
	public   error // non-sensitive, public error
	internal error // sensitive, underlying error
}

// Format a new redacted error
func Redactf(internal error, f string, a ...interface{}) Redacted {
	return Redact(internal, fmt.Errorf(f, a...))
}

// Redact an error by providing a representation that is safe to expose
// externally
func Redact(internal error, public error) Redacted {
	return redactedError{
		public:   public,
		internal: internal,
	}
}

func (e redactedError) Error() string {
	if e.internal != nil {
		return fmt.Sprintf("%s (additional details redacted)", e.public.Error())
	} else {
		return e.public.Error()
	}
}

func (e redactedError) Unwrap() error {
	return e.internal
}

func (e redactedError) Unredact() error {
	return e.internal
}

func (e redactedError) Reference() string {
	var referr Referenced
	if errors.As(e.public, &referr) {
		return referr.Reference()
	} else if errors.As(e.internal, &referr) {
		return referr.Reference()
	} else {
		return ""
	}
}
