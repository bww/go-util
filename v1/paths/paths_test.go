package paths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	tests := []struct {
		Input            string
		First, Remainder string
	}{
		{
			Input:     "",
			First:     "",
			Remainder: "",
		},
		{
			Input:     "a",
			First:     "a",
			Remainder: "",
		},
		{
			Input:     "a/b/c",
			First:     "a",
			Remainder: "b/c",
		},
		{
			Input:     "a\\/b/c",
			First:     "a/b",
			Remainder: "c",
		},
		{
			Input:     "\\/\\/\\/",
			First:     "///",
			Remainder: "",
		},
		{
			Input:     "\\\\\\\\\\\\/b",
			First:     "\\\\\\",
			Remainder: "b",
		},
	}
	for _, test := range tests {
		a, b := First(test.Input)
		fmt.Println(">>>", test.Input, "→", a, b)
		assert.Equal(t, test.First, a)
		assert.Equal(t, test.Remainder, b)
	}
}

func TestFirstConfig(t *testing.T) {
	tests := []struct {
		Input            string
		First, Remainder string
	}{
		{
			Input:     "",
			First:     "",
			Remainder: "",
		},
		{
			Input:     "a",
			First:     "a",
			Remainder: "",
		},
		{
			Input:     "a#b#c",
			First:     "a",
			Remainder: "b#c",
		},
		{
			Input:     "a\\#b#c",
			First:     "a#b",
			Remainder: "c",
		},
		{
			Input:     "\\#\\#\\#",
			First:     "###",
			Remainder: "",
		},
		{
			Input:     "\\\\\\\\\\\\#b",
			First:     "\\\\\\",
			Remainder: "b",
		},
	}
	for _, test := range tests {
		a, b := FirstConfig(test.Input, Config{Sep: '#', Esc: '\\'})
		fmt.Println(">>>", test.Input, "→", a, b)
		assert.Equal(t, test.First, a)
		assert.Equal(t, test.Remainder, b)
	}
}
