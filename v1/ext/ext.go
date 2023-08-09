package ext

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
