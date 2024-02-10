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

// Merge the specified queury parameters into the provided URL. Parameters
// with existing keys are added, not replaced.
func MergeQuery(s string, p ...url.Values) (string, error) {
	if len(p) < 1 {
		return s, nil
	}
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for _, e := range p {
		for k, v := range e {
			for _, x := range v {
				q.Add(k, x)
			}
		}
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// Merge the specified queury parameters into the provided URL. This version
// accepts a set of options which control how the merge is performed. By
// default, parameters with existing keys are replaced, not added.
func MergeQueryOpts(s string, p url.Values, opts ...MergeOption) (string, error) {
	if len(p) < 1 {
		return s, nil
	}
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	u.RawQuery = MergeValues(u.Query(), p, opts...).Encode()
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
