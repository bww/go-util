package sets

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	tests := []struct {
		A, B     []string
		Add, Del []string
	}{
		{
			nil,
			nil,
			[]string{},
			[]string{},
		},
		{
			nil,
			[]string{},
			[]string{},
			[]string{},
		},
		{
			[]string{},
			nil,
			[]string{},
			[]string{},
		},
		{
			[]string{},
			[]string{},
			[]string{},
			[]string{},
		},
		{
			[]string{"A"},
			[]string{"A"},
			[]string{},
			[]string{},
		},
		{
			[]string{"A"},
			[]string{"B"},
			[]string{"B"},
			[]string{"A"},
		},
		{
			[]string{"A", "A"},
			[]string{"B"},
			[]string{"B"},
			[]string{"A"},
		},
		{
			[]string{"A", "A"},
			[]string{"B", "B"},
			[]string{"B"},
			[]string{"A"},
		},
		{
			[]string{"A", "B"},
			[]string{"B"},
			[]string{},
			[]string{"A"},
		},
		{
			[]string{"A", "B", "C"},
			[]string{"B", "B"},
			[]string{},
			[]string{"A", "C"},
		},
		{
			[]string{"A", "B", "A", "C"},
			[]string{"C", "B"},
			[]string{},
			[]string{"A"},
		},
	}
	for i, e := range tests {
		fmt.Printf("--> #%d: [%s] → [%s]\n", i+1, strings.Join(e.A, ", "), strings.Join(e.B, ", "))
		add, del := DiffStrings(e.A, e.B)
		sort.Strings(add)
		sort.Strings(del)
		sort.Strings(e.Add)
		sort.Strings(e.Del)
		valid := true
		valid = valid && assert.Equal(t, e.Add, add)
		valid = valid && assert.Equal(t, e.Del, del)
		if !valid {
			fmt.Printf("--> #%d: [%s] → [%s] != add [%s], del [%s]\n", i+1, strings.Join(e.A, ", "), strings.Join(e.B, ", "), strings.Join(add, ", "), strings.Join(del, ", "))
		}
	}
}
