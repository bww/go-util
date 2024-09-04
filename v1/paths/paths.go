package paths

import (
	"strings"
)

var esc = '\\'

// First separates a path at the first component, not the last. The separator
// used for this purpose is '/' and the escape character used is '\'. To use
// the OS path separator, use FirstDelim(path, os.PathSeparator) instead.
func First(path string) (string, string) {
	return FirstDelim(path, '/')
}

// First separates a path at the first component, not the last. The separator
// used for this purpose is the one provided and the escape character used is
// '\'.
func FirstDelim(path string, sep rune) (string, string) {
	path = strings.TrimSpace(path)
	if path == "" {
		return "", ""
	}
	sb := &strings.Builder{}
	var (
		x bool
		i int
		e rune
	)
	for i, e = range path {
		if x && (e == sep || e == esc) {
			sb.WriteRune(e)
			x = false
		} else if e == esc {
			x = true
		} else if e == sep {
			break
		} else {
			sb.WriteRune(e)
			x = false
		}
	}
	if l := len(path); i < l {
		return sb.String(), path[i+1:]
	} else {
		return path, ""
	}
}
