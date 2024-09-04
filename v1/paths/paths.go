package paths

import (
	"os"
	"strings"
)

var esc = '\\'

// First separates a path at the first component, not the last. The separator
// used for this purpose is os.PathSeparator and the escape character used is
// '\'.
func First(path string) (string, string) {
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
		if x && (e == os.PathSeparator || e == esc) {
			sb.WriteRune(e)
			x = false
		} else if e == esc {
			x = true
		} else if e == os.PathSeparator {
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
