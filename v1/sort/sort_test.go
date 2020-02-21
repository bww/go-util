package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type cmp int

func (c cmp) Compare(a Comparable) int {
	z := a.(cmp)
	return int(c - z)
}

func TestSort(t *testing.T) {
	tests := []struct {
		Unsorted []Comparable
		Expect   []Comparable
	}{
		{
			[]Comparable{cmp(3), cmp(1), cmp(7), cmp(-22)},
			[]Comparable{cmp(-22), cmp(1), cmp(3), cmp(7)},
		},
	}
	for _, e := range tests {
		Sort(e.Unsorted)
		fmt.Println("-->", e.Unsorted)
		assert.Equal(t, e.Expect, e.Unsorted)
	}
}
