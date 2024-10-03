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
// error in a new referenced error with that identifier. If the parameter
// error is already a Referenced, it is simply returned unmodified.
func Reference(err error) Referenced {
	var referr Referenced
	if errors.As(err, &referr) {
		return referr
	} else {
		return referencedError{
			err: err,
			ref: fmt.Sprintf("err-%v", uuid.New()),
		}
	}
}

// Refstr inspects the provided error to determine if any error in its chain
// implements the interface Referenced. If so, the reference string from the first
// Referenced error encountered is returned.
//
// If you just want the first error in the chain the implements Refererenced,
// use [errors.As] instead.
func Refstr(err error) string {
	var r Referenced
	if errors.As(err, &r) {
		return r.Reference()
	} else {
		return ""
	}
}

func (e referencedError) Unwrap() error {
	return e.err
}

func (e referencedError) Error() string {
	return e.err.Error() // Error() does not include the reference
}

func (e referencedError) Reference() string {
	return e.ref
}

func (e referencedError) String() string {
	return fmt.Sprintf("%v (ref: %s)", e.err, e.ref) // String() does bake in the reference
}
