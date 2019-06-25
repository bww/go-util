package urls

import (
	"fmt"
	"path"
)

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
