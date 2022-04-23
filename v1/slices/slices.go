package slices

// Apply a function to every element in a slice, returning
// the parameter slice, whose elements may be mutated.
func Apply[T any](s []T, f func(T) T) []T {
	for i, e := range s {
		s[i] = f(e)
	}
	return s
}

// Convert every element in a slice to another value by applying
// a function to it and returning a new slice which contains the
// converted elements.
//
// If any element conversion fails, the routine stops and the
// error returned from converting that element is returned from
// this function.
//
// The primary difference between Convert and Apply is that Apply
// operates on the parameter array in place, and therefore the
// input and output types must be the same.
func Convert[A, B any](s []A, f func(A) (B, error)) ([]B, error) {
	d := make([]B, len(s))
	var err error
	for i, e := range s {
		d[i], err = f(e)
		if err != nil {
			return nil, err
		}
	}
	return d, nil
}
