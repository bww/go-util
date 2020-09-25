package scan

import (
	"strings"
)

// Split a string using a delimiter, which may be esacaped by an escape
// character. When escaped, it is treated as a literal instead of a delimiter.
// The input string up to the the first delimiter and everything after the
// first delimiter is returned. If there are no delimiters in the string the
// entire input string is the first return value. If the input string is empty
// or only contains a single delimiter, both return values will be empty strings.
func Split(s string, d, e rune) (string, string, error) {
	b := &strings.Builder{}

	var esc bool
	for i, r := range s {
		if esc {
			switch r {
			case d:
				b.WriteRune(d)
			case e:
				b.WriteRune(e)
			default:
				return "", "", ErrInvalidEscape
			}
			esc = false
		} else {
			if r == d {
				return b.String(), s[i+1:], nil
			} else if r == e {
				esc = true
			} else {
				b.WriteRune(r)
			}
		}
	}

	return b.String(), "", nil
}

// Join elements with the provided delimiter, escaping that delimiter when
// it occurs in an input component.
func Join(a []string, d, e rune) string {
	b := &strings.Builder{}

	for i, s := range a {
		if i > 0 {
			b.WriteRune(d)
		}
		for _, r := range s {
			if r == d || r == e { // if this is either the delimiter or the escape char, escape it
				b.WriteRune(e)
			}
			b.WriteRune(r)
		}
	}

	return b.String()
}
