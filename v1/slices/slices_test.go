package slices

import (
	"fmt"
	"strconv"
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

func TestConvert(t *testing.T) {
	tests := []struct {
		Input  []string
		Output []int
		Func   func(string) (int, error)
	}{
		{
			[]string{"10", "20", "50"},
			[]int{10, 20, 50},
			strconv.Atoi,
		},
	}

	for _, e := range tests {
		v, err := Convert(e.Input, e.Func)
		if assert.Nil(t, err, fmt.Sprint(err)) {
			assert.Equal(t, e.Output, v)
		}
	}
}
