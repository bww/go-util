package paths

import (
	"os"
	"strings"
)

type Config struct {
	Sep, Esc rune
}

// OSConfig defines a path configuration using the OS path separator
var OSConfig = Config{
	Sep: os.PathSeparator,
	Esc: '\\',
}

// First separates a path at the first component, not the last. The separator
// used for this purpose is '/' and the escape character used is '\'. To use
// the OS path separator, use FirstDelim(path, OSConfig) instead.
func First(path string) (string, string) {
	return FirstConfig(path, Config{Sep: '/', Esc: '\\'})
}

// First separates a path at the first component, not the last. The separator
// used for this purpose is the one provided and the escape character used is
// '\'.
func FirstConfig(path string, conf Config) (string, string) {
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
		if x && (e == conf.Sep || e == conf.Esc) {
			sb.WriteRune(e)
			x = false
		} else if e == conf.Esc {
			x = true
		} else if e == conf.Sep {
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
