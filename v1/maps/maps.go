package maps

// Copy makes a shallow copy of a map
func Copy[K comparable, V any](m map[K]V) map[K]V {
	d := make(map[K]V)
	for k, v := range m {
		d[k] = v
	}
	return d
}

// Merge merges the contents of several maps together into the first provided
// map, which is also returned. Each successive instance of a key found in any
// map overwrites a previously encountered instance. Provide the maps in
// inverse order of priority.
//
// If no maps are provided as input, nil is returned, not an empty map.
func Merge[K comparable, V any](s ...map[K]V) map[K]V {
	l := len(s)
	if l == 0 {
		return nil
	} else if l == 1 {
		return s[0]
	}
	d := s[0] // first map is also the output
	for i := 1; i < l; i++ {
		for k, v := range s[i] {
			d[k] = v
		}
	}
	return d
}

// Merged has the same behavior as Merge(), except that the maps are merged
// into a newly allocated map, not one of the parameters.
func Merged[K comparable, V any](s ...map[K]V) map[K]V {
	if len(s) > 0 {
		return Merge(append([]map[K]V{make(map[K]V)}, s...)...)
	} else {
		return nil
	}
}
