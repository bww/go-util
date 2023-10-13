package slices

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {
	tests := []struct {
		Input  []string
		Output []string
		Func   func(string) string
	}{
		{
			[]string{"a", "b", "c"},
			[]string{"A", "B", "C"},
			strings.ToUpper,
		},
	}

	for _, e := range tests {
		assert.Equal(t, e.Output, Apply(e.Input, e.Func))
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		Input  []rune
		Output []int
		Func   func(rune) int
	}{
		{
			[]rune{'a', 'b', 'c'},
			[]int{0, 1, 2},
			func(c rune) int { // map a lowercase latin rune to it's alphabetic ordinal
				return int(c) - int('a')
			},
		},
	}

	for _, e := range tests {
		assert.Equal(t, e.Output, Map(e.Input, e.Func))
	}
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		Input  [][]int
		Output []int
	}{
		{
			[][]int{},
			[]int{},
		},
		{
			[][]int{
				{},
				{},
			},
			[]int{},
		},
		{
			[][]int{
				{0, 1, 2},
			},
			[]int{
				0, 1, 2,
			},
		},
		{
			[][]int{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
			},
			[]int{
				0, 1, 2, 3, 4, 5, 6, 7, 8,
			},
		},
	}

	for _, e := range tests {
		assert.Equal(t, e.Output, Flatten(e.Input))
	}
}

func TestFlatMap(t *testing.T) {
	tests := []struct {
		Input  []rune
		Output []int
		Func   func(rune) []int
	}{
		{
			[]rune{'a', 'b', 'c'},
			[]int{0, 0, 1, 0, 1, 2},
			func(c rune) []int { // map a lowercase latin rune to the set of indexes through it's alphabetic ordinal
				var s []int
				n := int(c) - int('a')
				for i := 0; i <= n; i++ {
					s = append(s, i)
				}
				return s
			},
		},
	}

	for _, e := range tests {
		assert.Equal(t, e.Output, FlatMap(e.Input, e.Func))
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		Input  []int
		Output []int
		Cmp    Comparator[int]
	}{
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
			func(a, b int) int {
				return a - b
			},
		},
		{
			[]int{1, 2, 3},
			[]int{3, 2, 1},
			func(a, b int) int { // compare inverse
				return b - a
			},
		},
	}

	for _, e := range tests {
		assert.Equal(t, e.Output, Sorted(e.Input, e.Cmp))
		Sort(e.Input, e.Cmp) // sorts input in place
		assert.Equal(t, e.Output, e.Input)
	}
}
