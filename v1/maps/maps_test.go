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

func TestMergeMerged(t *testing.T) {
	tests := []struct {
		Input  []map[int]string
		Output map[int]string
	}{
		{
			nil,
			nil,
		},
		{
			[]map[int]string{},
			nil,
		},
		{
			[]map[int]string{
				map[int]string{1: "One", 2: "Two"},
			},
			map[int]string{1: "One", 2: "Two"},
		},
		{
			[]map[int]string{
				map[int]string{1: "One", 2: "Two"},
				map[int]string{1: "One", 2: "Two"},
			},
			map[int]string{1: "One", 2: "Two"},
		},
		{
			[]map[int]string{
				map[int]string{1: "One", 2: "Two"},
				map[int]string{3: "Three"},
			},
			map[int]string{1: "One", 2: "Two", 3: "Three"},
		},
		{
			[]map[int]string{
				map[int]string{1: "One", 2: "Two"},
				map[int]string{1: "Three"},
			},
			map[int]string{1: "Three", 2: "Two"},
		},
		{
			[]map[int]string{
				map[int]string{1: "One", 2: "Two"},
				map[int]string{1: "Three"},
				map[int]string{1: "Four"},
			},
			map[int]string{1: "Four", 2: "Two"},
		},
	}

	for _, e := range tests {
		assert.Equal(t, e.Output, Merged(e.Input...))
		assert.Equal(t, e.Output, Merge(e.Input...))
		if len(e.Input) > 0 {
			assert.Equal(t, e.Output, e.Input[0]) // input was mutated
		}
	}
}
