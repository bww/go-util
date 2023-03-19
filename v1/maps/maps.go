package maps

// Copy makes a shallow copy of a map
func Copy[K comparable, V any](m map[K]V) map[K]V {
	d := make(map[K]V)
	for k, v := range m {
		d[k] = v
	}
	return d
}
