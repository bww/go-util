package urls

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
