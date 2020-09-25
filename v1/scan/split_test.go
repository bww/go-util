package scan

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const delim = ','

func TestSplit(t *testing.T) {
	tests := []struct {
		Str        string
		Parts      []string
		Delim, Esc rune
		Error      error
	}{
		{
			`A,B,C`,
			[]string{
				"A", "B", "C",
			},
			delim, esc,
			nil,
		},
		{
			`A\,B,C`,
			[]string{
				"A,B", "C",
			},
			delim, esc,
			nil,
		},
		{
			`A\,B\,C`,
			[]string{
				"A,B,C",
			},
			delim, esc,
			nil,
		},
		{
			`A\\,B\\,C`,
			[]string{
				"A\\", "B\\", "C",
			},
			delim, esc,
			nil,
		},
		{
			`A\x`,
			nil,
			delim, esc,
			ErrInvalidEscape,
		},
		{
			`,`,
			[]string{
				"",
			},
			delim, esc,
			nil,
		},
		{
			`\,`,
			[]string{
				",",
			},
			delim, esc,
			nil,
		},
		{
			``,
			nil,
			delim, esc,
			nil,
		},
	}
outer:
	for _, e := range tests {
		var p []string
		fmt.Println("-->", e.Str)

		var a string
		for s := e.Str; s != ""; {
			var err error
			a, s, err = Split(s, e.Delim, e.Esc)
			if e.Error != nil {
				assert.Equal(t, e.Error, err)
			} else {
				assert.Nil(t, err, fmt.Sprint(err))
			}
			if e.Error != nil || err != nil {
				continue outer
			}
			p = append(p, a)
		}

		fmt.Println("<--", strings.Join(p, "; "))
		assert.Equal(t, e.Parts, p)
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		Parts      []string
		Str        string
		Delim, Esc rune
	}{
		{
			[]string{
				"A", "B", "C",
			},
			`A,B,C`,
			delim, esc,
		},
		{
			[]string{
				"A,", "B,", "C",
			},
			`A\,,B\,,C`,
			delim, esc,
		},
		{
			[]string{
				"\\", "\\", "C",
			},
			`\\,\\,C`,
			delim, esc,
		},
		{
			[]string{
				"A,B", "C",
			},
			`A\,B,C`,
			delim, esc,
		},
		{
			[]string{},
			``,
			delim, esc,
		},
		{
			[]string{
				`,`,
			},
			`\,`,
			delim, esc,
		},
		{
			[]string{
				`\`,
			},
			`\\`,
			delim, esc,
		},
	}
	for _, e := range tests {
		fmt.Println("-->", strings.Join(e.Parts, "; "))
		s := Join(e.Parts, e.Delim, e.Esc)
		fmt.Println("<--", s)
		assert.Equal(t, e.Str, s)
	}
}
