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

func TestMapAny(t *testing.T) {
	tests := []struct {
		Input  []rune
		Output []any
	}{
		{
			[]rune{'a', 'b', 'c'},
			[]any{'a', 'b', 'c'},
		},
	}

	for _, e := range tests {
		assert.Equal(t, e.Output, MapAny(e.Input))
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

func TestFind(t *testing.T) {
	tests := []struct {
		Input  []int
		Search int
		Found  bool
	}{
		{
			Input:  []int{1, 2, 3},
			Search: 3,
			Found:  true,
		},
	}

	for _, e := range tests {
		res, ok := FindFunc(e.Input, MatchValue(e.Search))
		assert.Equal(t, e.Search, res)
		assert.Equal(t, e.Found, ok)
	}
}

func TestSummary(t *testing.T) {
	tests := []struct {
		Input  []int
		Limit  int
		Sep    string
		Expect string
	}{
		{
			Input:  []int{},
			Limit:  0,
			Sep:    ",",
			Expect: "",
		},
		{
			Input:  []int{},
			Limit:  3,
			Sep:    ",",
			Expect: "",
		},
		{
			Input:  []int{1, 2, 3},
			Limit:  3,
			Sep:    ",",
			Expect: "1,2,3",
		},
		{
			Input:  []int{1, 2, 3},
			Limit:  2,
			Sep:    ",",
			Expect: "1,2...1 more",
		},
		{
			Input:  []int{1, 2, 3},
			Limit:  1,
			Sep:    ",",
			Expect: "1...2 more",
		},
		{
			Input:  []int{1, 2, 3},
			Limit:  0,
			Sep:    ",",
			Expect: "1,2,3",
		},
	}
	for i, test := range tests {
		assert.Equal(t, test.Expect, Summary(test.Input, test.Sep, "...", test.Limit), "#%d", i)
	}
}
