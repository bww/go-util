package urls

import (
	"net/url"
	"strings"
)

// Parse parameters from the provided Values and produce a list of
// results for the specified key. Both values defined by repeatedly
// providing a key and those provided in a delimited list for a single
// key are returned.
//
// For example, using the delimiter ',':
//   a=one&a=two        -> [one, two]
//   a=one,two          -> [one, two]
//   a=one,two&a=three  -> [one, two, three]
//
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
