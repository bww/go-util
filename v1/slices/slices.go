package slices

// Apply a function to every element in a slice, returning
// the parameter slice, whose elements may be mutated.
func Apply[T any](s []T, f func(T) T) []T {
	for i, e := range s {
		s[i] = f(e)
	}
	return s
}

// Map every element in an input slice to a countepart output
// element by applying the specified function.
func Map[X, Y any](s []X, f func(X) Y) []Y {
	r := make([]Y, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}
