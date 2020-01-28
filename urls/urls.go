package urls

import (
	"fmt"
	"path"
)

// Join a base URL together with path components, similar to the
// functionality of path.Join.
func Join(b string, c ...interface{}) string {
	if len(c) < 1 {
		return b
	}
	s := make([]string, len(c))
	for i, e := range c {
		s[i] = fmt.Sprint(e)
	}
	p := path.Join(s...)
	for i, e := range p {
		if e != '/' {
			p = p[i:]
			break
		}
	}
	if l := len(b); l < 1 || b[l-1] != '/' {
		b = b + "/"
	}
	return b + p
}

// Return a file URL for the provided path
func File(p string) string {
	var s string
	for i := 0; i < len(p); i++ {
		if p[i] != '/' {
			s = p[i:]
			break
		}
	}
	return "file:///" + s
}
