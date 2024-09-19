package errors

import (
	"errors"
	"fmt"

	"github.com/bww/go-util/v1/uuid"
)

type Referenced interface {
	error
	Unwrap() error
	Reference() string
}

// Reference wraps an error that has a reference identifier which can
// be used to identify related information, probably in logs.
type referencedError struct {
	err error
	ref string
}

// Reference generates a random reference identifier and wraps the provided
// error in a new referenced error with that identifier.
func Reference(err error) Referenced {
	return referencedError{
		err: err,
		ref: fmt.Sprintf("err-%v", uuid.New()),
	}
}

// Unreference inspects the provided error to determine if it implements the
// interface Referenced. If so, the result of err.Unwrap() is returned along
// with the reference identifier; and if not, the input error itself is returned
// with an empty string as the identifier.
func Unreference(err error) (error, string) {
	var r Referenced
	if errors.As(err, &r) {
		return r.Unwrap(), r.Reference()
	} else {
		return err, ""
	}
}

func (e referencedError) Unwrap() error {
	return e.err
}

func (e referencedError) Error() string {
	return fmt.Sprintf("%v (ref: %s)", e.err, e.ref)
}

func (e referencedError) Reference() string {
	return e.ref
}
