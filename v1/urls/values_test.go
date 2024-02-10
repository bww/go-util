package urls

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeValues(t *testing.T) {
	tests := []struct {
		A      url.Values
		B      url.Values
		Opts   []MergeOption
		Expect url.Values
	}{
		{
			A:      url.Values{},
			B:      url.Values{},
			Expect: url.Values{},
		},
		{
			A: url.Values{
				"a": []string{"1", "2"},
			},
			B: url.Values{},
			Expect: url.Values{
				"a": []string{"1", "2"},
			},
		},
		{
			A: url.Values{},
			B: url.Values{
				"a": []string{"1", "2"},
			},
			Expect: url.Values{
				"a": []string{"1", "2"},
			},
		},
		{
			A: url.Values{
				"a": []string{"x", "y"},
			},
			B: url.Values{
				"a": []string{"1", "2"},
			},
			Expect: url.Values{
				"a": []string{"1", "2"},
			},
		},
		{
			A: url.Values{
				"a": []string{"x", "y"},
			},
			B: url.Values{
				"a": []string{"1", "2"},
			},
			Opts: []MergeOption{
				Append(true),
			},
			Expect: url.Values{
				"a": []string{"x", "y", "1", "2"},
			},
		},
		{
			A: url.Values{
				"a": []string{"x", "y"},
			},
			B: url.Values{
				"b": []string{"1", "2"},
			},
			Expect: url.Values{
				"a": []string{"x", "y"},
				"b": []string{"1", "2"},
			},
		},
		{
			A: url.Values{
				"a": []string{"x", "y"},
			},
			B: url.Values{
				"a": []string{"x", "y"},
				"b": []string{"1", "2"},
			},
			Opts: []MergeOption{
				Append(true),
			},
			Expect: url.Values{
				"a": []string{"x", "y", "x", "y"},
				"b": []string{"1", "2"},
			},
		},
	}
	for _, e := range tests {
		r := MergeValues(e.A, e.B, e.Opts...)
		fmt.Println("-->", r)
		assert.Equal(t, e.Expect, r)
	}
}

func TestParseValueList(t *testing.T) {
	tests := []struct {
		Values url.Values
		Key    string
		Delim  string
		Expect []string
	}{
		{
			url.Values{
				"a": []string{"x", "y"},
			},
			"a", ",",
			[]string{"x", "y"},
		},
		{
			url.Values{
				"a": []string{"x,y"},
			},
			"a", ",",
			[]string{"x", "y"},
		},
		{
			url.Values{
				"a": []string{"x, y"},
			},
			"a", ",",
			[]string{"x", "y"},
		},
		{
			url.Values{
				"a": []string{"x, y", "z"},
			},
			"a", ",",
			[]string{"x", "y", "z"},
		},
	}
	for _, e := range tests {
		r := ParseValueList(e.Values, e.Key, e.Delim)
		fmt.Println("-->", r)
		assert.Equal(t, e.Expect, r)
	}
}
