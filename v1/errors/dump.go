package errors

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func Fdump(w io.Writer, err error) (int, error) {
	var n int
	for e := err; e != nil; e = errors.Unwrap(e) {
		if n > 0 {
			if x, err := fmt.Fprint(w, "\n\tbecause: "); err != nil {
				return n, err
			} else {
				n += x
			}
		}
		if x, err := fmt.Fprint(w, e.Error()); err != nil {
			return n, err
		} else {
			n += x
		}
	}
	return n, nil
}

func Dump(err error) (int, error) {
	return Fdump(os.Stdout, err)
}

func Sdump(err error) string {
	b := &strings.Builder{}
	_, reserr := Fdump(b, err)
	if reserr != nil {
		panic(reserr)
	}
	return b.String()
}
