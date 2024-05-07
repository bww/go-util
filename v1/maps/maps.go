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
func Merge[K comparable, V any](s ...map[K]V) {
	l := len(s)
	if l < 2 {
		return // if less than two maps are provided, the input is already merged
	}
	d := s[0] // first map is also the output
	for i := 1; i < l; i++ {
		for k, v := range s[i] {
			d[k] = v
		}
	}
}

// Merged has the same behavior as Merge(), except that the maps are merged
// into a newly allocated map, not one of the parameters.
//
// If no maps are provided as input, nil is returned, not an empty map.
func Merged[K comparable, V any](s ...map[K]V) map[K]V {
	if len(s) == 0 {
		return nil
	}
	d := make(map[K]V)
	Merge(append([]map[K]V{d}, s...)...)
	return d
}
