package urls

import (
	"fmt"
	"net/url"
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

// Merge the specified qweury parameters into the provided URL. Parameters
// with existing keys are added, not replaced.
func MergeQuery(s string, p url.Values) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	e := u.Query()
	for k, v := range p {
		for _, x := range v {
			e.Add(k, x)
		}
	}
	u.RawQuery = e.Encode()
	return u.String(), nil
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
