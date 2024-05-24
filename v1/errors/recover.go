package errors

type Recovery interface {
	error

	// Recoverable indicates whether the error is recoverable or not.  What
	// exactly this means depends on the context the error occurs in and it is
	// expeced to be defined externally from this interface.
	Recoverable() bool
}

// Recoverable inspects the provided error to determine if it both implements
// the interface Recovery and returns true from the method Recoverable.
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

func (e recoveryError) Recoverable() bool {
	return e.rec
}
