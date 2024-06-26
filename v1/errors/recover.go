package errors

// Recovery is implemented by errors which can describe whether or not
// they can be recovered from.
type Recovery interface {
	error

	// Recoverable indicates whether the error is recoverable or not. What
	// exactly this means depends on the context in which the error occurs and is
	// expected to be defined externally from this interface.
	Recoverable() bool
}

// Recoverable inspects the provided error to determine if it both implements
// the interface Recovery and returns true from the method Recoverable. If both
// of these conditions are true, this method returns true.
func Recoverable(err error) bool {
	if r, ok := err.(Recovery); ok {
		return r.Recoverable()
	} else {
		return false
	}
}

// recoveryError wraps another error and provides the Recovery interface
type recoveryError struct {
	error
	rec bool
}

// Wrap an error and make it recoverable
func NewRecoverable(err error, rec bool) Recovery {
	return recoveryError{
		error: err,
		rec:   rec,
	}
}

func (e recoveryError) Error() string {
	return e.error.Error()
}

func (e recoveryError) Unwrap() error {
	return e.error
}

func (e recoveryError) Recoverable() bool {
	return e.rec
}
