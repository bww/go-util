package ext

// Zeroer is implemented by types that can represent a non-literal zero value
type Zeroer interface {
	IsZero() bool
}

// Choose implements the missing ternary (?:) operator. When the provided
// expression evalutes to true, the first result parameter is returned;
// otherwise the second.
//
//	Choose(len(s) == 1, "true", "false")
func Choose[T any](expr bool, ift T, iff T) T {
	if expr {
		return ift
	} else {
		return iff
	}
}

// Coalesce returns the first non-zero value in the provided arguments. If
// no argument is non-zero, the zero value is returned.
//
//	Coalesce("", "", "hello", "") // "hello"
//	Coalesce(nil, ptr, nil) // ptr
func Coalesce[T comparable](v ...T) T {
	var z T
	for _, e := range v {
		if e != z {
			return e
		}
	}
	return z
}
