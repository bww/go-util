package slices

// Apply a function to every element in a slice, returning the parameter
// slice, whose elements may be mutated.
func Apply[T any](s []T, f func(T) T) []T {
	for i, e := range s {
		s[i] = f(e)
	}
	return s
}

// Map converts every element in an input slice to a countepart output
// element by applying the specified function.
func Map[X, Y any](s []X, f func(X) Y) []Y {
	r := make([]Y, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

// MapAny converts every element in an input slice to a countepart output
// element containing the same input but cast to any/interface{}. This is a
// special case of Map which provides a convenience for a common use case.
func MapAny[X any](s []X) []any {
	r := make([]any, len(s))
	for i, e := range s {
		r[i] = e
	}
	return r
}

// Flatten creates a new output slice that contains the elements from the
// specified input slices.
func Flatten[X any](s [][]X) []X {
	r := make([]X, 0, len(s))
	for _, e := range s {
		r = append(r, e...)
	}
	return r
}

// FlatMap converts every element in an input slice to a countepart output
// slice by applying the specified function, then flatten the output slices
// into a single array of elements.
//
// This is conceptually the equivalent of Flatten(Map(...))
func FlatMap[X, Y any](s []X, f func(X) []Y) []Y {
	r := make([]Y, 0, len(s))
	for _, e := range s {
		r = append(r, f(e)...)
	}
	return r
}

// FindValue matches a comparable value
func MatchValue[T comparable](v T) func(T) bool {
	return func(e T) bool {
		return e == v
	}
}

// Find searches for an element in a slice, sequentially. This should be used
// on small slices where performance is not a consideration. For large slices,
// consider the standard library function slices.BinarySearch instead.
func FindFunc[T any](s []T, match func(T) bool) (T, bool) {
	var zero T
	for _, e := range s {
		if match(e) {
			return e, true
		}
	}
	return zero, false
}
