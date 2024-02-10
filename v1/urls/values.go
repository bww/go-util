package urls

import (
	"net/url"
	"strings"
)

type MergeConfig struct {
	Append bool
}

func (c MergeConfig) WithOptions(opts []MergeOption) MergeConfig {
	for _, opt := range opts {
		c = opt(c)
	}
	return c
}

type MergeOption func(MergeConfig) MergeConfig

func Append(b bool) MergeOption {
	return func(c MergeConfig) MergeConfig {
		c.Append = b
		return c
	}
}

// Merge the specified values. Either one of the parameters or a new set of
// values is returned.
func MergeValues(a, b url.Values, opts ...MergeOption) (url.Values, error) {
	return mergeValues(a, b, MergeConfig{}.WithOptions(opts))
}

func mergeValues(a, b url.Values, conf MergeConfig) (url.Values, error) {
	if len(a) == 0 && len(b) == 0 {
		return a, nil
	} else if len(a) == 0 {
		return b, nil
	} else if len(b) == 0 {
		return a, nil
	}
	q := make(url.Values)
	for _, e := range []url.Values{a, b} {
		for k, v := range e {
			if conf.Append {
				q[k] = append(q[k], v...)
			} else {
				q[k] = v
			}
		}
	}
	return q, nil
}

// Parse parameters from the provided Values and produce a list of
// results for the specified key. Both values defined by repeatedly
// providing a key and those provided in a delimited list for a single
// key are returned.
//
// For example, using the delimiter ',':
//
//	a=one&a=two        -> [one, two]
//	a=one,two          -> [one, two]
//	a=one,two&a=three  -> [one, two, three]
func ParseValueList(v url.Values, k, d string) []string {
	var l []string
	for _, e := range v[k] {
		if e != "" {
			for _, c := range strings.Split(e, d) {
				l = append(l, strings.TrimSpace(c))
			}
		}
	}
	return l
}
