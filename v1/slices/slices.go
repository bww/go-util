package slices

import (
	"sort"
)

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
		for _, x := range e {
			r = append(r, x)
		}
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
		for _, x := range f(e) {
			r = append(r, x)
		}
	}
	return r
}

// Comparator evalutes two arguments of the same type and returns
// an integer value describing how they compare, returning:
//
//   - 1 if a > b
//   - 0 if the elements are equal
//   - -1 if a < b
type Comparator[T any] func(a, b T) int

type byComparator[T any] struct {
	data []T
	cmp  Comparator[T]
}

func (c byComparator[T]) Len() int           { return len(c.data) }
func (c byComparator[T]) Swap(i, j int)      { c.data[i], c.data[j] = c.data[j], c.data[i] }
func (c byComparator[T]) Less(i, j int) bool { return c.cmp(c.data[i], c.data[j]) < 0 }

// Sort a slice in place by appying the specified comparator
// function to elements.
func Sort[T any](s []T, f Comparator[T]) {
	sort.Sort(byComparator[T]{s, f})
}

// Sort a copy of a slice by appying the specified comparator
// function to elements. The sorted copy is returned.
func Sorted[T any](s []T, f Comparator[T]) []T {
	d := make([]T, len(s))
	copy(d, s)
	sort.Sort(byComparator[T]{d, f})
	return d
}
