package errors

import (
	"fmt"
)

type Error struct {
	Message string      `json:"message"`
	Detail  interface{} `json:"detail,omitempty"`
	Cause   error       `json:"-"`
}

func Errorf(f string, a ...interface{}) *Error {
	return &Error{fmt.Sprintf(f, a...), nil, nil}
}

func (e *Error) SetDetail(d interface{}) *Error {
	e.Detail = d
	return e
}

func (e *Error) SetCause(c error) *Error {
	e.Cause = c
	return e
}

func (e Error) Unwrap() error {
	return e.Cause
}

func (e Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%v: %v", e.Message, e.Cause.Error())
	} else {
		return e.Message
	}
}

type Set []error

// Create a set of errors. Only non-nil parameters are included. If only
// one non-nil parameter is provided it is simply returned and a set is
// not actually created.
func NewSet(e ...error) error {
	s := make(Set, 0)
	for _, v := range e {
		if v != nil {
			s = append(s, v)
		}
	}
	if len(s) == 1 {
		return s[0]
	} else {
		return s
	}
}

// Conform to error. This method simply concatenates the result of Error()
// for all the elements of the set and returns the result.
func (e Set) Error() string {
	var s string
	for i, v := range e {
		if i > 0 {
			s += "; "
		}
		s += v.Error()
	}
	return s
}
