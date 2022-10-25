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
