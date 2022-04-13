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
