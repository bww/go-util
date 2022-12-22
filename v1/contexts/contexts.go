package contexts

import (
	"context"
)

// Continue is a convenience wrapper around a common context
// cancelation check. It is intended to be used when processing
// a loop that may be canceled; like so:
//
//	for contexts.Continue(cxt) {
//	  // continue processing something...
//	}
func Continue(cxt context.Context) bool {
	select {
	case <-cxt.Done():
		return false
	default:
		return true
	}
}

// Continue is a convenience wrapper around a common context
// cancelation check. It is essentially the inverse of Continue:
//
//	for {
//	  if contexts.Done(cxt) {
//	    break
//	  }
//	  // otherwise...
//	}
func Done(cxt context.Context) bool {
	select {
	case <-cxt.Done():
		return true
	default:
		return false
	}
}
