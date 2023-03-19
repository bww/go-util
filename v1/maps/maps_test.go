package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		Input  map[int]string
		Output map[int]string
	}{
		{
			map[int]string{1: "One", 2: "Two"},
			map[int]string{1: "One", 2: "Two"},
		},
	}

	for _, e := range tests {
		assert.Equal(t, e.Output, Copy(e.Input))
	}
}
