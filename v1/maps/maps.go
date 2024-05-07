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
// map. Each successive instance of a key found in any map overwrites a
// previously encountered instance. Provide the maps in inverse order of
// priority.
func Merge[K comparable, V any](d map[K]V, s ...map[K]V) {
	l := len(s)
	if l < 1 {
		return // if less than two maps are provided, the input is already merged
	}
	for _, e := range s {
		for k, v := range e {
			d[k] = v
		}
	}
}

// Merged has the same behavior as Merge(), except that the maps are merged
// into a newly allocated map, not one of the parameters.
//
// If no maps are provided as input, nil is returned, not an empty map.
func Merged[K comparable, V any](s ...map[K]V) map[K]V {
	if l := len(s); l == 0 {
		return nil
	} else if l == 1 {
		return Copy(s[0])
	}
	d := make(map[K]V)
	Merge(d, s...)
	return d
}
