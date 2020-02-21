package sort

import (
	"sort"
)

type Comparable interface {
	// Return <0, 0, or >0 if the receiver is less than, equal to, or greater than the parameter
	Compare(Comparable) int
}

type Sortable []Comparable

func (s Sortable) Len() int {
	return len(s)
}

func (s Sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sortable) Less(i, j int) bool {
	return s[i].Compare(s[j]) < 0
}

func Sort(s Sortable) {
	sort.Sort(s)
}
